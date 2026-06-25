package interfaces

import (
	"testing"

	"github.com/sila-chain/Sila-Prysm-Core/v7/runtime/version"
	"github.com/sila-chain/Sila-Prysm-Core/v7/testing/require"
	"github.com/pkg/errors"
)

func TestNewInvalidCastError(t *testing.T) {
	err := NewInvalidCastError(version.Phase0, version.Electra)
	require.Equal(t, true, errors.Is(err, ErrInvalidCast))
}
