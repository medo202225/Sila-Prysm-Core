package minimal

import (
	"testing"

	"github.com/sila-chain/Sila-Consensus-Core/v7/testing/spectest/shared/electra/ssz_static"
)

func TestMinimal_Electra_SSZStatic(t *testing.T) {
	ssz_static.RunSSZStaticTests(t, "minimal")
}
