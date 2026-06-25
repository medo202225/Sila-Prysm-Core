package utils

import (
	"testing"

	"github.com/sila-chain/Sila-Prysm-Core/v7/config/params"
	"github.com/sila-chain/Sila-Prysm-Core/v7/consensus-types/primitives"
	"github.com/sila-chain/Sila-Prysm-Core/v7/testing/require"
)

func TestConfig(t *testing.T) {
	require.NoError(t, SetConfig(t, "minimal"))
	require.Equal(t, primitives.Slot(8), params.BeaconConfig().SlotsPerEpoch)
	require.NoError(t, SetConfig(t, "mainnet"))
	require.Equal(t, primitives.Slot(32), params.BeaconConfig().SlotsPerEpoch)
}
