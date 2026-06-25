### Fixed

- Skip peers returning empty `SilaPayloadEnvelopesByRange` responses when the block bids prove payloads exist; downscore only when the same peer served both the blocks and the empty response.
