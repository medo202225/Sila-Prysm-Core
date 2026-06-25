package api

import (
	"testing"

	"github.com/sila-chain/Sila-Prysm-Core/v7/testing/require"
)

func TestGenerateRandomHexString(t *testing.T) {
	token, err := GenerateRandomHexString()
	require.NoError(t, err)
	require.NoError(t, ValidateAuthToken(token))
}
