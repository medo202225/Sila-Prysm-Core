package sync

import (
	"context"
	"fmt"

	"github.com/sila-chain/Sila-Consensus-Core/v7/beacon-chain/blockchain"
	p2ptypes "github.com/sila-chain/Sila-Consensus-Core/v7/beacon-chain/p2p/types"
	"github.com/sila-chain/Sila-Consensus-Core/v7/config/params"
	consensusblocks "github.com/sila-chain/Sila-Consensus-Core/v7/consensus-types/blocks"
	"github.com/sila-chain/Sila-Consensus-Core/v7/consensus-types/interfaces"
	"github.com/sila-chain/Sila-Consensus-Core/v7/crypto/rand"
	"github.com/sila-chain/Sila-Consensus-Core/v7/runtime/version"
	"github.com/sila-chain/Sila-Consensus-Core/v7/time/slots"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
	"github.com/pkg/errors"
)

// validateSilaPayloadBid validates sila payload bid gossip rules.
// [REJECT] The bid's parent (defined by bid.parent_block_root) equals the block's parent (defined by block.parent_root).
// [REJECT] The length of KZG commitments is less than or equal to the limitation defined in the consensus layer.
func (s *Service) validateSilaPayloadBid(ctx context.Context, blk interfaces.ReadOnlyBeaconBlock) (pubsub.ValidationResult, error) {
	if blk.Version() < version.Gloas {
		return pubsub.ValidationAccept, nil
	}
	signedBid, err := blk.Body().SignedSilaPayloadBid()
	if err != nil {
		return pubsub.ValidationIgnore, errors.Wrap(err, "unable to read bid from block")
	}
	wrappedBid, err := consensusblocks.WrappedROSignedSilaPayloadBid(signedBid)
	if err != nil {
		return pubsub.ValidationIgnore, errors.Wrap(err, "unable to wrap signed sila payload bid")
	}
	bid, err := wrappedBid.Bid()
	if err != nil {
		return pubsub.ValidationIgnore, errors.Wrap(err, "unable to read bid from signed sila payload bid")
	}

	if bid.ParentBlockRoot() != blk.ParentRoot() {
		return pubsub.ValidationReject, errors.New("bid parent block root does not match block parent root")
	}

	maxBlobsPerBlock := params.BeaconConfig().MaxBlobsPerBlockAtEpoch(slots.ToEpoch(blk.Slot()))
	if bid.BlobKzgCommitmentCount() > uint64(maxBlobsPerBlock) {
		return pubsub.ValidationReject, errors.Wrapf(errRejectCommitmentLen, "%d > %d", bid.BlobKzgCommitmentCount(), maxBlobsPerBlock)
	}

	return pubsub.ValidationAccept, nil
}

// validateSilaPayloadBidParentSeen validates parent payload gossip rules.
// [IGNORE] The block's parent sila payload (defined by bid.parent_block_hash) has been seen
// (via gossip or non-gossip sources) (a client MAY queue blocks for processing once the parent payload is retrieved).
func (s *Service) validateSilaPayloadBidParentSeen(_ context.Context, blk interfaces.ReadOnlyBeaconBlock) (pubsub.ValidationResult, error) {
	if blk.Version() < version.Gloas {
		return pubsub.ValidationAccept, nil
	}
	if s.cfg.chain.ParentPayloadReady(blk) {
		return pubsub.ValidationAccept, nil
	}
	return pubsub.ValidationIgnore, errors.New("parent payload not yet available")
}

// validateSilaPayloadBidParentValid validates parent payload verification status.
// If sila_payload verification of block's sila payload parent by a Sila node is complete:
// [REJECT] The block's sila payload parent (defined by bid.parent_block_hash) passes all validation.
func (s *Service) validateSilaPayloadBidParentValid(_ context.Context, blk interfaces.ReadOnlyBeaconBlock) (pubsub.ValidationResult, error) {
	if blk.Version() < version.Gloas {
		return pubsub.ValidationAccept, nil
	}
	if s.hasBadPayload(blk.ParentRoot()) {
		return pubsub.ValidationReject, errors.New("parent payload is invalid")
	}
	return pubsub.ValidationAccept, nil
}

func (s *Service) requestPayloadEnvelope(root [32]byte) {
	if s.cfg.chain.HasFullNode(root) || s.hasBadPayload(root) {
		return
	}
	key := fmt.Sprintf("%#x", root)
	_, _, _ = s.payloadEnvelopeRequestSingleFlight.Do(key, func() (any, error) {
		s.fetchPayloadEnvelope(root)
		return nil, nil
	})
}

const maxPayloadEnvelopeFetchAttempts = 3

func (s *Service) fetchPayloadEnvelope(root [32]byte) {
	bestPeers := s.getBestPeers()
	if len(bestPeers) == 0 {
		return
	}
	gen := rand.NewGenerator()
	gen.Shuffle(len(bestPeers), func(i, j int) { bestPeers[i], bestPeers[j] = bestPeers[j], bestPeers[i] })
	if len(bestPeers) > maxPayloadEnvelopeFetchAttempts {
		bestPeers = bestPeers[:maxPayloadEnvelopeFetchAttempts]
	}
	req := p2ptypes.SilaPayloadEnvelopesByRootReq{root}
	for _, pid := range bestPeers {
		if s.cfg.chain.HasFullNode(root) {
			return
		}
		envelopes, err := SendSilaPayloadEnvelopesByRootRequest(s.ctx, s.cfg.clock, s.cfg.p2p, pid, s.ctxMap, &req)
		if err != nil {
			log.WithError(err).WithField("peer", pid).Debug("Could not request payload envelope by root")
			continue
		}
		if len(envelopes) == 0 {
			continue
		}
		wrapped, err := consensusblocks.WrappedROSignedSilaPayloadEnvelope(envelopes[0])
		if err != nil {
			log.WithError(err).Debug("Could not wrap requested payload envelope")
			continue
		}
		if err := s.cfg.chain.ReceiveSilaPayloadEnvelope(s.ctx, wrapped); err != nil {
			if blockchain.IsInvalidBlock(err) {
				s.setBadPayload(s.ctx, root)
				return
			}
			log.WithError(err).Debug("Could not process requested payload envelope")
			continue
		}
		return
	}
}
