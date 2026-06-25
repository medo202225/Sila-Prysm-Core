package blocks

import (
	"bytes"
	"context"
	"errors"

	"github.com/sila-chain/Sila-Consensus-Core/v7/beacon-chain/state"
	"github.com/sila-chain/Sila-Consensus-Core/v7/config/params"
	"github.com/sila-chain/Sila-Consensus-Core/v7/monitoring/tracing/trace"
	silapb "github.com/sila-chain/Sila-Consensus-Core/v7/proto/sila/v1alpha1"
)

// ProcessSilaExecutionDataInBlock is an operation performed on each
// beacon block to ensure the SILAEXEC data votes are processed
// into the beacon state.
//
// Official spec definition:
//
//	def process_sila_execution_data(state: BeaconState, body: BeaconBlockBody) -> None:
//	 state.sila_execution_data_votes.append(body.sila_execution_data)
//	 if state.sila_execution_data_votes.count(body.sila_execution_data) * 2 > EPOCHS_PER_SilaExecution_VOTING_PERIOD * SLOTS_PER_EPOCH:
//	     state.sila_execution_data = body.sila_execution_data
func ProcessSilaExecutionDataInBlock(ctx context.Context, beaconState state.BeaconState, silaexecData *silapb.SilaExecutionData) (state.BeaconState, error) {
	_, span := trace.StartSpan(ctx, "blocks.ProcessSilaExecutionDataInBlock")
	defer span.End()

	if beaconState == nil || beaconState.IsNil() {
		return nil, errors.New("nil state")
	}
	if err := beaconState.AppendSilaExecutionDataVotes(silaexecData); err != nil {
		return nil, err
	}
	hasSupport, err := SilaExecutionDataHasEnoughSupport(beaconState, silaexecData)
	if err != nil {
		return nil, err
	}
	if hasSupport {
		if err := beaconState.SetSilaExecutionData(silaexecData); err != nil {
			return nil, err
		}
	}
	return beaconState, nil
}

// AreSilaExecutionDataEqual checks equality between two silaexec data objects.
func AreSilaExecutionDataEqual(a, b *silapb.SilaExecutionData) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	return a.DepositCount == b.DepositCount &&
		bytes.Equal(a.BlockHash, b.BlockHash) &&
		bytes.Equal(a.DepositRoot, b.DepositRoot)
}

// SilaExecutionDataHasEnoughSupport returns true when the given silaExecutionData has more than 50% votes in the
// silaexec voting period. A vote is cast by including silaExecutionData in a block and part of state processing
// appends silaExecutionData to the state in the SilaExecutionDataVotes list. Iterating through this list checks the
// votes to see if they match the silaExecutionData.
func SilaExecutionDataHasEnoughSupport(beaconState state.ReadOnlyBeaconState, data *silapb.SilaExecutionData) (bool, error) {
	voteCount := uint64(0)

	for _, vote := range beaconState.SilaExecutionDataVotes() {
		if AreSilaExecutionDataEqual(vote, data) {
			voteCount++
		}
	}

	// If 50+% majority converged on the same silaExecutionData, then it has enough support to update the
	// state.
	support := params.BeaconConfig().SlotsPerEpoch.Mul(uint64(params.BeaconConfig().EpochsPerSilaExecutionVotingPeriod))
	return voteCount*2 > uint64(support), nil
}
