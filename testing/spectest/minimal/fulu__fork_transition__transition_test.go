package minimal

import (
	"testing"

	"github.com/sila-chain/Sila-Prysm-Core/v7/testing/spectest/shared/fulu/fork"
)

func TestMinimal_Fulu_Transition(t *testing.T) {
	fork.RunForkTransitionTest(t, "minimal")
}
