package minimal

import (
	"testing"

	"github.com/sila-chain/Sila-Consensus-Core/v7/testing/spectest/shared/fulu/operations"
)

func TestMainnet_Fulu_Operations_DepositRequests(t *testing.T) {
	operations.RunDepositRequestsTest(t, "minimal")
}
