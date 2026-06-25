package mainnet

import (
	"testing"

	"github.com/sila-chain/Sila-Prysm-Core/v7/testing/spectest/shared/fulu/fork"
)

func TestMainnet_UpgradeToFulu(t *testing.T) {
	fork.RunUpgradeToFulu(t, "mainnet")
}
