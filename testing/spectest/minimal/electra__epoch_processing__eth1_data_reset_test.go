package minimal

import (
	"testing"

	"github.com/sila-chain/Sila-Consensus-Core/v7/testing/spectest/shared/electra/epoch_processing"
)

func TestMinimal_Electra_EpochProcessing_Eth1DataReset(t *testing.T) {
	epoch_processing.RunEth1DataResetTests(t, "minimal")
}
