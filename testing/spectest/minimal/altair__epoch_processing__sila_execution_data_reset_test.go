package minimal

import (
	"testing"

	"github.com/sila-chain/Sila-Consensus-Core/v7/testing/spectest/shared/altair/epoch_processing"
)

func TestMinimal_Altair_EpochProcessing_SilaExecutionDataReset(t *testing.T) {
	epoch_processing.RunSilaExecutionDataResetTests(t, "minimal")
}
