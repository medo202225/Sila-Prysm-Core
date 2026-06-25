package cache

import (
	"github.com/sila-chain/Sila-Prysm-Core/v7/consensus-types/primitives"
)

// ProposerIndices defines the cached struct for proposer indices.
type ProposerIndices struct {
	BlockRoot       [32]byte
	ProposerIndices []primitives.ValidatorIndex
}
