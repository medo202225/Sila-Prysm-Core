package minimal

import (
	"testing"

	"github.com/sila-chain/Sila-Consensus-Core/v7/testing/spectest/shared/phase0/epoch_processing"
)

func TestMinimal_Phase0_EpochProcessing_ParticipationRecordUpdates(t *testing.T) {
	epoch_processing.RunParticipationRecordUpdatesTests(t, "minimal")
}
