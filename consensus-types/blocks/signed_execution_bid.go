package blocks

import (
	"github.com/sila-chain/Sila-Consensus-Core/v7/beacon-chain/core/signing"
	consensus_types "github.com/sila-chain/Sila-Consensus-Core/v7/consensus-types"
	"github.com/sila-chain/Sila-Consensus-Core/v7/consensus-types/interfaces"
	"github.com/sila-chain/Sila-Consensus-Core/v7/consensus-types/primitives"
	"github.com/sila-chain/Sila-Consensus-Core/v7/encoding/bytesutil"
	silapb "github.com/sila-chain/Sila-Consensus-Core/v7/proto/sila/v1alpha1"
)

// signedSilaPayloadBid wraps the protobuf signed sila payload bid
// and implements the ROSignedSilaPayloadBid interface.
type signedSilaPayloadBid struct {
	bid *silapb.SignedSilaPayloadBid
}

// silaPayloadBidGloas wraps the protobuf sila payload bid for Gloas fork
// and implements the ROSilaPayloadBidGloas interface.
type silaPayloadBidGloas struct {
	payload *silapb.SilaPayloadBid
}

// IsNil checks if the signed sila payload bid is nil or invalid.
func (s signedSilaPayloadBid) IsNil() bool {
	if s.bid == nil {
		return true
	}

	if _, err := WrappedROSilaPayloadBid(s.bid.Message); err != nil {
		return true
	}

	return len(s.bid.Signature) != 96
}

// IsNil checks if the sila payload bid is nil or has invalid fields.
func (h silaPayloadBidGloas) IsNil() bool {
	if h.payload == nil {
		return true
	}

	if len(h.payload.ParentBlockHash) != 32 ||
		len(h.payload.ParentBlockRoot) != 32 ||
		len(h.payload.BlockHash) != 32 ||
		len(h.payload.PrevRandao) != 32 ||
		len(h.payload.FeeRecipient) != 20 {
		return true
	}

	for _, commitment := range h.payload.BlobKzgCommitments {
		if len(commitment) != 48 {
			return true
		}
	}

	return false
}

// WrappedROSignedSilaPayloadBid creates a new read-only signed sila payload bid
// wrapper from the given protobuf message.
func WrappedROSignedSilaPayloadBid(pb *silapb.SignedSilaPayloadBid) (interfaces.ROSignedSilaPayloadBid, error) {
	wrapper := signedSilaPayloadBid{bid: pb}
	if wrapper.IsNil() {
		return nil, consensus_types.ErrNilObjectWrapped
	}
	return wrapper, nil
}

// WrappedROSilaPayloadBid creates a new read-only sila payload bid
// wrapper for the Gloas fork from the given protobuf message.
func WrappedROSilaPayloadBid(pb *silapb.SilaPayloadBid) (interfaces.ROSilaPayloadBid, error) {
	wrapper := silaPayloadBidGloas{payload: pb}
	if wrapper.IsNil() {
		return nil, consensus_types.ErrNilObjectWrapped
	}
	return wrapper, nil
}

// Bid returns the sila payload bid as a read-only interface.
func (s signedSilaPayloadBid) Bid() (interfaces.ROSilaPayloadBid, error) {
	return WrappedROSilaPayloadBid(s.bid.Message)
}

// SigningRoot computes the signing root for the sila payload bid with the given domain.
func (s signedSilaPayloadBid) SigningRoot(domain []byte) ([32]byte, error) {
	return signing.ComputeSigningRoot(s.bid.Message, domain)
}

// Signature returns the BLS signature as a 96-byte array.
func (s signedSilaPayloadBid) Signature() [96]byte {
	return [96]byte(s.bid.Signature)
}

// ParentBlockHash returns the hash of the parent execution block.
func (h silaPayloadBidGloas) ParentBlockHash() [32]byte {
	return [32]byte(h.payload.ParentBlockHash)
}

// ParentBlockRoot returns the beacon block root of the parent block.
func (h silaPayloadBidGloas) ParentBlockRoot() [32]byte {
	return [32]byte(h.payload.ParentBlockRoot)
}

// PrevRandao returns the previous randao value for the execution block.
func (h silaPayloadBidGloas) PrevRandao() [32]byte {
	return [32]byte(h.payload.PrevRandao)
}

// BlockHash returns the hash of the execution block.
func (h silaPayloadBidGloas) BlockHash() [32]byte {
	return [32]byte(h.payload.BlockHash)
}

// GasLimit returns the gas limit for the execution block.
func (h silaPayloadBidGloas) GasLimit() uint64 {
	return h.payload.GasLimit
}

// BuilderIndex returns the builder index of the builder who created this bid.
func (h silaPayloadBidGloas) BuilderIndex() primitives.BuilderIndex {
	return h.payload.BuilderIndex
}

// Slot returns the beacon chain slot for which this bid was created.
func (h silaPayloadBidGloas) Slot() primitives.Slot {
	return h.payload.Slot
}

// Value returns the payment value offered by the builder in Gwei.
func (h silaPayloadBidGloas) Value() primitives.Gwei {
	return primitives.Gwei(h.payload.Value)
}

// ExecutionPayment returns the execution payment offered by the builder.
func (h silaPayloadBidGloas) ExecutionPayment() primitives.Gwei {
	return primitives.Gwei(h.payload.ExecutionPayment)
}

// BlobKzgCommitments returns the KZG commitments for blobs.
func (h silaPayloadBidGloas) BlobKzgCommitments() [][]byte {
	return bytesutil.SafeCopy2dBytes(h.payload.BlobKzgCommitments)
}

// BlobKzgCommitmentCount returns the number of blob KZG commitments.
func (h silaPayloadBidGloas) BlobKzgCommitmentCount() uint64 {
	return uint64(len(h.payload.BlobKzgCommitments))
}

// FeeRecipient returns the execution address that will receive the builder payment.
func (h silaPayloadBidGloas) FeeRecipient() [20]byte {
	return [20]byte(h.payload.FeeRecipient)
}

// SilaRequestsRoot returns the hash tree root of the sila requests.
func (h silaPayloadBidGloas) SilaRequestsRoot() [32]byte {
	if h.payload == nil || len(h.payload.SilaRequestsRoot) < 32 {
		return [32]byte{}
	}
	return [32]byte(h.payload.SilaRequestsRoot)
}
