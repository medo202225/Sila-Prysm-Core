package attestations

import (
	"github.com/sila-chain/Sila-Consensus-Core/v7/beacon-chain/operations/attestations/kv"
)

var _ Pool = (*kv.AttCaches)(nil)
