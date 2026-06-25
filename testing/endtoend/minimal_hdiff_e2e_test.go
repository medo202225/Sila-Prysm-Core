package endtoend

import (
	"testing"

	"github.com/sila-chain/Sila-Consensus-Core/v7/config/params"
	"github.com/sila-chain/Sila-Consensus-Core/v7/runtime/version"
	"github.com/sila-chain/Sila-Consensus-Core/v7/testing/endtoend/types"
)

func TestEndToEnd_MinimalConfig_WithStateDiff(t *testing.T) {
	r := e2eMinimal(t, types.InitForkCfg(version.Bellatrix, version.Electra, params.E2ETestConfig()),
		types.WithStateDiff(),
	)
	r.run()
}
