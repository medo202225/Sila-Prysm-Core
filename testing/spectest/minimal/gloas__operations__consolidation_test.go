package minimal

import (
	"testing"

	"github.com/sila-chain/Sila-Prysm-Core/v7/testing/spectest/shared/gloas/operations"
)

func TestMinimal_Gloas_Operations_Consolidation(t *testing.T) {
	operations.RunConsolidationTest(t, "minimal")
}
