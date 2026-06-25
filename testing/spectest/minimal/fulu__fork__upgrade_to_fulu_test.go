package minimal

import (
	"testing"

	"github.com/sila-chain/Sila-Prysm-Core/v7/testing/spectest/shared/fulu/fork"
)

func TestMinimal_UpgradeToFulu(t *testing.T) {
	fork.RunUpgradeToFulu(t, "minimal")
}
