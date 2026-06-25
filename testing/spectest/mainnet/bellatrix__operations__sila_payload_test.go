package mainnet

import (
	"testing"

	"github.com/sila-chain/Sila-Consensus-Core/v7/testing/spectest/shared/bellatrix/operations"
)

func TestMainnet_Bellatrix_Operations_PayloadExecution(t *testing.T) {
	operations.RunSilaPayloadTest(t, "mainnet")
}
