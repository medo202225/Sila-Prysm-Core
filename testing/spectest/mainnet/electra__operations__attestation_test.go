package mainnet

import (
	"testing"

	"github.com/sila-chain/Sila-Consensus-Core/v7/testing/spectest/shared/electra/operations"
)

func TestMainnet_Electra_Operations_Attestation(t *testing.T) {
	operations.RunAttestationTest(t, "mainnet")
}
