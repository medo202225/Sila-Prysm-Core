package mainnet

import (
	"testing"

	"github.com/sila-chain/Sila-Consensus-Core/v7/testing/spectest/shared/capella/fork"
)

func TestMainnet_Capella_UpgradeToCapella(t *testing.T) {
	fork.RunUpgradeToCapella(t, "mainnet")
}
