package operations

import (
	"github.com/sila-chain/Sila-Prysm-Core/v7/beacon-chain/state"
	state_native "github.com/sila-chain/Sila-Prysm-Core/v7/beacon-chain/state/state-native"
	"github.com/sila-chain/Sila-Prysm-Core/v7/consensus-types/blocks"
	"github.com/sila-chain/Sila-Prysm-Core/v7/consensus-types/interfaces"
	ethpb "github.com/sila-chain/Sila-Prysm-Core/v7/proto/prysm/v1alpha1"
)

func sszToState(b []byte) (state.BeaconState, error) {
	base := &ethpb.BeaconStateGloas{}
	if err := base.UnmarshalSSZ(b); err != nil {
		return nil, err
	}
	return state_native.InitializeFromProtoUnsafeGloas(base)
}
func sszToBlock(b []byte) (interfaces.SignedBeaconBlock, error) {
	base := &ethpb.BeaconBlockGloas{}
	if err := base.UnmarshalSSZ(b); err != nil {
		return nil, err
	}
	return blocks.NewSignedBeaconBlock(&ethpb.SignedBeaconBlockGloas{Block: base})
}

func sszToBlockBody(b []byte) (interfaces.ReadOnlyBeaconBlockBody, error) {
	base := &ethpb.BeaconBlockBodyGloas{}
	if err := base.UnmarshalSSZ(b); err != nil {
		return nil, err
	}
	return blocks.NewBeaconBlockBody(base)
}
