package minimal

import (
	"testing"

	"github.com/sila-chain/Sila-Consensus-Core/v7/testing/spectest/shared/bellatrix/fork"
)

func TestMinimal_Bellatrix_UpgradeToBellatrix(t *testing.T) {
	fork.RunUpgradeToBellatrix(t, "minimal")
}
