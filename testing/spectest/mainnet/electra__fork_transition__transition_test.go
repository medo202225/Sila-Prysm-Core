package mainnet

import (
	"testing"

	"github.com/sila-chain/Sila-Consensus-Core/v7/testing/spectest/shared/electra/fork"
)

func TestMainnet_Electra_Transition(t *testing.T) {
	fork.RunForkTransitionTest(t, "mainnet")
}
