package genesis

import (
	"testing"

	"github.com/sila-chain/Sila-Consensus-Core/v7/config/params"
	"github.com/sila-chain/Sila-Consensus-Core/v7/genesis/internal/embedded"
	"github.com/sila-chain/Sila-Consensus-Core/v7/testing/require"
)

func TestEmbededGenesisDataMatchesMainnet(t *testing.T) {
	st, err := embedded.ByName(params.MainnetName)
	require.NoError(t, err)
	gvr := st.GenesisValidatorsRoot()

	data := embeddedGenesisData[params.MainnetName]
	require.DeepEqual(t, gvr, data.ValidatorsRoot[:])
	require.Equal(t, st.GenesisTime(), data.Time)
}
