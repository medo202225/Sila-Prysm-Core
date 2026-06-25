package minimal

import (
	"testing"

	"github.com/sila-chain/Sila-Consensus-Core/v7/testing/spectest/shared/fulu/fork"
)

func TestMinimal_UpgradeToFulu(t *testing.T) {
	fork.RunUpgradeToFulu(t, "minimal")
}
