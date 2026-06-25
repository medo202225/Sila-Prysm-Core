package mainnet

import (
	"testing"

	"github.com/sila-chain/Sila-Consensus-Core/v7/testing/spectest/shared/fulu/epoch_processing"
)

func TestMainnet_fulu_EpochProcessing_EffectiveBalanceUpdates(t *testing.T) {
	epoch_processing.RunEffectiveBalanceUpdatesTests(t, "mainnet")
}
