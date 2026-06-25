package state_native

import (
	"github.com/sila-chain/Sila-Consensus-Core/v7/beacon-chain/state/state-native/types"
	"github.com/sila-chain/Sila-Consensus-Core/v7/beacon-chain/state/stateutil"
	silapb "github.com/sila-chain/Sila-Consensus-Core/v7/proto/sila/v1alpha1"
)

// SetSilaExecutionData for the beacon state.
func (b *BeaconState) SetSilaExecutionData(val *silapb.SilaExecutionData) error {
	b.lock.Lock()
	defer b.lock.Unlock()

	b.silaexecData = val
	b.markFieldAsDirty(types.SilaExecutionData)
	return nil
}

// SetSilaExecutionDataVotes for the beacon state. Updates the entire
// list to a new value by overwriting the previous one.
func (b *BeaconState) SetSilaExecutionDataVotes(val []*silapb.SilaExecutionData) error {
	b.lock.Lock()
	defer b.lock.Unlock()

	b.sharedFieldReferences[types.SilaExecutionDataVotes].MinusRef()
	b.sharedFieldReferences[types.SilaExecutionDataVotes] = stateutil.NewRef(1)

	b.silaExecutionDataVotes = val
	b.markFieldAsDirty(types.SilaExecutionDataVotes)
	b.rebuildTrie[types.SilaExecutionDataVotes] = true
	return nil
}

// SetSilaExecutionDepositIndex for the beacon state.
func (b *BeaconState) SetSilaExecutionDepositIndex(val uint64) error {
	b.lock.Lock()
	defer b.lock.Unlock()

	b.silaExecutionDepositIndex = val
	b.markFieldAsDirty(types.SilaExecutionDepositIndex)
	return nil
}

// AppendSilaExecutionDataVotes for the beacon state. Appends the new value
// to the end of list.
func (b *BeaconState) AppendSilaExecutionDataVotes(val *silapb.SilaExecutionData) error {
	b.lock.Lock()
	defer b.lock.Unlock()

	votes := b.silaExecutionDataVotes
	if b.sharedFieldReferences[types.SilaExecutionDataVotes].Refs() > 1 {
		// Copy elements in underlying array by reference.
		votes = make([]*silapb.SilaExecutionData, 0, len(b.silaExecutionDataVotes)+1)
		votes = append(votes, b.silaExecutionDataVotes...)
		b.sharedFieldReferences[types.SilaExecutionDataVotes].MinusRef()
		b.sharedFieldReferences[types.SilaExecutionDataVotes] = stateutil.NewRef(1)
	}

	b.silaExecutionDataVotes = append(votes, val)
	b.markFieldAsDirty(types.SilaExecutionDataVotes)
	b.addDirtyIndices(types.SilaExecutionDataVotes, []uint64{uint64(len(b.silaExecutionDataVotes) - 1)})
	return nil
}
