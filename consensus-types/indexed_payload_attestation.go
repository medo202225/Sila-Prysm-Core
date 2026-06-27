package consensus_types

import (
	"github.com/sila-chain/Sila-Consensus-Core/v7/consensus-types/primitives"
	sila "github.com/sila-chain/Sila-Consensus-Core/v7/proto/sila/v1alpha1"
)

type IndexedPayloadAttestation struct {
	AttestingIndices []primitives.ValidatorIndex
	Data             *sila.PayloadAttestationData
	Signature        []byte
}

// GetAttestingIndices returns the attesting indices or nil when the receiver is nil.
func (x *IndexedPayloadAttestation) GetAttestingIndices() []primitives.ValidatorIndex {
	if x == nil {
		return nil
	}
	return x.AttestingIndices
}

// GetData returns the attestation data or nil when the receiver is nil.
func (x *IndexedPayloadAttestation) GetData() *sila.PayloadAttestationData {
	if x == nil {
		return nil
	}
	return x.Data
}

// GetSignature returns the signature bytes or nil when the receiver is nil.
func (x *IndexedPayloadAttestation) GetSignature() []byte {
	if x == nil {
		return nil
	}
	return x.Signature
}
