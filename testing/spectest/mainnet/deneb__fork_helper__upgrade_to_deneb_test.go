package mainnet

import (
	"testing"

	"github.com/sila-chain/Sila-Consensus-Core/v7/testing/spectest/shared/deneb/fork"
)

func TestMainnet_UpgradeToDeneb(t *testing.T) {
	fork.RunUpgradeToDeneb(t, "mainnet")
}
