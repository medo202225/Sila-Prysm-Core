package state_native_test

import (
	"testing"

	state_native "github.com/sila-chain/Sila-Consensus-Core/v7/beacon-chain/state/state-native"
	"github.com/sila-chain/Sila-Consensus-Core/v7/consensus-types/primitives"
	sila "github.com/sila-chain/Sila-Consensus-Core/v7/proto/sila/v1alpha1"
	"github.com/sila-chain/Sila-Consensus-Core/v7/testing/require"
)

func TestAppendPendingConsolidation(t *testing.T) {
	s, err := state_native.InitializeFromProtoElectra(&sila.BeaconStateElectra{})
	require.NoError(t, err)
	num, err := s.NumPendingConsolidations()
	require.NoError(t, err)
	require.Equal(t, uint64(0), num)
	require.NoError(t, s.AppendPendingConsolidation(&sila.PendingConsolidation{}))
	num, err = s.NumPendingConsolidations()
	require.NoError(t, err)
	require.Equal(t, uint64(1), num)

	pc := make([]*sila.PendingConsolidation, 0, 4)
	require.NoError(t, s.SetPendingConsolidations(pc))
	require.NoError(t, s.AppendPendingConsolidation(&sila.PendingConsolidation{SourceIndex: 1}))
	s2 := s.Copy()
	require.NoError(t, s2.AppendPendingConsolidation(&sila.PendingConsolidation{SourceIndex: 3}))
	require.NoError(t, s.AppendPendingConsolidation(&sila.PendingConsolidation{SourceIndex: 2}))
	pc, err = s.PendingConsolidations()
	require.NoError(t, err)
	require.Equal(t, primitives.ValidatorIndex(1), pc[0].SourceIndex)
	require.Equal(t, primitives.ValidatorIndex(2), pc[1].SourceIndex)
	pc, err = s2.PendingConsolidations()
	require.NoError(t, err)
	require.Equal(t, primitives.ValidatorIndex(1), pc[0].SourceIndex)
	require.Equal(t, primitives.ValidatorIndex(3), pc[1].SourceIndex)

	// Fails for versions older than electra
	s, err = state_native.InitializeFromProtoDeneb(&sila.BeaconStateDeneb{})
	require.NoError(t, err)
	require.ErrorContains(t, "not supported", s.AppendPendingConsolidation(&sila.PendingConsolidation{}))
}

func TestSetPendingConsolidations(t *testing.T) {
	s, err := state_native.InitializeFromProtoElectra(&sila.BeaconStateElectra{})
	require.NoError(t, err)
	num, err := s.NumPendingConsolidations()
	require.NoError(t, err)
	require.Equal(t, uint64(0), num)
	require.NoError(t, s.SetPendingConsolidations([]*sila.PendingConsolidation{{}, {}, {}}))
	num, err = s.NumPendingConsolidations()
	require.NoError(t, err)
	require.Equal(t, uint64(3), num)

	// Fails for versions older than electra
	s, err = state_native.InitializeFromProtoDeneb(&sila.BeaconStateDeneb{})
	require.NoError(t, err)
	require.ErrorContains(t, "not supported", s.SetPendingConsolidations([]*sila.PendingConsolidation{{}, {}, {}}))
}

func TestSetEarliestConsolidationEpoch(t *testing.T) {
	s, err := state_native.InitializeFromProtoElectra(&sila.BeaconStateElectra{})
	require.NoError(t, err)
	ece, err := s.EarliestConsolidationEpoch()
	require.NoError(t, err)
	require.Equal(t, primitives.Epoch(0), ece)
	require.NoError(t, s.SetEarliestConsolidationEpoch(10))
	ece, err = s.EarliestConsolidationEpoch()
	require.NoError(t, err)
	require.Equal(t, primitives.Epoch(10), ece)

	// Fails for versions older than electra
	s, err = state_native.InitializeFromProtoDeneb(&sila.BeaconStateDeneb{})
	require.NoError(t, err)
	require.ErrorContains(t, "not supported", s.SetEarliestConsolidationEpoch(10))
}

func TestSetConsolidationBalanceToConsume(t *testing.T) {
	s, err := state_native.InitializeFromProtoElectra(&sila.BeaconStateElectra{})
	require.NoError(t, err)
	require.NoError(t, s.SetConsolidationBalanceToConsume(10))
	cbtc, err := s.ConsolidationBalanceToConsume()
	require.NoError(t, err)
	require.Equal(t, primitives.Gwei(10), cbtc)

	// Fails for versions older than electra
	s, err = state_native.InitializeFromProtoDeneb(&sila.BeaconStateDeneb{})
	require.NoError(t, err)
	require.ErrorContains(t, "not supported", s.SetConsolidationBalanceToConsume(10))
}
