package mainnet

import (
	"testing"

	"github.com/sila-chain/Sila-Prysm-Core/v7/testing/spectest/shared/phase0/operations"
)

func TestMainnet_Phase0_Operations_Attestation(t *testing.T) {
	operations.RunAttestationTest(t, "mainnet")
}
