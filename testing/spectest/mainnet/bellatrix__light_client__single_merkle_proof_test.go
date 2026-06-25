package mainnet

import (
	"testing"

	"github.com/sila-chain/Sila-Consensus-Core/v7/runtime/version"
	"github.com/sila-chain/Sila-Consensus-Core/v7/testing/spectest/shared/common/light_client"
)

func TestMainnet_Bellatrix_LightClient_SingleMerkleProof(t *testing.T) {
	light_client.RunLightClientSingleMerkleProofTests(t, "mainnet", version.Bellatrix)
}
