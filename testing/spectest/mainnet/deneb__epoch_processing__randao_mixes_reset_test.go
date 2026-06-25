package mainnet

import (
	"testing"

	"github.com/sila-chain/Sila-Consensus-Core/v7/testing/spectest/shared/deneb/epoch_processing"
)

func TestMainnet_Deneb_EpochProcessing_RandaoMixesReset(t *testing.T) {
	epoch_processing.RunRandaoMixesResetTests(t, "mainnet")
}
