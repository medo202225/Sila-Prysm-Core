package validator

import (
	"github.com/sila-chain/Sila-Prysm-Core/v7/beacon-chain/blockchain"
	"github.com/sila-chain/Sila-Prysm-Core/v7/beacon-chain/db"
	"github.com/sila-chain/Sila-Prysm-Core/v7/beacon-chain/rpc/core"
	"github.com/sila-chain/Sila-Prysm-Core/v7/beacon-chain/rpc/lookup"
)

type Server struct {
	BeaconDB            db.ReadOnlyDatabase
	Stater              lookup.Stater
	CanonicalFetcher    blockchain.CanonicalFetcher
	FinalizationFetcher blockchain.FinalizationFetcher
	ChainInfoFetcher    blockchain.ChainInfoFetcher
	CoreService         *core.Service
}
