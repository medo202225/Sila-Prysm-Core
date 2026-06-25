package sync

type silaPayloadEnvelopeRPCResult string

const (
	silaPayloadEnvelopeRPCResultServed              silaPayloadEnvelopeRPCResult = "served"
	silaPayloadEnvelopeRPCResultInvalid             silaPayloadEnvelopeRPCResult = "invalid"
	silaPayloadEnvelopeRPCResultRateLimited         silaPayloadEnvelopeRPCResult = "rate_limited"
	silaPayloadEnvelopeRPCResultResourceUnavailable silaPayloadEnvelopeRPCResult = "resource_unavailable"
	silaPayloadEnvelopeRPCResultError               silaPayloadEnvelopeRPCResult = "error"
)
