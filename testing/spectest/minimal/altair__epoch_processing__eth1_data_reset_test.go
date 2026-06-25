package minimal

import (
	"testing"

	"github.com/sila-chain/Sila-Consensus-Core/v7/testing/spectest/shared/altair/epoch_processing"
)

func TestMinimal_Altair_EpochProcessing_Eth1DataReset(t *testing.T) {
	epoch_processing.RunEth1DataResetTests(t, "minimal")
}
