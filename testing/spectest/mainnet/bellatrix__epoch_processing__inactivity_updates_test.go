package mainnet

import (
	"testing"

	"github.com/sila-chain/Sila-Consensus-Core/v7/testing/spectest/shared/bellatrix/epoch_processing"
)

func TestMainnet_Bellatrix_EpochProcessing_InactivityUpdates(t *testing.T) {
	epoch_processing.RunInactivityUpdatesTest(t, "mainnet")
}
