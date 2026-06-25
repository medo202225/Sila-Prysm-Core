package db

import "github.com/sila-chain/Sila-Prysm-Core/v7/beacon-chain/db/kv"

var _ Database = (*kv.Store)(nil)
