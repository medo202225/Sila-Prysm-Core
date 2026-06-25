package mainnet

import (
	"testing"

	"github.com/sila-chain/Sila-Consensus-Core/v7/testing/spectest/shared/electra/epoch_processing"
)

func TestMainnet_Electra_EpochProcessing_InactivityUpdates(t *testing.T) {
	epoch_processing.RunInactivityUpdatesTest(t, "mainnet")
}
