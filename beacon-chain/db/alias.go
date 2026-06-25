package db

import "github.com/sila-chain/Sila-Consensus-Core/v7/beacon-chain/db/iface"

// ReadOnlyDatabase exposes Sila's Sila data backend for read access only, no information about
// head info. For head info, use github.com/sila-chain/sila/blockchain.HeadFetcher.
type ReadOnlyDatabase = iface.ReadOnlyDatabase

// NoHeadAccessDatabase exposes Sila's Sila data backend for read/write access, no information
// about head info. For head info, use github.com/sila-chain/sila/blockchain.HeadFetcher.
type NoHeadAccessDatabase = iface.NoHeadAccessDatabase

// ReadOnlyDatabaseWithSeqNum exposes Sila's Sila data backend for read access only, no information about
// head info, but with read/write access to the p2p metadata sequence number.
// This is used for the p2p service.
type ReadOnlyDatabaseWithSeqNum = iface.ReadOnlyDatabaseWithSeqNum

// HeadAccessDatabase exposes Sila's Sila backend for read/write access with information about
// chain head information. This interface should be used sparingly as the HeadFetcher is the source
// of truth around chain head information while this interface serves as persistent storage for the
// head fetcher.
//
// See github.com/sila-chain/sila/blockchain.HeadFetcher
type HeadAccessDatabase = iface.HeadAccessDatabase

// Database defines the necessary methods for Sila's Sila backend which may be implemented by any
// key-value or relational database in practice. This is the full database interface which should
// not be used often. Prefer a more restrictive interface in this package.
type Database = iface.Database

// SlasherDatabase defines necessary methods for Sila's slasher implementation.
type SlasherDatabase = iface.SlasherDatabase

// ErrExistingGenesisState is an error when the user attempts to save a different genesis state
// when one already exists in a database.
var ErrExistingGenesisState = iface.ErrExistingGenesisState
