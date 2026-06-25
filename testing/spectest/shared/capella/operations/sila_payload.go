package operations

import (
	"testing"

	"github.com/sila-chain/Sila-Consensus-Core/v7/runtime/version"
	common "github.com/sila-chain/Sila-Consensus-Core/v7/testing/spectest/shared/common/operations"
)

func RunSilaPayloadTest(t *testing.T, config string) {
	common.RunSilaPayloadTest(t, config, version.String(version.Capella), sszToBlockBody, sszToState)
}
