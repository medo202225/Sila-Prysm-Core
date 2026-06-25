package minimal

import (
	"testing"

	"github.com/sila-chain/Sila-Consensus-Core/v7/runtime/version"
	"github.com/sila-chain/Sila-Consensus-Core/v7/testing/spectest/shared/common/light_client"
)

func TestMainnet_Altair_LightClient_UpdateRanking(t *testing.T) {
	light_client.RunLightClientUpdateRankingTests(t, "minimal", version.Altair)
}
