package minimal

import (
	"testing"

	"github.com/sila-chain/Sila-Prysm-Core/v7/testing/spectest/shared/fulu/epoch_processing"
)

func TestMinimal_Fulu_EpochProcessing_PendingDeposits(t *testing.T) {
	epoch_processing.RunPendingDepositsTests(t, "minimal")
}
