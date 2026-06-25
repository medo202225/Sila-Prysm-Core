package kv

import (
	"testing"

	testpb "github.com/sila-chain/Sila-Consensus-Core/v7/proto/testing"
	"github.com/sila-chain/Sila-Consensus-Core/v7/testing/require"
)

func Test_encode_handlesNilFromFunction(t *testing.T) {
	foo := func() *testpb.Puzzle {
		return nil
	}
	_, err := encode(t.Context(), foo())
	require.ErrorContains(t, "cannot encode nil message", err)
}
