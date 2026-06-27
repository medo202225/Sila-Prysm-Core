package testing

import (
	"github.com/sila-chain/Sila-Consensus-Core/v7/consensus-types/blocks"
	"github.com/sila-chain/Sila-Consensus-Core/v7/consensus-types/interfaces"
	"github.com/sila-chain/Sila-Consensus-Core/v7/consensus-types/primitives"
	sila "github.com/sila-chain/Sila-Consensus-Core/v7/proto/sila/v1alpha1"
)

type blockMutator struct {
	Phase0    func(beaconBlock *sila.SignedBeaconBlock)
	Altair    func(beaconBlock *sila.SignedBeaconBlockAltair)
	Bellatrix func(beaconBlock *sila.SignedBeaconBlockBellatrix)
	Capella   func(beaconBlock *sila.SignedBeaconBlockCapella)
}

func (m blockMutator) apply(b interfaces.SignedBeaconBlock) (interfaces.SignedBeaconBlock, error) {
	pb, err := b.Proto()
	if err != nil {
		return nil, err
	}
	switch pbStruct := pb.(type) {
	case *sila.SignedBeaconBlock:
		m.Phase0(pbStruct)
	case *sila.SignedBeaconBlockAltair:
		m.Altair(pbStruct)
	case *sila.SignedBeaconBlockBellatrix:
		m.Bellatrix(pbStruct)
	case *sila.SignedBeaconBlockCapella:
		m.Capella(pbStruct)
	default:
		return nil, blocks.ErrUnsupportedSignedBeaconBlock
	}
	return blocks.NewSignedBeaconBlock(pb)
}

// SetBlockStateRoot modifies the block's state root.
func SetBlockStateRoot(b interfaces.SignedBeaconBlock, sr [32]byte) (interfaces.SignedBeaconBlock, error) {
	return blockMutator{
		Phase0:    func(bb *sila.SignedBeaconBlock) { bb.Block.StateRoot = sr[:] },
		Altair:    func(bb *sila.SignedBeaconBlockAltair) { bb.Block.StateRoot = sr[:] },
		Bellatrix: func(bb *sila.SignedBeaconBlockBellatrix) { bb.Block.StateRoot = sr[:] },
		Capella:   func(bb *sila.SignedBeaconBlockCapella) { bb.Block.StateRoot = sr[:] },
	}.apply(b)
}

// SetBlockParentRoot modifies the block's parent root.
func SetBlockParentRoot(b interfaces.SignedBeaconBlock, pr [32]byte) (interfaces.SignedBeaconBlock, error) {
	return blockMutator{
		Phase0:    func(bb *sila.SignedBeaconBlock) { bb.Block.ParentRoot = pr[:] },
		Altair:    func(bb *sila.SignedBeaconBlockAltair) { bb.Block.ParentRoot = pr[:] },
		Bellatrix: func(bb *sila.SignedBeaconBlockBellatrix) { bb.Block.ParentRoot = pr[:] },
		Capella:   func(bb *sila.SignedBeaconBlockCapella) { bb.Block.ParentRoot = pr[:] },
	}.apply(b)
}

// SetBlockSlot modifies the block's slot.
func SetBlockSlot(b interfaces.SignedBeaconBlock, s primitives.Slot) (interfaces.SignedBeaconBlock, error) {
	return blockMutator{
		Phase0:    func(bb *sila.SignedBeaconBlock) { bb.Block.Slot = s },
		Altair:    func(bb *sila.SignedBeaconBlockAltair) { bb.Block.Slot = s },
		Bellatrix: func(bb *sila.SignedBeaconBlockBellatrix) { bb.Block.Slot = s },
		Capella:   func(bb *sila.SignedBeaconBlockCapella) { bb.Block.Slot = s },
	}.apply(b)
}

// SetProposerIndex modifies the block's proposer index.
func SetProposerIndex(b interfaces.SignedBeaconBlock, idx primitives.ValidatorIndex) (interfaces.SignedBeaconBlock, error) {
	return blockMutator{
		Phase0:    func(bb *sila.SignedBeaconBlock) { bb.Block.ProposerIndex = idx },
		Altair:    func(bb *sila.SignedBeaconBlockAltair) { bb.Block.ProposerIndex = idx },
		Bellatrix: func(bb *sila.SignedBeaconBlockBellatrix) { bb.Block.ProposerIndex = idx },
		Capella:   func(bb *sila.SignedBeaconBlockCapella) { bb.Block.ProposerIndex = idx },
	}.apply(b)
}
