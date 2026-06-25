package minimal

import (
	"testing"

	"github.com/sila-chain/Sila-Consensus-Core/v7/testing/spectest/shared/deneb/epoch_processing"
)

func TestMinimal_Deneb_EpochProcessing_ResetRegistryUpdates(t *testing.T) {
	epoch_processing.RunRegistryUpdatesTests(t, "minimal")
}
