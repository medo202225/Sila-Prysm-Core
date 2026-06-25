package operations

import (
	"testing"

	b "github.com/sila-chain/Sila-Prysm-Core/v7/beacon-chain/core/blocks"
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
	b := util.NewBeaconBlock()
	b.Block.Body = &ethpb.BeaconBlockBody{Attestations: []*ethpb.Attestation{att}}
	return blocks.NewSignedBeaconBlock(b)
}

// RunAttestationTest executes "operations/attestation" tests.
func RunAttestationTest(t *testing.T, config string) {
	common.RunAttestationTest(t, config, version.String(version.Phase0), blockWithAttestation, b.ProcessAttestationsNoVerifySignature, sszToState)
}
