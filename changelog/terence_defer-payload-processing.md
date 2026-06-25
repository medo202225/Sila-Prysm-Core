### Changed

- Implement defer payload processing spec change (consensus-specs#5094): move sila payload state mutations to `process_parent_sila_payload` in the next block, make `process_sila_payload` pure verification, add `sila_requests_root` to `SilaPayloadBid`, remove `state_root` from `SilaPayloadEnvelope`, and add `parent_sila_requests` to `BeaconBlockBody`.
