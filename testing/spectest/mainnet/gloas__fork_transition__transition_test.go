package mainnet

import (
	"testing"

	"github.com/sila-chain/Sila-Prysm-Core/v7/testing/spectest/shared/gloas/fork"
)

func TestMainnet_Gloas_Transition(t *testing.T) {
	fork.RunForkTransitionTest(t, "mainnet")
}
