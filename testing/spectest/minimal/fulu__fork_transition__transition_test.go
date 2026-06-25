package minimal

import (
	"testing"

	"github.com/sila-chain/Sila-Consensus-Core/v7/testing/spectest/shared/fulu/fork"
)

func TestMinimal_Fulu_Transition(t *testing.T) {
	fork.RunForkTransitionTest(t, "minimal")
}
