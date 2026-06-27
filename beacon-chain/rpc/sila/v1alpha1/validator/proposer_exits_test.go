package validator

import (
	"testing"

	"github.com/sila-chain/Sila-Consensus-Core/v7/beacon-chain/operations/voluntaryexits"
	"github.com/sila-chain/Sila-Consensus-Core/v7/config/params"
	"github.com/sila-chain/Sila-Consensus-Core/v7/consensus-types/primitives"
	sila "github.com/sila-chain/Sila-Consensus-Core/v7/proto/sila/v1alpha1"
	"github.com/sila-chain/Sila-Consensus-Core/v7/testing/require"
	"github.com/sila-chain/Sila-Consensus-Core/v7/testing/util"
)

func TestServer_getExits(t *testing.T) {
	params.SetupTestConfigCleanup(t)
	config := params.BeaconConfig()
	config.ShardCommitteePeriod = 0
	params.OverrideBeaconConfig(config)

	beaconState, privKeys := util.DeterministicGenesisState(t, 256)

	proposerServer := &Server{
		ExitPool: voluntaryexits.NewPool(),
	}

	exits := make([]*sila.SignedVoluntaryExit, params.BeaconConfig().MaxVoluntaryExits)
	for i := primitives.ValidatorIndex(0); uint64(i) < params.BeaconConfig().MaxVoluntaryExits; i++ {
		exit, err := util.GenerateVoluntaryExits(beaconState, privKeys[i], i)
		require.NoError(t, err)
		proposerServer.ExitPool.InsertVoluntaryExit(exit)
		exits[i] = exit
	}

	e := proposerServer.getExits(beaconState, 1)
	require.Equal(t, len(e), int(params.BeaconConfig().MaxVoluntaryExits))
	require.DeepEqual(t, e, exits)
}
