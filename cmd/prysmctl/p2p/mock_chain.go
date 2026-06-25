package p2p

import (
	"time"

	"github.com/sila-chain/Sila-Prysm-Core/v7/beacon-chain/forkchoice"
	"github.com/sila-chain/Sila-Prysm-Core/v7/consensus-types/primitives"
	ethpb "github.com/sila-chain/Sila-Prysm-Core/v7/proto/prysm/v1alpha1"
	"github.com/sila-chain/Sila-Prysm-Core/v7/time/slots"
)

type mockChain struct {
	currentFork     *ethpb.Fork
	genesisValsRoot [32]byte
	genesisTime     time.Time
}

func (m *mockChain) ForkChoicer() forkchoice.ForkChoicer {
	return nil
}

func (m *mockChain) CurrentFork() *ethpb.Fork {
	return m.currentFork
}

func (m *mockChain) GenesisValidatorsRoot() [32]byte {
	return m.genesisValsRoot
}

func (m *mockChain) GenesisTime() time.Time {
	return m.genesisTime
}

func (m *mockChain) CurrentSlot() primitives.Slot {
	return slots.CurrentSlot(m.genesisTime)
}
