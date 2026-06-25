package minimal

import (
	"testing"

	"github.com/sila-chain/Sila-Consensus-Core/v7/testing/spectest/shared/deneb/operations"
)

func TestMinimal_Deneb_Operations_Withdrawals(t *testing.T) {
	operations.RunWithdrawalsTest(t, "minimal")
}
