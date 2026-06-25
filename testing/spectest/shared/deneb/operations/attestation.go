package operations

import (
	"testing"

	"github.com/sila-chain/Sila-Prysm-Core/v7/beacon-chain/core/altair"
	"github.com/sila-chain/Sila-Prysm-Core/v7/consensus-types/blocks"
	"github.com/sila-chain/Sila-Prysm-Core/v7/consensus-types/interfaces"
	ethpb "github.com/sila-chain/Sila-Prysm-Core/v7/proto/prysm/v1alpha1"
	"github.com/sila-chain/Sila-Prysm-Core/v7/runtime/version"
	common "github.com/sila-chain/Sila-Prysm-Core/v7/testing/spectest/shared/common/operations"
	"github.com/sila-chain/Sila-Prysm-Core/v7/testing/util"
)

func blockWithAttestation(attestationSSZ []byte) (interfaces.SignedBeaconBlock, error) {
	att := &ethpb.Attestation{}
	if err := att.UnmarshalSSZ(attestationSSZ); err != nil {
		return nil, err
	}
	b := util.NewBeaconBlockDeneb()
	b.Block.Body = &ethpb.BeaconBlockBodyDeneb{Attestations: []*ethpb.Attestation{att}}
	return blocks.NewSignedBeaconBlock(b)
}

func RunAttestationTest(t *testing.T, config string) {
	common.RunAttestationTest(t, config, version.String(version.Deneb), blockWithAttestation, altair.ProcessAttestationsNoVerifySignature, sszToState)
}
