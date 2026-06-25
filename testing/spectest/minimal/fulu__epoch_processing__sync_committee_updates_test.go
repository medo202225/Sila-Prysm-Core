package minimal

import (
	"testing"

	"github.com/sila-chain/Sila-Consensus-Core/v7/testing/spectest/shared/fulu/epoch_processing"
)

func TestMinimal_Fulu_EpochProcessing_SyncCommitteeUpdates(t *testing.T) {
	epoch_processing.RunSyncCommitteeUpdatesTests(t, "minimal")
}
