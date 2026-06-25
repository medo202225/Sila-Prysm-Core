package minimal

import (
	"testing"

	"github.com/sila-chain/Sila-Consensus-Core/v7/testing/spectest/shared/capella/finality"
)

func TestMinimal_Capella_Finality(t *testing.T) {
	finality.RunFinalityTest(t, "minimal")
}
