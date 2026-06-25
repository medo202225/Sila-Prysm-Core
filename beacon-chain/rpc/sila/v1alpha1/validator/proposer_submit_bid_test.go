package validator

import (
	"testing"

	p2pmock "github.com/sila-chain/Sila-Consensus-Core/v7/beacon-chain/p2p/testing"
	mockSync "github.com/sila-chain/Sila-Consensus-Core/v7/beacon-chain/sync/initial-sync/testing"
	"github.com/sila-chain/Sila-Consensus-Core/v7/config/params"
	silapb "github.com/sila-chain/Sila-Consensus-Core/v7/proto/sila/v1alpha1"
	"github.com/sila-chain/Sila-Consensus-Core/v7/testing/assert"
	"github.com/sila-chain/Sila-Consensus-Core/v7/testing/require"
	"google.golang.org/protobuf/types/known/emptypb"
)

func TestSubmitSignedSilaPayloadBid_OK(t *testing.T) {
	params.SetupTestConfigCleanup(t)
	cfg := params.BeaconConfig().Copy()
	cfg.GloasForkEpoch = 0
	params.OverrideBeaconConfig(cfg)

	p2p := &p2pmock.MockBroadcaster{}
	vs := &Server{
		SyncChecker: &mockSync.Sync{IsSyncing: false},
		P2P:         p2p,
	}

	req := &silapb.SignedSilaPayloadBid{
		Message: &silapb.SilaPayloadBid{
			ParentBlockHash:       make([]byte, 32),
			ParentBlockRoot:       make([]byte, 32),
			BlockHash:             make([]byte, 32),
			PrevRandao:            make([]byte, 32),
			FeeRecipient:          make([]byte, 20),
			GasLimit:              30_000_000,
			BuilderIndex:          1,
			Slot:                  10,
			Value:                 100,
			SilaRequestsRoot: make([]byte, 32),
		},
		Signature: make([]byte, 96),
	}

	resp, err := vs.SubmitSignedSilaPayloadBid(t.Context(), req)
	require.NoError(t, err)
	require.DeepEqual(t, &emptypb.Empty{}, resp)
	assert.Equal(t, true, p2p.BroadcastCalled.Load())
	require.Equal(t, 1, len(p2p.BroadcastMessages))
}

func TestSubmitSignedSilaPayloadBid_NilRequest(t *testing.T) {
	vs := &Server{
		SyncChecker: &mockSync.Sync{IsSyncing: false},
	}
	_, err := vs.SubmitSignedSilaPayloadBid(t.Context(), nil)
	require.ErrorContains(t, "nil", err)
}

func TestSubmitSignedSilaPayloadBid_NilMessage(t *testing.T) {
	vs := &Server{
		SyncChecker: &mockSync.Sync{IsSyncing: false},
	}
	_, err := vs.SubmitSignedSilaPayloadBid(t.Context(), &silapb.SignedSilaPayloadBid{})
	require.ErrorContains(t, "nil", err)
}

func TestSubmitSignedSilaPayloadBid_Syncing(t *testing.T) {
	vs := &Server{
		SyncChecker: &mockSync.Sync{IsSyncing: true},
	}
	req := &silapb.SignedSilaPayloadBid{
		Message: &silapb.SilaPayloadBid{Slot: 10},
	}
	_, err := vs.SubmitSignedSilaPayloadBid(t.Context(), req)
	require.ErrorContains(t, "Syncing", err)
}

func TestSubmitSignedSilaPayloadBid_PreGloas(t *testing.T) {
	params.SetupTestConfigCleanup(t)
	cfg := params.BeaconConfig().Copy()
	cfg.GloasForkEpoch = 100
	params.OverrideBeaconConfig(cfg)

	vs := &Server{
		SyncChecker: &mockSync.Sync{IsSyncing: false},
	}
	req := &silapb.SignedSilaPayloadBid{
		Message: &silapb.SilaPayloadBid{Slot: 10},
	}
	_, err := vs.SubmitSignedSilaPayloadBid(t.Context(), req)
	require.ErrorContains(t, "not supported before Gloas", err)
}
