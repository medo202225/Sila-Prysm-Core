package mainnet

import (
	"testing"

	"github.com/sila-chain/Sila-Consensus-Core/v7/testing/spectest/shared/deneb/finality"
)

func TestMainnet_Deneb_Finality(t *testing.T) {
	finality.RunFinalityTest(t, "mainnet")
}
