package minimal

import (
	"testing"

	"github.com/sila-chain/Sila-Prysm-Core/v7/testing/spectest/shared/bellatrix/rewards"
)

func TestMinimal_Bellatrix_Rewards(t *testing.T) {
	rewards.RunPrecomputeRewardsAndPenaltiesTests(t, "minimal")
}
