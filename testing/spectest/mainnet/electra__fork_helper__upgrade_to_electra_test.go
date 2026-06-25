package mainnet

import (
	"testing"

	"github.com/sila-chain/Sila-Consensus-Core/v7/testing/spectest/shared/electra/fork"
)

func TestMainnet_UpgradeToElectra(t *testing.T) {
	fork.RunUpgradeToElectra(t, "mainnet")
}
