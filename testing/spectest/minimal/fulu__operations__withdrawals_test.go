package minimal

import (
	"testing"

	"github.com/sila-chain/Sila-Consensus-Core/v7/testing/spectest/shared/fulu/operations"
)

func TestMinimal_Fulu_Operations_Withdrawals(t *testing.T) {
	operations.RunWithdrawalsTest(t, "minimal")
}
