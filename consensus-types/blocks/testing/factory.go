package testing

import (
	"github.com/pkg/errors"
	"github.com/sila-chain/Sila-Consensus-Core/v7/consensus-types/blocks"
	"github.com/sila-chain/Sila-Consensus-Core/v7/consensus-types/interfaces"
	sila "github.com/sila-chain/Sila-Consensus-Core/v7/proto/sila/v1alpha1"
)

// NewSignedBeaconBlockFromGeneric creates a signed beacon block
// from a protobuf generic signed beacon block.
func NewSignedBeaconBlockFromGeneric(gb *sila.GenericSignedBeaconBlock) (interfaces.ReadOnlySignedBeaconBlock, error) {
	if gb == nil {
		return nil, blocks.ErrNilObject
	}
	switch bb := gb.Block.(type) {
	case *sila.GenericSignedBeaconBlock_Phase0:
		return blocks.NewSignedBeaconBlock(bb.Phase0)
	case *sila.GenericSignedBeaconBlock_Altair:
		return blocks.NewSignedBeaconBlock(bb.Altair)
	case *sila.GenericSignedBeaconBlock_Bellatrix:
		return blocks.NewSignedBeaconBlock(bb.Bellatrix)
	case *sila.GenericSignedBeaconBlock_BlindedBellatrix:
		return blocks.NewSignedBeaconBlock(bb.BlindedBellatrix)
	case *sila.GenericSignedBeaconBlock_Capella:
		return blocks.NewSignedBeaconBlock(bb.Capella)
	case *sila.GenericSignedBeaconBlock_BlindedCapella:
		return blocks.NewSignedBeaconBlock(bb.BlindedCapella)
	case *sila.GenericSignedBeaconBlock_Deneb:
		return blocks.NewSignedBeaconBlock(bb.Deneb.Block)
	case *sila.GenericSignedBeaconBlock_BlindedDeneb:
		return blocks.NewSignedBeaconBlock(bb.BlindedDeneb)
	case *sila.GenericSignedBeaconBlock_Electra:
		return blocks.NewSignedBeaconBlock(bb.Electra.Block)
		// Generic Signed Beacon Block Deneb can't be used here as it is not a block, but block content with blobs
	case *sila.GenericSignedBeaconBlock_Fulu:
		return blocks.NewSignedBeaconBlock(bb.Fulu.Block)
		// Generic Signed Beacon Block Deneb can't be used here as it is not a block, but block content with blobs
	default:
		return nil, errors.Wrapf(blocks.ErrUnsupportedSignedBeaconBlock, "unable to create block from type %T", gb)
	}
}
