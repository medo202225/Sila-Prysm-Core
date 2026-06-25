### Fixed

- Preserve the full signed sila payload envelope in pubsub validation by setting `msg.ValidatorData = signedEnvelope` after successful validation.
