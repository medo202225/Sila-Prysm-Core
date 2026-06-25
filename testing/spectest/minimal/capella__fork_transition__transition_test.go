package minimal

import (
	"testing"

	"github.com/sila-chain/Sila-Consensus-Core/v7/testing/spectest/shared/capella/fork"
)

func TestMinimal_Capella_Transition(t *testing.T) {
	fork.RunForkTransitionTest(t, "minimal")
}
