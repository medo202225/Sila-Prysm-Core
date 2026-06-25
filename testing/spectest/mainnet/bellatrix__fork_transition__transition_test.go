package mainnet

import (
	"testing"

	"github.com/sila-chain/Sila-Consensus-Core/v7/testing/spectest/shared/bellatrix/fork"
)

func TestMainnet_Bellatrix_Transition(t *testing.T) {
	fork.RunForkTransitionTest(t, "mainnet")
}
