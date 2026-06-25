package util

import (
	"testing"

	"github.com/sila-chain/Sila-Consensus-Core/v7/consensus-types/primitives"
	"github.com/sila-chain/Sila-Consensus-Core/v7/testing/require"
	"github.com/sila-chain/Sila-Consensus-Core/v7/time/slots"
)

func SlotAtEpoch(t *testing.T, e primitives.Epoch) primitives.Slot {
	s, err := slots.EpochStart(e)
	require.NoError(t, err)
	return s
}
