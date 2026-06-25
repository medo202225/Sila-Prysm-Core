package minimal

import (
	"testing"

	"github.com/sila-chain/Sila-Consensus-Core/v7/testing/spectest/shared/gloas/epoch_processing"
)

func TestMinimal_Gloas_EpochProcessing_ParticipationFlagUpdates(t *testing.T) {
	epoch_processing.RunParticipationFlagUpdatesTests(t, "minimal")
}
