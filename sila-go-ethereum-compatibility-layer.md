# Sila go-ethereum Compatibility Layer

This document defines how Sila-Prysm treats go-ethereum dependencies during the Sila transition.

## Scope

This document is a compatibility-layer plan only.

It does not remove go-ethereum imports.
It does not change execution or consensus logic.
It does not replace Ethereum-derived types.
It does not rename compatibility APIs.

## Current status

Sila-Prysm still uses selected go-ethereum packages as compatibility substrate.

These packages provide mature, tested primitives required by execution integration, RPC handling, cryptographic helpers, ENR/enode handling, hex utilities, and test infrastructure.

## Preserved compatibility

The following categories remain compatibility substrate:

- common.Address / common.Hash
- common/hexutil
- rpc helpers
- execution payload related types
- core/types
- ethclient compatibility helpers
- p2p/enode
- p2p/enr
- kzg4844 helpers
- test fixtures that depend on go-ethereum behavior

## Sila usage rule

Sila-Prysm may use go-ethereum packages internally only as compatibility infrastructure.

User-facing identity, logs, docs, and Sila RPC paths must use Sila terminology.

## Do not remove yet

The following must remain until Sila-native replacements exist:

- github.com/ethereum/go-ethereum/common
- github.com/ethereum/go-ethereum/common/hexutil
- github.com/ethereum/go-ethereum/rpc
- github.com/ethereum/go-ethereum/core/types
- github.com/ethereum/go-ethereum/ethclient
- github.com/ethereum/go-ethereum/p2p/enode
- github.com/ethereum/go-ethereum/p2p/enr
- github.com/ethereum/go-ethereum/crypto/kzg4844
- ethereum.NotFound
- ethereum.CallMsg
- ethereum.FilterQuery
- ethereum.Subscription

## Required before removal

1. Sila common type layer
2. Sila hex utility layer
3. Sila RPC abstraction layer
4. Sila execution payload type mapping
5. Sila ENR/enode compatibility abstraction
6. Sila KZG compatibility abstraction
7. Full regression tests proving identical behavior
8. End-to-end SilaChain + Sila-Prysm validation
