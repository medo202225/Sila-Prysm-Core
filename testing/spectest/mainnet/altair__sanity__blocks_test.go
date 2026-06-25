package mainnet

import (
	"testing"

	"github.com/sila-chain/Sila-Consensus-Core/v7/testing/spectest/shared/altair/sanity"
)

func TestMainnet_Altair_Sanity_Blocks(t *testing.T) {
	sanity.RunBlockProcessingTest(t, "mainnet", "sanity/blocks/pyspec_tests")
}
