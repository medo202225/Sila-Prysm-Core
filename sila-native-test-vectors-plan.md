# Sila Native Test Vectors Plan

This document defines the plan for creating Sila-native test vectors before removing any compatibility-derived behavior.

## Scope

This document is a test-vector plan only.

It does not change consensus logic.
It does not change execution logic.
It does not remove Ethereum-derived test vectors.
It does not replace compatibility dependencies.

## Purpose

Sila-native test vectors are required before replacing compatibility substrate with Sila-native implementations.

## Required vector groups

### Consensus vectors

- Sila genesis state vectors
- Sila fork schedule vectors
- Sila state transition vectors
- Sila block processing vectors
- Sila attestation processing vectors
- Sila finality and justification vectors
- Sila slashing protection vectors

### Execution API vectors

- silaEngine_newPayload vectors
- silaEngine_forkchoiceUpdated vectors
- silaEngine_getPayload vectors
- silaEngine_exchangeTransitionConfiguration vectors
- sila_call vectors
- sila_getCode vectors

### P2P vectors

- Sila fork digest vectors
- Sila ENR vectors
- Sila message-id vectors
- Sila gossip topic vectors
- Sila peer status vectors
- Sila subnet vectors

### Cryptography vectors

- Sila BLS compatibility vectors
- Sila KZG commitment vectors
- Sila KZG proof verification vectors
- Sila SSZ hash tree root vectors

## Removal rule

No Ethereum-derived compatibility behavior may be removed until the corresponding Sila-native vector group exists and passes regression tests.

## Required before implementation replacement

1. Generate Sila-native vectors
2. Add regression tests
3. Compare against compatibility behavior
4. Run SilaChain + Sila-Prysm end-to-end validation
5. Document equivalence
6. Only then replace compatibility substrate
