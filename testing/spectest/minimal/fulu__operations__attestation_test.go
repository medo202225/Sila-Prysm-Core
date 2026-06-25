package minimal

import (
	"testing"

	"github.com/sila-chain/Sila-Prysm-Core/v7/testing/spectest/shared/fulu/operations"
)

func TestMinimal_Fulu_Operations_Attestation(t *testing.T) {
	operations.RunAttestationTest(t, "minimal")
}
