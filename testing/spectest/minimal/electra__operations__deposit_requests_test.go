package minimal

import (
	"testing"

	"github.com/sila-chain/Sila-Consensus-Core/v7/testing/spectest/shared/electra/operations"
)

func TestMainnet_Electra_Operations_DepositRequests(t *testing.T) {
	operations.RunDepositRequestsTest(t, "minimal")
}
