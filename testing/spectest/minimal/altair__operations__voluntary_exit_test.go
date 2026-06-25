package minimal

import (
	"testing"

	"github.com/sila-chain/Sila-Consensus-Core/v7/testing/spectest/shared/altair/operations"
)

func TestMinimal_Altair_Operations_VoluntaryExit(t *testing.T) {
	operations.RunVoluntaryExitTest(t, "minimal")
}
