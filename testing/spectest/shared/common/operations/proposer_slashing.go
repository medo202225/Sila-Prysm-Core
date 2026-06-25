package operations

import (
	"context"
	"testing"

	"github.com/sila-chain/Sila-Consensus-Core/v7/beacon-chain/core/blocks"
	v "github.com/sila-chain/Sila-Consensus-Core/v7/beacon-chain/core/validators"
	"github.com/sila-chain/Sila-Consensus-Core/v7/beacon-chain/state"
	"github.com/sila-chain/Sila-Consensus-Core/v7/consensus-types/interfaces"
)

func RunProposerSlashingTest(t *testing.T, config string, fork string, block blockWithSSZObject, sszToState SSZToState) {
	runSlashingTest(t, config, fork, "proposer_slashing", block, sszToState, func(ctx context.Context, s state.BeaconState, b interfaces.ReadOnlySignedBeaconBlock) (state.BeaconState, error) {
		return blocks.ProcessProposerSlashings(ctx, s, b.Block().Body().ProposerSlashings(), v.ExitInformation(s))
	})
}
