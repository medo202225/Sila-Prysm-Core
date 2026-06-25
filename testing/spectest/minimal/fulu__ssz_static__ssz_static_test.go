package minimal

import (
	"testing"

	"github.com/sila-chain/Sila-Prysm-Core/v7/testing/spectest/shared/fulu/ssz_static"
)

func TestMinimal_Fulu_SSZStatic(t *testing.T) {
	ssz_static.RunSSZStaticTests(t, "minimal")
}
