package minimal

import (
	"testing"

	"github.com/sila-chain/Sila-Consensus-Core/v7/testing/spectest/shared/capella/operations"
)

func TestMinimal_Capella_Operations_BlockHeader(t *testing.T) {
	operations.RunBlockHeaderTest(t, "minimal")
}
