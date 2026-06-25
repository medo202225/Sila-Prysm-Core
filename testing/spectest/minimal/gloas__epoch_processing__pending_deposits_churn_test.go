package minimal

import (
	"testing"

	"github.com/sila-chain/Sila-Prysm-Core/v7/testing/spectest/shared/gloas/epoch_processing"
)

func TestMinimal_Gloas_EpochProcessing_PendingDepositsChurn(t *testing.T) {
	epoch_processing.RunPendingDepositsChurnTests(t, "minimal")
}
