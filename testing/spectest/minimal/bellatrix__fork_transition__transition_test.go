package minimal

import (
	"testing"

	"github.com/sila-chain/Sila-Prysm-Core/v7/testing/spectest/shared/bellatrix/fork"
)

func TestMinimal_Bellatrix_Transition(t *testing.T) {
	fork.RunForkTransitionTest(t, "minimal")
}
