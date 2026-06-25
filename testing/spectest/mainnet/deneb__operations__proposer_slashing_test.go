package mainnet

import (
	"testing"

	"github.com/sila-chain/Sila-Consensus-Core/v7/testing/spectest/shared/deneb/operations"
)

func TestMainnet_Deneb_Operations_ProposerSlashing(t *testing.T) {
	operations.RunProposerSlashingTest(t, "mainnet")
}
