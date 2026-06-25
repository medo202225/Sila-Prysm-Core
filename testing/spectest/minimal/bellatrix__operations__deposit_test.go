package minimal

import (
	"testing"

	"github.com/sila-chain/Sila-Prysm-Core/v7/testing/spectest/shared/bellatrix/operations"
)

func TestMinimal_Bellatrix_Operations_Deposit(t *testing.T) {
	operations.RunDepositTest(t, "minimal")
}
