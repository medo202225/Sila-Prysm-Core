package blocks_test

import (
	"testing"

	"github.com/sila-chain/Sila-Prysm-Core/v7/beacon-chain/core/blocks"
	"github.com/sila-chain/Sila-Prysm-Core/v7/encoding/bytesutil"
	"github.com/sila-chain/Sila-Prysm-Core/v7/testing/assert"
)

func TestGenesisBlock_InitializedCorrectly(t *testing.T) {
	stateHash := bytesutil.PadTo([]byte{0}, 32)
	b1 := blocks.NewGenesisBlock(stateHash)

	assert.NotNil(t, b1.Block.ParentRoot, "Genesis block missing ParentHash field")
	assert.DeepEqual(t, b1.Block.StateRoot, stateHash, "Genesis block StateRootHash32 isn't initialized correctly")
}
