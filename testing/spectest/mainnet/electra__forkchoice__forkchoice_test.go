package mainnet

import (
	"testing"

	"github.com/sila-chain/Sila-Consensus-Core/v7/runtime/version"
	"github.com/sila-chain/Sila-Consensus-Core/v7/testing/spectest/shared/common/forkchoice"
)

func TestMainnet_Electra_Forkchoice(t *testing.T) {
	forkchoice.Run(t, "mainnet", version.Electra)
}
