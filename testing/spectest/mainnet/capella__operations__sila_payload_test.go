package mainnet

import (
	"testing"

	"github.com/sila-chain/Sila-Consensus-Core/v7/testing/spectest/shared/capella/operations"
)

func TestMainnet_Capella_Operations_PayloadExecution(t *testing.T) {
	operations.RunSilaPayloadTest(t, "mainnet")
}
