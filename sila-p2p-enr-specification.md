# Sila P2P / ENR Specification

This document defines the current Sila P2P and ENR identity boundaries used by Sila-Prysm.

## Scope

This specification documents Sila P2P identity and compatibility boundaries.

It does not change networking logic.
It does not change ENR behavior.
It does not remove Ethereum-derived compatibility dependencies.
It does not rename protocol identifiers required for interoperability.

## Current Sila P2P identity

Sila-Prysm identifies the P2P layer as:

- Sila P2P
- Sila consensus networking
- Sila beacon node P2P

## Preserved compatibility

The following remain compatibility substrate:

- discv5 behavior
- libp2p behavior
- ENR record structure
- fork digest behavior
- gossip message-id behavior
- SSZ/snappy gossip encoding
- peer status compatibility
- subnet and custody-group compatibility where required

## ENR compatibility

Sila-Prysm may continue to use compatibility-derived ENR keys and structures until Sila-native ENR conventions exist.

Do not remove or rename ENR identifiers until a tested Sila replacement exists.

## Message ID compatibility

Sila-Prysm currently preserves message-id semantics compatible with consensus networking behavior.

Do not replace message-id computation until Sila-native test vectors exist.

## Do not remove yet

The following must remain until Sila-native replacements exist:

- go-ethereum p2p/enode imports
- go-ethereum p2p/enr imports
- consensus-specs P2P references
- ENR compatibility keys
- message-id compatibility rules
- gossip/subnet compatibility behavior

## Required before removing compatibility substrate

1. Sila ENR key specification
2. Sila fork digest compatibility statement
3. Sila message-id test vectors
4. Sila gossip topic compatibility statement
5. Sila subnet compatibility statement
6. Sila peer-status compatibility tests
7. Sila devnet P2P interoperability validation
