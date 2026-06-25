package mainnet

import (
	"testing"

	"github.com/sila-chain/Sila-Consensus-Core/v7/testing/spectest/shared/gloas/operations"
)

func TestMainnet_Gloas_Operations_ParentSilaPayload(t *testing.T) {
	operations.RunParentSilaPayloadTest(t, "mainnet")
}
