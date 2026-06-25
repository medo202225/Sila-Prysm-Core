package endtoend

// This file contains the dependencies required for github.com/sila-chain/Sila/cmd/geth.
// Having these dependencies listed here helps go mod understand that these dependencies are
// necessary for end to end tests since we build go-ethereum binary for this test.
import (
	_ "github.com/sila-chain/Sila/accounts"          // Required for go-ethereum e2e.
	_ "github.com/sila-chain/Sila/accounts/keystore" // Required for go-ethereum e2e.
	_ "github.com/sila-chain/Sila/cmd/utils"         // Required for go-ethereum e2e.
	_ "github.com/sila-chain/Sila/common"            // Required for go-ethereum e2e.
	_ "github.com/sila-chain/Sila/console"           // Required for go-ethereum e2e.
	_ "github.com/sila-chain/Sila/eth"               // Required for go-ethereum e2e.
	_ "github.com/sila-chain/Sila/eth/downloader"    // Required for go-ethereum e2e.
	_ "github.com/sila-chain/Sila/ethclient"         // Required for go-ethereum e2e.
	_ "github.com/sila-chain/Sila/log"               // Required for go-ethereum e2e.
	_ "github.com/sila-chain/Sila/metrics"           // Required for go-ethereum e2e.
	_ "github.com/sila-chain/Sila/node"              // Required for go-ethereum e2e.
)
