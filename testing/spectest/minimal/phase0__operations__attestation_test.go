package minimal

import (
	"testing"

	"github.com/sila-chain/Sila-Prysm-Core/v7/testing/spectest/shared/phase0/operations"
)

func TestMinimal_Phase0_Operations_Attestation(t *testing.T) {
	operations.RunAttestationTest(t, "minimal")
}
