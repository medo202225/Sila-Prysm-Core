package minimal

import (
	"testing"

	"github.com/sila-chain/Sila-Prysm-Core/v7/testing/spectest/shared/gloas/fork"
)

func TestMinimal_UpgradeToGloas(t *testing.T) {
	fork.RunUpgradeToGloas(t, "minimal")
}
