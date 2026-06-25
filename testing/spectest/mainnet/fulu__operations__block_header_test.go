package mainnet

import (
	"testing"

	"github.com/sila-chain/Sila-Consensus-Core/v7/testing/spectest/shared/fulu/operations"
)

func TestMainnet_Fulu_Operations_BlockHeader(t *testing.T) {
	operations.RunBlockHeaderTest(t, "mainnet")
}
