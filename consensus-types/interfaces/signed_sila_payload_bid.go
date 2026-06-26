package interfaces

import (
	field_params "github.com/sila-chain/Sila-Consensus-Core/v7/config/fieldparams"
	"github.com/sila-chain/Sila-Consensus-Core/v7/consensus-types/primitives"
)

type ROSignedSilaPayloadBid interface {
	Bid() (ROSilaPayloadBid, error)
	Signature() [field_params.BLSSignatureLength]byte
	SigningRoot([]byte) ([32]byte, error)
	IsNil() bool
}

type ROSilaPayloadBid interface {
	ParentBlockHash() [32]byte
	ParentBlockRoot() [32]byte
	PrevRandao() [32]byte
	BlockHash() [32]byte
	GasLimit() uint64
	BuilderIndex() primitives.BuilderIndex
	Slot() primitives.Slot
	Value() primitives.Gwei
	SilaPayment() primitives.Gwei
	BlobKzgCommitments() [][]byte
	BlobKzgCommitmentCount() uint64
	FeeRecipient() [20]byte
	SilaRequestsRoot() [32]byte
	IsNil() bool
}
