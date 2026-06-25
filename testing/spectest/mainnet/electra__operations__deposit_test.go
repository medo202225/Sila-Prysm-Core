package mainnet

import (
	"testing"

	"github.com/sila-chain/Sila-Prysm-Core/v7/testing/spectest/shared/electra/operations"
)

func TestMainnet_Electra_Operations_Deposit(t *testing.T) {
	operations.RunDepositTest(t, "mainnet")
}
