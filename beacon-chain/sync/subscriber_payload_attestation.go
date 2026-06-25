package sync

import (
	"context"

	"github.com/sila-chain/Sila-Consensus-Core/v7/beacon-chain/core/feed"
	opfeed "github.com/sila-chain/Sila-Consensus-Core/v7/beacon-chain/core/feed/operation"
	"github.com/sila-chain/Sila-Consensus-Core/v7/beacon-chain/core/gloas"
	eth "github.com/sila-chain/Sila-Consensus-Core/v7/proto/sila/v1alpha1"
	"google.golang.org/protobuf/proto"
)

func (s *Service) payloadAttestationSubscriber(ctx context.Context, msg proto.Message) error {
	a, ok := msg.(*eth.PayloadAttestationMessage)
	if !ok {
		return errWrongMessage
	}
	if a == nil || a.Data == nil {
		return errNilMessage
	}

	s.cfg.operationNotifier.OperationFeed().Send(&feed.Event{
		Type: opfeed.PayloadAttestationMessageReceived,
		Data: &opfeed.PayloadAttestationMessageReceivedData{
			Message: a,
		},
	})

	if err := s.payloadAttestationCache.Add(a.Data.Slot, a.ValidatorIndex); err != nil {
		return err
	}

	st, err := s.cfg.chain.HeadStateReadOnly(ctx)
	if err != nil {
		return err
	}
	idx, err := gloas.PayloadCommitteeIndex(ctx, st, a.Data.Slot, a.ValidatorIndex)
	if err != nil {
		return err
	}
	if err := s.cfg.payloadAttestationPool.InsertPayloadAttestation(a, idx); err != nil {
		return err
	}

	return s.cfg.chain.ReceivePayloadAttestationMessage(ctx, a)
}
