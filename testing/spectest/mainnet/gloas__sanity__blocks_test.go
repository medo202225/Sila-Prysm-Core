package mainnet

import (
	"testing"

	"github.com/sila-chain/Sila-Consensus-Core/v7/testing/spectest/shared/gloas/sanity"
)

func TestMainnet_Gloas_Sanity_Blocks(t *testing.T) {
	sanity.RunBlockProcessingTest(t, "mainnet", "sanity/blocks/pyspec_tests")
}
