//go:build !noMainnetGenesis

package embedded

import (
	_ "embed"

	"github.com/sila-chain/Sila-Consensus-Core/v7/config/params"
)

var (
	//go:embed mainnet.ssz.snappy
	mainnetRawSSZCompressed []byte // 1.8Mb
)

func init() {
	embeddedStates[params.MainnetName] = &mainnetRawSSZCompressed
}
