package mainnet

import (
	"testing"

	"github.com/sila-chain/Sila-Consensus-Core/v7/testing/spectest/shared/electra/operations"
)

func TestMainnet_Electra_Operations_AttesterSlashing(t *testing.T) {
	operations.RunAttesterSlashingTest(t, "mainnet")
}
