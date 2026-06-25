package epoch_processing

import (
	"context"
	"path"
	"testing"

	"github.com/sila-chain/Sila-Consensus-Core/v7/beacon-chain/core/electra"
	"github.com/sila-chain/Sila-Consensus-Core/v7/beacon-chain/core/helpers"
	"github.com/sila-chain/Sila-Consensus-Core/v7/beacon-chain/state"
	"github.com/sila-chain/Sila-Consensus-Core/v7/testing/require"
	"github.com/sila-chain/Sila-Consensus-Core/v7/testing/spectest/utils"
)

func RunSyncCommitteeUpdatesTests(t *testing.T, config string) {
	require.NoError(t, utils.SetConfig(t, config))

	testFolders, testsFolderPath := utils.TestFolders(t, config, "fulu", "epoch_processing/sync_committee_updates/pyspec_tests")
	for _, folder := range testFolders {
		t.Run(folder.Name(), func(t *testing.T) {
			helpers.ClearCache()
			folderPath := path.Join(testsFolderPath, folder.Name())
			RunEpochOperationTest(t, folderPath, processSyncCommitteeUpdates)
		})
	}
}

func processSyncCommitteeUpdates(t *testing.T, st state.BeaconState) (state.BeaconState, error) {
	return electra.ProcessSyncCommitteeUpdates(context.Background(), st)
}
