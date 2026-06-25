package mainnet

import (
	"testing"

	"github.com/sila-chain/Sila-Consensus-Core/v7/testing/spectest/shared/deneb/fork"
)

func TestMainnet_Deneb_Transition(t *testing.T) {
	fork.RunForkTransitionTest(t, "mainnet")
}
