package mainnet

import (
	"testing"

	"github.com/sila-chain/Sila-Consensus-Core/v7/testing/spectest/shared/gloas/operations"
)

func TestMainnet_Gloas_Operations_Attestation(t *testing.T) {
	operations.RunAttestationTest(t, "mainnet")
}
