package minimal

import (
	"testing"

	"github.com/sila-chain/Sila-Consensus-Core/v7/testing/spectest/shared/electra/operations"
)

func TestMinimal_Electra_Operations_PayloadExecution(t *testing.T) {
	operations.RunExecutionPayloadTest(t, "minimal")
}
