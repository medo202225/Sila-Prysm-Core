package minimal

import (
	"testing"

	"github.com/sila-chain/Sila-Consensus-Core/v7/testing/spectest/shared/bellatrix/operations"
)

func TestMinimal_Bellatrix_Operations_Attestation(t *testing.T) {
	operations.RunAttestationTest(t, "minimal")
}
