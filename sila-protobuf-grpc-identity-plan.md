# Sila protobuf / gRPC Identity Plan

This document defines the Sila plan for protobuf and gRPC identity in Sila-Prysm.

## Scope

This document is a compatibility plan only.

It does not change protobuf definitions.
It does not rename generated packages.
It does not change gRPC service names.
It does not break external validator, beacon, or tooling compatibility.

## Current status

Some protobuf and gRPC package identifiers remain Ethereum-derived because they are part of external API compatibility.

These identifiers must remain until Sila-native protobuf and gRPC replacements are fully designed, generated, tested, and migrated.

## Preserved compatibility

The following remain compatibility substrate:

- generated protobuf package names
- gRPC service identifiers
- REST/gRPC API compatibility behavior
- external validator client API compatibility
- beacon API compatibility
- test expectations that depend on existing service names

## Sila identity layer

Sila-Prysm user-facing documentation and logs should use Sila terminology.

Protocol-facing package identifiers may remain compatibility-derived until a full Sila migration is available.

## Do not remove yet

The following must not be renamed directly:

- ethereum.eth.v1alpha1 package/service identifiers
- generated protobuf package names
- generated gRPC service names
- compatibility API paths used by external tools
- tests that assert external service names

## Required before migration

1. Sila protobuf namespace design
2. Sila gRPC service naming plan
3. Generated Sila protobuf code
4. Backward compatibility strategy
5. Migration tests for validators and clients
6. Sila REST/gRPC API compatibility matrix
7. Deprecation plan for compatibility-derived identifiers
