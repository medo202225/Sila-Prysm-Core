package minimal

import (
	"testing"

	"github.com/sila-chain/Sila-Consensus-Core/v7/testing/spectest/shared/deneb/operations"
)

func TestMinimal_Deneb_Operations_Deposit(t *testing.T) {
	operations.RunDepositTest(t, "minimal")
}
