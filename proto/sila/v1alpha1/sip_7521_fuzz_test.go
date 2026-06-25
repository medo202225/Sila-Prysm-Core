package eth_test

import (
	"testing"

	eth "github.com/sila-chain/Sila-Consensus-Core/v7/proto/sila/v1alpha1"
)

func TestCopySip7521Types_Fuzz(t *testing.T) {
	fuzzCopies(t, &eth.PendingDeposit{})
	fuzzCopies(t, &eth.PendingPartialWithdrawal{})
	fuzzCopies(t, &eth.PendingConsolidation{})
}
