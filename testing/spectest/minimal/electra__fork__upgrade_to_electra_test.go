package minimal

import (
	"testing"

	"github.com/sila-chain/Sila-Consensus-Core/v7/testing/spectest/shared/electra/fork"
)

func TestMinimal_UpgradeToElectra(t *testing.T) {
	fork.RunUpgradeToElectra(t, "minimal")
}
