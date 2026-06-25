package operations

import (
	"github.com/sila-chain/Sila-Consensus-Core/v7/beacon-chain/state"
	state_native "github.com/sila-chain/Sila-Consensus-Core/v7/beacon-chain/state/state-native"
	"github.com/sila-chain/Sila-Consensus-Core/v7/consensus-types/blocks"
	"github.com/sila-chain/Sila-Consensus-Core/v7/consensus-types/interfaces"
	ethpb "github.com/sila-chain/Sila-Consensus-Core/v7/proto/sila/v1alpha1"
)

func sszToState(b []byte) (state.BeaconState, error) {
	base := &ethpb.BeaconStateFulu{}
	if err := base.UnmarshalSSZ(b); err != nil {
		return nil, err
	}
	return state_native.InitializeFromProtoUnsafeFulu(base)
}

func sszToBlock(b []byte) (interfaces.SignedBeaconBlock, error) {
	base := &ethpb.BeaconBlockElectra{}
	if err := base.UnmarshalSSZ(b); err != nil {
		return nil, err
	}
	return blocks.NewSignedBeaconBlock(&ethpb.SignedBeaconBlockElectra{Block: base})
}

func sszToBlockBody(b []byte) (interfaces.ReadOnlyBeaconBlockBody, error) {
	base := &ethpb.BeaconBlockBodyElectra{}
	if err := base.UnmarshalSSZ(b); err != nil {
		return nil, err
	}
	return blocks.NewBeaconBlockBody(base)
}
