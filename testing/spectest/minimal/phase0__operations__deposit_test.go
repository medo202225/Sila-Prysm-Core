package minimal

import (
	"testing"

	"github.com/sila-chain/Sila-Consensus-Core/v7/testing/spectest/shared/phase0/operations"
)

func TestMinimal_Phase0_Operations_Deposit(t *testing.T) {
	operations.RunDepositTest(t, "minimal")
}
