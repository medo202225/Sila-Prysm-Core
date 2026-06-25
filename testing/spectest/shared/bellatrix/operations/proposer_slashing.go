package operations

import (
	"testing"

	"github.com/sila-chain/Sila-Consensus-Core/v7/consensus-types/blocks"
	"github.com/sila-chain/Sila-Consensus-Core/v7/consensus-types/interfaces"
	ethpb "github.com/sila-chain/Sila-Consensus-Core/v7/proto/sila/v1alpha1"
	"github.com/sila-chain/Sila-Consensus-Core/v7/runtime/version"
	common "github.com/sila-chain/Sila-Consensus-Core/v7/testing/spectest/shared/common/operations"
	"github.com/sila-chain/Sila-Consensus-Core/v7/testing/util"
)

func blockWithProposerSlashing(ssz []byte) (interfaces.SignedBeaconBlock, error) {
	ps := &ethpb.ProposerSlashing{}
	if err := ps.UnmarshalSSZ(ssz); err != nil {
		return nil, err
	}
	b := util.NewBeaconBlockBellatrix()
	b.Block.Body = &ethpb.BeaconBlockBodyBellatrix{ProposerSlashings: []*ethpb.ProposerSlashing{ps}}
	return blocks.NewSignedBeaconBlock(b)
}

func RunProposerSlashingTest(t *testing.T, config string) {
	common.RunProposerSlashingTest(t, config, version.String(version.Bellatrix), blockWithProposerSlashing, sszToState)
}
