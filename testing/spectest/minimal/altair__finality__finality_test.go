package minimal

import (
	"testing"

	"github.com/sila-chain/Sila-Consensus-Core/v7/testing/spectest/shared/altair/finality"
)

func TestMinimal_Altair_Finality(t *testing.T) {
	finality.RunFinalityTest(t, "minimal")
}
