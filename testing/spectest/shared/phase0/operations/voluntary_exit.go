package operations

import (
	"testing"

	"github.com/sila-chain/Sila-Prysm-Core/v7/consensus-types/blocks"
	"github.com/sila-chain/Sila-Prysm-Core/v7/consensus-types/interfaces"
	ethpb "github.com/sila-chain/Sila-Prysm-Core/v7/proto/prysm/v1alpha1"
	"github.com/sila-chain/Sila-Prysm-Core/v7/runtime/version"
	common "github.com/sila-chain/Sila-Prysm-Core/v7/testing/spectest/shared/common/operations"
	"github.com/sila-chain/Sila-Prysm-Core/v7/testing/util"
)

func blockWithVoluntaryExit(ssz []byte) (interfaces.SignedBeaconBlock, error) {
	e := &ethpb.SignedVoluntaryExit{}
	if err := e.UnmarshalSSZ(ssz); err != nil {
		return nil, err
	}
	b := util.NewBeaconBlock()
	b.Block.Body = &ethpb.BeaconBlockBody{VoluntaryExits: []*ethpb.SignedVoluntaryExit{e}}
	return blocks.NewSignedBeaconBlock(b)
}

func RunVoluntaryExitTest(t *testing.T, config string) {
	common.RunVoluntaryExitTest(t, config, version.String(version.Phase0), blockWithVoluntaryExit, sszToState)
}
