package mainnet

import (
	"testing"

	"github.com/sila-chain/Sila-Consensus-Core/v7/testing/spectest/shared/deneb/operations"
)

func TestMainnet_Deneb_Operations_BLSToExecutionChange(t *testing.T) {
	operations.RunBLSToExecutionChangeTest(t, "mainnet")
}
