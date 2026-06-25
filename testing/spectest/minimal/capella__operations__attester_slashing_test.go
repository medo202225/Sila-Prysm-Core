package minimal

import (
	"testing"

	"github.com/sila-chain/Sila-Consensus-Core/v7/testing/spectest/shared/capella/operations"
)

func TestMinimal_Capella_Operations_AttesterSlashing(t *testing.T) {
	operations.RunAttesterSlashingTest(t, "minimal")
}
