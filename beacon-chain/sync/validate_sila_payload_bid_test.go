package sync

import (
	"bytes"
	"context"
	"reflect"
	"testing"
	"time"

	mock "github.com/sila-chain/Sila-Consensus-Core/v7/beacon-chain/blockchain/testing"
	"github.com/sila-chain/Sila-Consensus-Core/v7/beacon-chain/cache"
	dbtest "github.com/sila-chain/Sila-Consensus-Core/v7/beacon-chain/db/testing"
	"github.com/sila-chain/Sila-Consensus-Core/v7/beacon-chain/p2p"
	p2ptest "github.com/sila-chain/Sila-Consensus-Core/v7/beacon-chain/p2p/testing"
	"github.com/sila-chain/Sila-Consensus-Core/v7/beacon-chain/startup"
	"github.com/sila-chain/Sila-Consensus-Core/v7/beacon-chain/state"
	mockSync "github.com/sila-chain/Sila-Consensus-Core/v7/beacon-chain/sync/initial-sync/testing"
	"github.com/sila-chain/Sila-Consensus-Core/v7/beacon-chain/verification"
	"github.com/sila-chain/Sila-Consensus-Core/v7/config/params"
	"github.com/sila-chain/Sila-Consensus-Core/v7/consensus-types/blocks"
	"github.com/sila-chain/Sila-Consensus-Core/v7/consensus-types/interfaces"
	"github.com/sila-chain/Sila-Consensus-Core/v7/consensus-types/primitives"
	"github.com/sila-chain/Sila-Consensus-Core/v7/encoding/bytesutil"
	silapb "github.com/sila-chain/Sila-Consensus-Core/v7/proto/sila/v1alpha1"
	"github.com/sila-chain/Sila-Consensus-Core/v7/testing/require"
	"github.com/sila-chain/Sila-Consensus-Core/v7/testing/util"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
	pb "github.com/libp2p/go-libp2p-pubsub/pb"
	"github.com/pkg/errors"
	"google.golang.org/protobuf/proto"
)

func TestValidateSilaPayloadBidGossip_InvalidTopic(t *testing.T) {
	ctx := context.Background()
	p := p2ptest.NewTestP2P(t)
	s := &Service{cfg: &config{p2p: p, initialSync: &mockSync.Sync{}}}

	result, err := s.validateSilaPayloadBidGossip(ctx, "", &pubsub.Message{Message: &pb.Message{}})
	require.ErrorIs(t, p2p.ErrInvalidTopic, err)
	require.Equal(t, pubsub.ValidationReject, result)
}

func TestValidateSilaPayloadBidGossip_AlreadySeenBuilder(t *testing.T) {
	ctx := context.Background()
	s, msg, signedBid := setupSilaPayloadBidService(t)
	s.newSilaPayloadBidVerifier = testNewSilaPayloadBidVerifier(mockSilaPayloadBidVerifier{})

	key := silaPayloadBidBuilderKey(signedBid.Message.Slot, signedBid.Message.BuilderIndex)
	s.setSeenSilaPayloadBidBuilder(signedBid.Message.Slot, key)
	result, err := s.validateSilaPayloadBidGossip(ctx, "", msg)
	require.NoError(t, err)
	require.Equal(t, pubsub.ValidationIgnore, result)
}

// Dedup must short-circuit before every later check; duplicates pay only the cache lookup.
func TestValidateSilaPayloadBidGossip_DedupShortCircuitsAllLaterChecks(t *testing.T) {
	ctx := context.Background()
	s, msg, signedBid := setupSilaPayloadBidService(t)
	key := silaPayloadBidBuilderKey(signedBid.Message.Slot, signedBid.Message.BuilderIndex)
	s.setSeenSilaPayloadBidBuilder(signedBid.Message.Slot, key)
	// Every subsequent verifier method would Reject/Ignore if it ran; the cache hit must skip them all.
	s.newSilaPayloadBidVerifier = testNewSilaPayloadBidVerifier(mockSilaPayloadBidVerifier{
		errCurrentOrNextSlot:    errors.New("slot"),
		errBuilderActive:        errors.New("builder"),
		errExecutionPayment:     errors.New("payment"),
		errFeeRecipientMismatch: errors.New("fee"),
		errSignature:            errors.New("sig"),
	})

	result, err := s.validateSilaPayloadBidGossip(ctx, "", msg)
	require.NoError(t, err)
	require.Equal(t, pubsub.ValidationIgnore, result)
}

func TestValidateSilaPayloadBidGossip_ProposerPreferencesUnseen(t *testing.T) {
	ctx := context.Background()
	s, msg, _ := setupSilaPayloadBidService(t)
	s.proposerPreferencesCache.Clear()
	s.newSilaPayloadBidVerifier = testNewSilaPayloadBidVerifier(mockSilaPayloadBidVerifier{})

	result, err := s.validateSilaPayloadBidGossip(ctx, "", msg)
	require.NoError(t, err)
	require.Equal(t, pubsub.ValidationIgnore, result)
}

func TestValidateSilaPayloadBidGossip_InitialSync(t *testing.T) {
	ctx := context.Background()
	p := p2ptest.NewTestP2P(t)
	s := &Service{
		cfg: &config{
			p2p:         p,
			initialSync: &mockSync.Sync{IsSyncing: true},
		},
	}

	result, err := s.validateSilaPayloadBidGossip(ctx, "", &pubsub.Message{})
	require.NoError(t, err)
	require.Equal(t, pubsub.ValidationIgnore, result)
}

func TestValidateSilaPayloadBidGossip_ErrorPathsWithMock(t *testing.T) {
	ctx := context.Background()
	tests := []struct {
		name      string
		verifier  mockSilaPayloadBidVerifier
		result    pubsub.ValidationResult
		wantError bool
	}{
		{
			name:      "slot out of range",
			verifier:  mockSilaPayloadBidVerifier{errCurrentOrNextSlot: errors.New("wrong slot")},
			result:    pubsub.ValidationIgnore,
			wantError: true,
		},
		{
			name:      "non-zero execution payment",
			verifier:  mockSilaPayloadBidVerifier{errExecutionPayment: errors.New("non-zero payment")},
			result:    pubsub.ValidationReject,
			wantError: true,
		},
		{
			name:      "fee recipient mismatch",
			verifier:  mockSilaPayloadBidVerifier{errFeeRecipientMismatch: errors.New("wrong fee recipient")},
			result:    pubsub.ValidationReject,
			wantError: true,
		},
		{
			name:      "gas limit incompatible",
			verifier:  mockSilaPayloadBidVerifier{errGasLimitIncompatible: errors.New("incompatible gas limit")},
			result:    pubsub.ValidationIgnore,
			wantError: true,
		},
		{
			name:      "parent root unknown",
			verifier:  mockSilaPayloadBidVerifier{errParentBlockRootSeen: errors.New("unknown root")},
			result:    pubsub.ValidationIgnore,
			wantError: true,
		},
		{
			name:      "inactive builder",
			verifier:  mockSilaPayloadBidVerifier{errBuilderActive: errors.New("inactive builder")},
			result:    pubsub.ValidationReject,
			wantError: true,
		},
		{
			name:      "slot not higher than parent",
			verifier:  mockSilaPayloadBidVerifier{errSlotHigherThanParent: errors.New("slot not higher than parent")},
			result:    pubsub.ValidationReject,
			wantError: true,
		},
		{
			name:      "parent hash mismatch",
			verifier:  mockSilaPayloadBidVerifier{errParentBlockHash: errors.New("wrong hash")},
			result:    pubsub.ValidationIgnore,
			wantError: true,
		},
		{
			name:      "builder cannot cover",
			verifier:  mockSilaPayloadBidVerifier{errBuilderCanCoverBid: errors.New("cannot cover")},
			result:    pubsub.ValidationIgnore,
			wantError: true,
		},
		{
			name:      "invalid signature",
			verifier:  mockSilaPayloadBidVerifier{errSignature: errors.New("bad signature")},
			result:    pubsub.ValidationReject,
			wantError: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			s, msg, _ := setupSilaPayloadBidService(t)
			s.newSilaPayloadBidVerifier = testNewSilaPayloadBidVerifier(tc.verifier)

			result, err := s.validateSilaPayloadBidGossip(ctx, "", msg)
			if tc.wantError {
				require.NotNil(t, err)
			}
			require.Equal(t, tc.result, result)
		})
	}
}

func TestValidateSilaPayloadBidGossip_LowerOrEqualBidIgnored(t *testing.T) {
	ctx := context.Background()
	s, msg, signedBid := setupSilaPayloadBidService(t)
	s.newSilaPayloadBidVerifier = testNewSilaPayloadBidVerifier(mockSilaPayloadBidVerifier{})

	s.setHighestSilaPayloadBid(signedBid)

	var err error
	result, err := s.validateSilaPayloadBidGossip(ctx, "", msg)
	require.NoError(t, err)
	require.Equal(t, pubsub.ValidationIgnore, result)
	builderKey := silaPayloadBidBuilderKey(signedBid.Message.Slot, signedBid.Message.BuilderIndex)
	require.Equal(t, true, s.hasSeenSilaPayloadBidBuilder(builderKey))
}

func TestValidateSilaPayloadBidGossip_LowerBidIgnoredStillMarksBuilderSeen(t *testing.T) {
	ctx := context.Background()
	s, msg, signedBid := setupSilaPayloadBidService(t)
	s.newSilaPayloadBidVerifier = testNewSilaPayloadBidVerifier(mockSilaPayloadBidVerifier{})

	higherBid := proto.Clone(signedBid).(*silapb.SignedSilaPayloadBid)
	higherBid.Message.Value = signedBid.Message.Value + 1
	s.setHighestSilaPayloadBid(higherBid)

	result, err := s.validateSilaPayloadBidGossip(ctx, "", msg)
	require.NoError(t, err)
	require.Equal(t, pubsub.ValidationIgnore, result)

	// If the lower valid bid did not mark the builder as seen, the same bid would
	// be accepted once the highest-bid cache is cleared.
	s.highestSilaPayloadBidCache = cache.NewHighestSilaPayloadBidCache()
	msg = silaPayloadBidToPubsub(t, s, s.cfg.p2p, signedBid)

	result, err = s.validateSilaPayloadBidGossip(ctx, "", msg)
	require.NoError(t, err)
	require.Equal(t, pubsub.ValidationIgnore, result)
}

func TestValidateSilaPayloadBidGossip_HigherBidAccepted(t *testing.T) {
	ctx := context.Background()
	s, msg, signedBid := setupSilaPayloadBidService(t)
	s.newSilaPayloadBidVerifier = testNewSilaPayloadBidVerifier(mockSilaPayloadBidVerifier{})

	wrapped, err := blocks.WrappedROSignedSilaPayloadBid(signedBid)
	require.NoError(t, err)
	bid, err := wrapped.Bid()
	require.NoError(t, err)
	lowerBid := proto.Clone(signedBid).(*silapb.SignedSilaPayloadBid)
	lowerBid.Message.Value = bid.Value() - 1
	s.setHighestSilaPayloadBid(lowerBid)

	result, err := s.validateSilaPayloadBidGossip(ctx, "", msg)
	require.NoError(t, err)
	require.Equal(t, pubsub.ValidationAccept, result)
}

func TestValidateSilaPayloadBidGossip_HappyPath(t *testing.T) {
	ctx := context.Background()
	s, msg, signedBid := setupSilaPayloadBidService(t)
	s.newSilaPayloadBidVerifier = testNewSilaPayloadBidVerifier(mockSilaPayloadBidVerifier{})

	result, err := s.validateSilaPayloadBidGossip(ctx, "", msg)
	require.NoError(t, err)
	require.Equal(t, pubsub.ValidationAccept, result)

	builderKey := silaPayloadBidBuilderKey(signedBid.Message.Slot, signedBid.Message.BuilderIndex)
	require.Equal(t, true, s.hasSeenSilaPayloadBidBuilder(builderKey))
	got, ok := msg.ValidatorData.(*silapb.SignedSilaPayloadBid)
	require.Equal(t, true, ok)
	require.DeepEqual(t, signedBid, got)
}

func TestValidateSilaPayloadBidGossip_FeeRecipientMismatch(t *testing.T) {
	ctx := context.Background()
	s, msg, _ := setupSilaPayloadBidService(t)
	s.newSilaPayloadBidVerifier = testNewSilaPayloadBidVerifier(
		mockSilaPayloadBidVerifier{errFeeRecipientMismatch: verification.ErrBidFeeRecipientMismatch},
	)

	result, err := s.validateSilaPayloadBidGossip(ctx, "", msg)
	require.NotNil(t, err)
	require.Equal(t, pubsub.ValidationReject, result)
	require.ErrorIs(t, err, verification.ErrBidFeeRecipientMismatch)
}

func TestValidateSilaPayloadBidGossip_GasLimitIncompatible(t *testing.T) {
	ctx := context.Background()
	s, msg, _ := setupSilaPayloadBidService(t)
	s.newSilaPayloadBidVerifier = testNewSilaPayloadBidVerifier(
		mockSilaPayloadBidVerifier{errGasLimitIncompatible: verification.ErrBidGasLimitIncompatible},
	)

	result, err := s.validateSilaPayloadBidGossip(ctx, "", msg)
	require.NotNil(t, err)
	require.Equal(t, pubsub.ValidationIgnore, result)
	require.ErrorIs(t, err, verification.ErrBidGasLimitIncompatible)
}

func TestSilaPayloadBidSubscriber_WrongMessage(t *testing.T) {
	s := &Service{}
	err := s.silaPayloadBidSubscriber(context.Background(), &silapb.BeaconBlock{})
	require.ErrorIs(t, errWrongMessage, err)
}

func TestSilaPayloadBidSubscriber_HappyPath(t *testing.T) {
	s := &Service{
		highestSilaPayloadBidCache: cache.NewHighestSilaPayloadBidCache(),
	}
	signedBid := util.GenerateTestSignedSilaPayloadBid(1)
	err := s.silaPayloadBidSubscriber(context.Background(), signedBid)
	require.NoError(t, err)
	bid := mustBid(t, signedBid)
	got, ok := s.highestSilaPayloadBidCache.Get(bid.Slot(), bid.ParentBlockHash(), bid.ParentBlockRoot())
	require.Equal(t, true, ok)
	require.DeepEqual(t, signedBid, got)
}

func TestSilaPayloadBidSubscriber_NilMessage(t *testing.T) {
	s := &Service{
		highestSilaPayloadBidCache: cache.NewHighestSilaPayloadBidCache(),
	}
	err := s.silaPayloadBidSubscriber(context.Background(), &silapb.SignedSilaPayloadBid{})
	require.ErrorIs(t, errNilMessage, err)
}

type mockSilaPayloadBidVerifier struct {
	errCurrentOrNextSlot    error
	errBuilderActive        error
	errExecutionPayment     error
	errFeeRecipientMismatch error
	errGasLimitIncompatible error
	errParentBlockRootSeen  error
	errSlotHigherThanParent error
	errParentBlockHash      error
	errBuilderCanCoverBid   error
	errSignature            error
}

var _ verification.SilaPayloadBidVerifier = &mockSilaPayloadBidVerifier{}

func (m *mockSilaPayloadBidVerifier) VerifyCurrentOrNextSlot() error {
	return m.errCurrentOrNextSlot
}

func (m *mockSilaPayloadBidVerifier) VerifyBuilderActive(state.ReadOnlyBeaconState) error {
	return m.errBuilderActive
}

func (m *mockSilaPayloadBidVerifier) VerifyExecutionPaymentZero() error {
	return m.errExecutionPayment
}

func (m *mockSilaPayloadBidVerifier) VerifyFeeRecipientMatches([]byte) error {
	return m.errFeeRecipientMismatch
}

func (m *mockSilaPayloadBidVerifier) VerifyGasLimitTargetCompatible(uint64, uint64) error {
	return m.errGasLimitIncompatible
}

func (m *mockSilaPayloadBidVerifier) VerifyParentBlockRootSeen(func([32]byte) bool) error {
	return m.errParentBlockRootSeen
}

func (m *mockSilaPayloadBidVerifier) VerifyBidSlotHigherThanParent(primitives.Slot) error {
	return m.errSlotHigherThanParent
}

func (m *mockSilaPayloadBidVerifier) VerifyParentBlockHash(func([32]byte) ([32]byte, error)) error {
	return m.errParentBlockHash
}

func (m *mockSilaPayloadBidVerifier) VerifyBuilderCanCoverBid(state.ReadOnlyBeaconState) error {
	return m.errBuilderCanCoverBid
}

func (m *mockSilaPayloadBidVerifier) VerifySignature(state.ReadOnlyBeaconState) error {
	return m.errSignature
}

func (*mockSilaPayloadBidVerifier) SatisfyRequirement(verification.Requirement) {}

func testNewSilaPayloadBidVerifier(m mockSilaPayloadBidVerifier) verification.NewSilaPayloadBidVerifier {
	return func(interfaces.ROSignedSilaPayloadBid, []verification.Requirement) verification.SilaPayloadBidVerifier {
		clone := m
		return &clone
	}
}

func setupSilaPayloadBidService(t *testing.T) (*Service, *pubsub.Message, *silapb.SignedSilaPayloadBid) {
	t.Helper()

	params.SetupTestConfigCleanup(t)
	cfg := params.BeaconConfig().Copy()
	cfg.FuluForkEpoch = 0
	cfg.GloasForkEpoch = 0
	params.OverrideBeaconConfig(cfg)

	ctx := context.Background()
	db := dbtest.SetupDB(t)
	p := p2ptest.NewTestP2P(t)

	// Save a genesis block so beaconDB.GenesisBlockRoot resolves; bids at slot 1
	// (epoch 0) hit the underflow branch in chain.DependentRootForEpoch which
	// falls back to the genesis block root.
	gb := util.NewBeaconBlock()
	signedGenesis, err := blocks.NewSignedBeaconBlock(gb)
	require.NoError(t, err)
	require.NoError(t, db.SaveBlock(ctx, signedGenesis))
	genesisRoot, err := signedGenesis.Block().HashTreeRoot()
	require.NoError(t, err)
	require.NoError(t, db.SaveGenesisBlockRoot(ctx, genesisRoot))

	state, err := util.NewBeaconStateGloas()
	require.NoError(t, err)
	signedBid := util.GenerateTestSignedSilaPayloadBid(1)
	signedBid.Message.BuilderIndex = 1
	chainService := &mock.ChainService{
		Genesis:    time.Now(),
		State:      state,
		TargetRoot: genesisRoot,
		ForkchoiceRoots: map[[32]byte]bool{
			[32]byte{0x02}: true,
		},
		ForkchoiceBlockHashes: map[[32]byte][32]byte{[32]byte{0x02}: [32]byte{0x01}},
		ForkchoiceGasLimits:   map[[32]byte]uint64{[32]byte{0x02}: 1},
	}
	s := &Service{
		seenSilaPayloadBidCache:    newSlotAwareCache(10),
		highestSilaPayloadBidCache: cache.NewHighestSilaPayloadBidCache(),
		proposerPreferencesCache:        cache.NewProposerPreferencesCache(),
		cfg: &config{
			p2p:         p,
			initialSync: &mockSync.Sync{},
			chain:       chainService,
			beaconDB:    db,
			clock:       startup.NewClock(chainService.Genesis, chainService.ValidatorsRoot),
		},
	}
	// The Gloas test state has a zero-filled proposer lookahead, so the
	// proposer for any slot is validator index 0.
	require.Equal(t, true, s.proposerPreferencesCache.Add(cache.ProposerPreference{
		DependentRoot:  genesisRoot,
		ValidatorIndex: 0,
		FeeRecipient:   bytesutil.ToBytes20(signedBid.Message.FeeRecipient),
		TargetGasLimit: signedBid.Message.GasLimit,
	}, signedBid.Message.Slot))
	msg := silaPayloadBidToPubsub(t, s, p, signedBid)
	return s, msg, signedBid
}

func silaPayloadBidToPubsub(t *testing.T, s *Service, p p2p.P2P, bid *silapb.SignedSilaPayloadBid) *pubsub.Message {
	t.Helper()

	buf := new(bytes.Buffer)
	_, err := p.Encoding().EncodeGossip(buf, bid)
	require.NoError(t, err)

	topic := p2p.GossipTypeMapping[reflect.TypeFor[*silapb.SignedSilaPayloadBid]()]
	digest, err := s.currentForkDigest()
	require.NoError(t, err)
	topic = s.addDigestToTopic(topic, digest)

	return &pubsub.Message{
		Message: &pb.Message{
			Data:  buf.Bytes(),
			Topic: &topic,
		},
	}
}

func mustBid(t *testing.T, signedBid *silapb.SignedSilaPayloadBid) interfaces.ROSilaPayloadBid {
	t.Helper()

	wrapped, err := blocks.WrappedROSignedSilaPayloadBid(signedBid)
	require.NoError(t, err)
	bid, err := wrapped.Bid()
	require.NoError(t, err)
	return bid
}
