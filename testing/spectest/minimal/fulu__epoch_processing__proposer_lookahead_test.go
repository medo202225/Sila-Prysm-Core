package minimal

import (
	"testing"

	"github.com/sila-chain/Sila-Prysm-Core/v7/testing/spectest/shared/fulu/epoch_processing"
)

func TestMinimal_fulu_EpochProcessing_ProposerLookahead(t *testing.T) {
	epoch_processing.RunProposerLookaheadTests(t, "minimal")
}
