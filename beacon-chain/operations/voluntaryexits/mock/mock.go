package mock

import (
	"github.com/sila-chain/Sila-Consensus-Core/v7/beacon-chain/state"
	"github.com/sila-chain/Sila-Consensus-Core/v7/consensus-types/primitives"
	sila "github.com/sila-chain/Sila-Consensus-Core/v7/proto/sila/v1alpha1"
)

// PoolMock is a fake implementation of PoolManager.
type PoolMock struct {
	Exits []*sila.SignedVoluntaryExit
}

// PendingExits --
func (m *PoolMock) PendingExits() ([]*sila.SignedVoluntaryExit, error) {
	return m.Exits, nil
}

// ExitsForInclusion --
func (m *PoolMock) ExitsForInclusion(_ state.ReadOnlyBeaconState, _ primitives.Slot) ([]*sila.SignedVoluntaryExit, error) {
	return m.Exits, nil
}

// InsertVoluntaryExit --
func (m *PoolMock) InsertVoluntaryExit(exit *sila.SignedVoluntaryExit) {
	m.Exits = append(m.Exits, exit)
}

// MarkIncluded --
func (*PoolMock) MarkIncluded(_ *sila.SignedVoluntaryExit) {
	panic("implement me") // lint:nopanic -- Mock / test code.
}
