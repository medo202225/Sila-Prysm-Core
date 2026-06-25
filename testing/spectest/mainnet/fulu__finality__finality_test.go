package mainnet

import (
	"testing"

	"github.com/sila-chain/Sila-Consensus-Core/v7/testing/spectest/shared/fulu/finality"
)

func TestMainnet_Fulu_Finality(t *testing.T) {
	finality.RunFinalityTest(t, "mainnet")
}
