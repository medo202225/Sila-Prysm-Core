package verification

import (
	"os"
	"testing"

	"github.com/sila-chain/Sila-Consensus-Core/v7/beacon-chain/blockchain/kzg"
)

func TestMain(t *testing.M) {
	if err := kzg.Start(); err != nil {
		os.Exit(1)
	}
	t.Run()
}
