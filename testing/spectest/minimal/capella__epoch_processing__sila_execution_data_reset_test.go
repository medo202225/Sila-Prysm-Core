package minimal

import (
	"testing"

	"github.com/sila-chain/Sila-Consensus-Core/v7/testing/spectest/shared/capella/epoch_processing"
)

func TestMinimal_Capella_EpochProcessing_SilaExecutionDataReset(t *testing.T) {
	epoch_processing.RunSilaExecutionDataResetTests(t, "minimal")
}
