package minimal

import (
	"testing"

	"github.com/sila-chain/Sila-Consensus-Core/v7/testing/spectest/shared/gloas/operations"
)

func TestMinimal_Gloas_Operations_ProposerSlashing(t *testing.T) {
	operations.RunProposerSlashingTest(t, "minimal")
}
