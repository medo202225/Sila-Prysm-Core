package minimal

import (
	"testing"

	"github.com/sila-chain/Sila-Prysm-Core/v7/runtime/version"
	"github.com/sila-chain/Sila-Prysm-Core/v7/testing/spectest/shared/common/forkchoice"
)

func TestMinimal_Fulu_Forkchoice(t *testing.T) {
	forkchoice.Run(t, "minimal", version.Fulu)
}
