package minimal

import (
	"testing"

	"github.com/sila-chain/Sila-Prysm-Core/v7/runtime/version"
	"github.com/sila-chain/Sila-Prysm-Core/v7/testing/spectest/shared/common/light_client"
)

func TestMinimal_Electra_LightClient_SingleMerkleProof(t *testing.T) {
	light_client.RunLightClientSingleMerkleProofTests(t, "minimal", version.Electra)
}
