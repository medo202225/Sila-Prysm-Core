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

func blockWithAttesterSlashing(asSSZ []byte) (interfaces.SignedBeaconBlock, error) {
	as := &ethpb.AttesterSlashing{}
	if err := as.UnmarshalSSZ(asSSZ); err != nil {
		return nil, err
	}
	b := util.NewBeaconBlockCapella()
	b.Block.Body = &ethpb.BeaconBlockBodyCapella{AttesterSlashings: []*ethpb.AttesterSlashing{as}}
	return blocks.NewSignedBeaconBlock(b)
}

func RunAttesterSlashingTest(t *testing.T, config string) {
	common.RunAttesterSlashingTest(t, config, version.String(version.Capella), blockWithAttesterSlashing, sszToState)
}
