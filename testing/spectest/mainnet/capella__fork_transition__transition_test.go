package mainnet

import (
	"testing"

	"github.com/sila-chain/Sila-Prysm-Core/v7/testing/spectest/shared/capella/fork"
)

func TestMainnet_Capella_Transition(t *testing.T) {
	fork.RunForkTransitionTest(t, "mainnet")
}
