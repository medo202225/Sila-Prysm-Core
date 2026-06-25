package attestations

import (
	"github.com/sila-chain/Sila-Prysm-Core/v7/beacon-chain/operations/attestations/kv"
)

var _ Pool = (*kv.AttCaches)(nil)
