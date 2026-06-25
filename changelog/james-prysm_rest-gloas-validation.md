### Added

- SSZ support for GET and POST of sila payload envelope and envelope contents.
- `broadcast_validation` query parameter on POST sila payload envelope.
- Spec-wire `WireBlindedSilaPayloadEnvelope` types and `Sila-Payload-Blinded`
  header for the stateful publish path (beacon-APIs #580).
- `202` response on POST sila payload envelope when the envelope is broadcast
  but fails database integration (beacon-APIs #580).
- `ProduceBlockV4` returns only the beacon block when the produced block uses an
  external builder bid, regardless of `include_payload` (beacon-APIs #580).

### Changed

- `GET /sila/v1/validator/sila_payload_envelope/{slot}` →
  `GET /sila/v1/validator/sila_payload_envelopes/{slot}/{beacon_block_root}`;
  the response is the spec-wire `BlindedSilaPayloadEnvelope` (payload replaced
  by `payload_root`, HTR equivalent to the full envelope). Returns only
  `Eth-Consensus-Version` (beacon-APIs #580 / PR #10).
- Stateful self-build now works end to end: the validator client fetches the blinded
  envelope from the BN, signs its (HTR-equivalent) root, and publishes the
  `SignedBlindedSilaPayloadEnvelope`.
- `POST /sila/v1/beacon/sila_payload_envelopes` body shape is now selected by
  the required `Sila-Payload-Blinded` request header:
  - `true` → `SignedBlindedSilaPayloadEnvelope` (stateful — BN reconstructs
    the full envelope from its cache).
  - `false` → `SignedSilaPayloadEnvelopeContents` (stateless — body carries
    blobs and KZG proofs).
  Replaces the prior SSZ-lead-offset / JSON wrapper-key probe.
- Pluralized gloas sila payload endpoint paths to match the REST naming
  convention (beacon-APIs #613): `POST /sila/v1/beacon/sila_payload_bid` →
  `/sila/v1/beacon/sila_payload_bids`, and the sila payload envelope
  paths use `sila_payload_envelopes`.