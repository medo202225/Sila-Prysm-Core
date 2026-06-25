package mainnet

import (
	"testing"

	"github.com/sila-chain/Sila-Consensus-Core/v7/testing/spectest/shared/capella/operations"
)

func TestMainnet_Capella_Operations_Attestation(t *testing.T) {
	operations.RunAttestationTest(t, "mainnet")
}
