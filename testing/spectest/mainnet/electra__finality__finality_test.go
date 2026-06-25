package mainnet

import (
	"testing"

	"github.com/sila-chain/Sila-Consensus-Core/v7/testing/spectest/shared/electra/finality"
)

func TestMainnet_Electra_Finality(t *testing.T) {
	finality.RunFinalityTest(t, "mainnet")
}
