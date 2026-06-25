package mainnet

import (
	"testing"

	"github.com/sila-chain/Sila-Prysm-Core/v7/testing/spectest/shared/fulu/operations"
)

func TestMainnet_Fulu_Operations_Consolidation(t *testing.T) {
	operations.RunConsolidationTest(t, "mainnet")
}
