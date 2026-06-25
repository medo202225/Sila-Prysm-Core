package operations

import (
	"testing"

	"github.com/sila-chain/Sila-Consensus-Core/v7/beacon-chain/core/altair"
	"github.com/sila-chain/Sila-Consensus-Core/v7/consensus-types/blocks"
	"github.com/sila-chain/Sila-Consensus-Core/v7/consensus-types/interfaces"
	ethpb "github.com/sila-chain/Sila-Consensus-Core/v7/proto/sila/v1alpha1"
	"github.com/sila-chain/Sila-Consensus-Core/v7/runtime/version"
	common "github.com/sila-chain/Sila-Consensus-Core/v7/testing/spectest/shared/common/operations"
)

func blockWithAttestation(attestationSSZ []byte) (interfaces.SignedBeaconBlock, error) {
	att := &ethpb.AttestationElectra{}
	if err := att.UnmarshalSSZ(attestationSSZ); err != nil {
		return nil, err
	}
	b := &ethpb.BeaconBlockGloas{}
	b.Body = &ethpb.BeaconBlockBodyGloas{Attestations: []*ethpb.AttestationElectra{att}}
	return blocks.NewSignedBeaconBlock(&ethpb.SignedBeaconBlockGloas{Block: b})
}

func RunAttestationTest(t *testing.T, config string) {
	common.RunAttestationTest(t, config, version.String(version.Gloas), blockWithAttestation, altair.ProcessAttestationsNoVerifySignature, sszToState)
}
