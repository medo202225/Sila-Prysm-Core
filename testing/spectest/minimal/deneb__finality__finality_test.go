package minimal

import (
	"testing"

	"github.com/sila-chain/Sila-Consensus-Core/v7/testing/spectest/shared/deneb/finality"
)

func TestMinimal_Deneb_Finality(t *testing.T) {
	finality.RunFinalityTest(t, "minimal")
}
