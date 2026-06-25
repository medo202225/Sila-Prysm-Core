### Changed

- Fix `SilaPayloadEnvelopesByRange` to only serve canonical payloads by walking the `ParentBlockHash` chain backward from a successor block.
- Add `ParentBlockHash` field to `BlindedSilaPayloadEnvelope` proto to enable the backward walk without loading full blocks.
- Add block hash indexing for `BlindedSilaPayloadEnvelope` in DB.
