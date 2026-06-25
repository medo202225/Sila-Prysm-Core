package minimal

import (
	"testing"

	"github.com/sila-chain/Sila-Consensus-Core/v7/testing/spectest/shared/bellatrix/operations"
)

func TestMinimal_Bellatrix_Operations_PayloadExecution(t *testing.T) {
	operations.RunExecutionPayloadTest(t, "minimal")
}
