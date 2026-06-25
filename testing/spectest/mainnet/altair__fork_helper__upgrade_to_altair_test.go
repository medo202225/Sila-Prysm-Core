package mainnet

import (
	"testing"

	"github.com/sila-chain/Sila-Consensus-Core/v7/testing/spectest/shared/altair/fork"
)

func TestMainnet_Altair_UpgradeToAltair(t *testing.T) {
	fork.RunUpgradeToAltair(t, "mainnet")
}
