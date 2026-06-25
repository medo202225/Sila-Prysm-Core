package mainnet

import (
	"testing"

	"github.com/sila-chain/Sila-Consensus-Core/v7/testing/spectest/shared/capella/epoch_processing"
)

func TestMainnet_Capella_EpochProcessing_Slashings(t *testing.T) {
	epoch_processing.RunSlashingsTests(t, "mainnet")
}
