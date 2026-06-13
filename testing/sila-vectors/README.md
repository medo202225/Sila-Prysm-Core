# Sila Native Test Vectors

This directory is reserved for Sila-native test vectors.

Rules:
- Do not remove Ethereum-derived vectors until Sila vectors exist and pass.
- Do not add mock vectors.
- Do not add placeholder values as valid vectors.
- Every Sila vector must be generated from real SilaChain / Sila-Prysm behavior.
- Every Sila vector must have a regression test before replacing compatibility behavior.

Initial vector groups:
- consensus
- execution-api
- p2p
- cryptography
