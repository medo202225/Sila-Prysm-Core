package mainnet

import (
	"testing"

	"github.com/sila-chain/Sila-Consensus-Core/v7/testing/spectest/shared/phase0/operations"
)

func TestMainnet_Phase0_Operations_BlockHeader(t *testing.T) {
	operations.RunBlockHeaderTest(t, "mainnet")
}
