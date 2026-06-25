package util

import (
	"testing"

	"github.com/sila-chain/Sila-Consensus-Core/v7/config/params"
	"github.com/sila-chain/Sila-Consensus-Core/v7/consensus-types/blocks"
	"github.com/sila-chain/Sila-Consensus-Core/v7/testing/require"
)

func TestInclusionProofs(t *testing.T) {
	ds := SlotAtEpoch(t, params.BeaconConfig().DenebForkEpoch)
	_, blobs := GenerateTestDenebBlockWithSidecar(t, [32]byte{}, ds, params.BeaconConfig().MaxBlobsPerBlock(ds))
	for i := range blobs {
		require.NoError(t, blocks.VerifyKZGInclusionProof(blobs[i]))
	}
}
