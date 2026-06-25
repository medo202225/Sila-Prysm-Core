package mainnet

import (
	"testing"

	"github.com/sila-chain/Sila-Consensus-Core/v7/testing/spectest/shared/altair/operations"
)

func TestMainnet_Altair_Operations_Attestation(t *testing.T) {
	operations.RunAttestationTest(t, "mainnet")
}
