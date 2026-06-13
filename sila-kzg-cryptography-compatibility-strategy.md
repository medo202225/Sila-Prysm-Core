# Sila KZG / Cryptography Compatibility Strategy

This document defines the Sila strategy for KZG, BLS, and cryptography compatibility in Sila-Prysm.

## Scope

This document is a compatibility strategy only.

It does not replace cryptography libraries.
It does not change consensus logic.
It does not change signature verification.
It does not change KZG proof verification.
It does not change trusted setup behavior.

## Current status

Sila-Prysm preserves cryptography compatibility for consensus safety.

The following remain compatibility substrate:

- BLS signature behavior
- SSZ hash tree root behavior
- KZG commitment behavior
- KZG proof verification behavior
- trusted setup loading behavior
- c-kzg-4844 bindings
- go-ethereum kzg4844 helpers
- consensus-specs cryptography references

## Sila usage rule

Sila-Prysm may use Ethereum-derived cryptography libraries internally only as compatibility infrastructure.

User-facing identity and documentation should use Sila terminology, while protocol-critical cryptography behavior must remain unchanged until Sila-native replacements are proven equivalent.

## Do not remove yet

The following must remain until Sila-native replacements exist:

- github.com/ethereum/c-kzg-4844
- github.com/ethereum/go-ethereum/crypto/kzg4844
- BLS compatibility behavior
- KZG trusted setup compatibility
- KZG proof verification compatibility
- SSZ hash tree root compatibility
- consensus-specs cryptography references
- Ethereum-derived cryptography test vectors

## Required before removal

1. Sila KZG compatibility specification
2. Sila BLS compatibility specification
3. Sila trusted setup policy
4. Sila cryptography test vectors
5. Regression tests proving identical BLS behavior
6. Regression tests proving identical KZG behavior
7. End-to-end SilaChain + Sila-Prysm blob and proof validation
8. Security review before replacing any cryptography substrate
