package kzg

import (
	"testing"

	"github.com/sila-chain/Sila-Prysm-Core/v7/testing/require"
)

func TestStart(t *testing.T) {
	require.NoError(t, Start())
	require.NotNil(t, kzgContext)
}
