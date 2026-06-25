package mainnet

import (
	"testing"

	"github.com/sila-chain/Sila-Consensus-Core/v7/testing/spectest/shared/deneb/operations"
)

func TestMainnet_Deneb_Operations_VoluntaryExit(t *testing.T) {
	operations.RunVoluntaryExitTest(t, "mainnet")
}
