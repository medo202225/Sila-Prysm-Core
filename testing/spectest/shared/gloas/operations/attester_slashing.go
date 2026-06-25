package operations

import (
	"testing"

	"github.com/sila-chain/Sila-Consensus-Core/v7/consensus-types/blocks"
	"github.com/sila-chain/Sila-Consensus-Core/v7/consensus-types/interfaces"
	ethpb "github.com/sila-chain/Sila-Consensus-Core/v7/proto/sila/v1alpha1"
	"github.com/sila-chain/Sila-Consensus-Core/v7/runtime/version"
	common "github.com/sila-chain/Sila-Consensus-Core/v7/testing/spectest/shared/common/operations"
)

func blockWithAttesterSlashing(asSSZ []byte) (interfaces.SignedBeaconBlock, error) {
	as := &ethpb.AttesterSlashingElectra{}
	if err := as.UnmarshalSSZ(asSSZ); err != nil {
		return nil, err
	}
	b := &ethpb.SignedBeaconBlockGloas{
		Block: &ethpb.BeaconBlockGloas{
			Body: &ethpb.BeaconBlockBodyGloas{AttesterSlashings: []*ethpb.AttesterSlashingElectra{as}},
		},
	}
	return blocks.NewSignedBeaconBlock(b)
}

func RunAttesterSlashingTest(t *testing.T, config string) {
	common.RunAttesterSlashingTest(t, config, version.String(version.Gloas), blockWithAttesterSlashing, sszToState)
}
