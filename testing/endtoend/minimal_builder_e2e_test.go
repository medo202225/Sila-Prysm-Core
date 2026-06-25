package endtoend

import (
	"testing"

	"github.com/sila-chain/Sila-Prysm-Core/v7/config/params"
	"github.com/sila-chain/Sila-Prysm-Core/v7/runtime/version"
	"github.com/sila-chain/Sila-Prysm-Core/v7/testing/endtoend/types"
)

func TestEndToEnd_MinimalConfig_WithBuilder(t *testing.T) {
	r := e2eMinimal(t, types.InitForkCfg(version.Bellatrix, version.Electra, params.E2ETestConfig()), types.WithCheckpointSync(), types.WithBuilder())
	r.run()
}

func TestEndToEnd_MinimalConfig_WithBuilder_ValidatorRESTApi(t *testing.T) {
	r := e2eMinimal(t, types.InitForkCfg(version.Bellatrix, version.Electra, params.E2ETestConfig()), types.WithCheckpointSync(), types.WithBuilder(), types.WithValidatorRESTApi())
	r.run()
}
