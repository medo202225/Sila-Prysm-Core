package minimal

import (
	"testing"

	"github.com/sila-chain/Sila-Consensus-Core/v7/testing/spectest/shared/deneb/fork"
)

func TestMinimal_UpgradeToDeneb(t *testing.T) {
	fork.RunUpgradeToDeneb(t, "minimal")
}
