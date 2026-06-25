package mainnet

import (
	"testing"

	"github.com/sila-chain/Sila-Consensus-Core/v7/testing/spectest/shared/fulu/epoch_processing"
)

func TestMainnet_Fulu_EpochProcessing_PendingConsolidations(t *testing.T) {
	epoch_processing.RunPendingConsolidationsTests(t, "mainnet")
}
