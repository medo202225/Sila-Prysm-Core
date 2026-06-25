package minimal

import (
	"testing"

	"github.com/sila-chain/Sila-Consensus-Core/v7/testing/spectest/shared/altair/fork"
)

func TestMinimal_Altair_UpgradeToAltair(t *testing.T) {
	fork.RunUpgradeToAltair(t, "minimal")
}
