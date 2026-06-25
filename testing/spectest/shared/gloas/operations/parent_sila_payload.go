package operations

import (
	"context"
	"path"
	"testing"

	"github.com/sila-chain/Sila-Consensus-Core/v7/beacon-chain/core/gloas"
	"github.com/sila-chain/Sila-Consensus-Core/v7/beacon-chain/core/helpers"
	"github.com/sila-chain/Sila-Consensus-Core/v7/beacon-chain/state"
	"github.com/sila-chain/Sila-Consensus-Core/v7/consensus-types/interfaces"
	"github.com/sila-chain/Sila-Consensus-Core/v7/runtime/version"
	"github.com/sila-chain/Sila-Consensus-Core/v7/testing/require"
	common "github.com/sila-chain/Sila-Consensus-Core/v7/testing/spectest/shared/common/operations"
	"github.com/sila-chain/Sila-Consensus-Core/v7/testing/spectest/utils"
	"github.com/sila-chain/Sila-Consensus-Core/v7/testing/util"
	"github.com/golang/snappy"
)

func RunParentSilaPayloadTest(t *testing.T, config string) {
	require.NoError(t, utils.SetConfig(t, config))
	testFolders, testsFolderPath := utils.TestFolders(t, config, version.String(version.Gloas), "operations/parent_sila_payload/pyspec_tests")
	if len(testFolders) == 0 {
		t.Fatalf("No test folders found for %s/%s/%s", config, version.String(version.Gloas), "operations/parent_sila_payload/pyspec_tests")
	}
	for _, folder := range testFolders {
		t.Run(folder.Name(), func(t *testing.T) {
			helpers.ClearCache()
			folderPath := path.Join(testsFolderPath, folder.Name())
			blockFile, err := util.BazelFileBytes(folderPath, "block.ssz_snappy")
			require.NoError(t, err)
			blockSSZ, err := snappy.Decode(nil, blockFile)
			require.NoError(t, err, "Failed to decompress")
			blk, err := sszToBlock(blockSSZ)
			require.NoError(t, err)

			common.RunBlockOperationTest(t, folderPath, blk, sszToState, func(ctx context.Context, s state.BeaconState, b interfaces.ReadOnlySignedBeaconBlock) (state.BeaconState, error) {
				if err := gloas.ProcessParentSilaPayload(ctx, s, b.Block()); err != nil {
					return nil, err
				}
				return s, nil
			})
		})
	}
}
