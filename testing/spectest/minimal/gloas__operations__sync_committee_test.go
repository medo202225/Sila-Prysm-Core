package minimal

import (
	"testing"

	"github.com/sila-chain/Sila-Consensus-Core/v7/testing/spectest/shared/gloas/operations"
)

func TestMinimal_Gloas_Operations_SyncCommittee(t *testing.T) {
	operations.RunSyncCommitteeTest(t, "minimal")
}
