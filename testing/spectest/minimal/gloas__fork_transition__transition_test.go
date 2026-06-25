package minimal

import (
	"testing"

	"github.com/sila-chain/Sila-Prysm-Core/v7/testing/spectest/shared/gloas/fork"
)

func TestMinimal_Gloas_Transition(t *testing.T) {
	fork.RunForkTransitionTest(t, "minimal")
}
