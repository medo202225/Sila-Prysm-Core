package mainnet

import (
	"testing"

	"github.com/sila-chain/Sila-Prysm-Core/v7/testing/spectest/shared/gloas/operations"
)

func TestMainnet_Gloas_Operations_ExecutionPayloadBid(t *testing.T) {
	operations.RunExecutionPayloadBidTest(t, "mainnet")
}
