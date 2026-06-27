package blocks

import (
	"fmt"

	"github.com/pkg/errors"
	consensus_types "github.com/sila-chain/Sila-Consensus-Core/v7/consensus-types"
	"github.com/sila-chain/Sila-Consensus-Core/v7/encoding/bytesutil"
	sila "github.com/sila-chain/Sila-Consensus-Core/v7/proto/sila/v1alpha1"
	silaenginev1 "github.com/sila-chain/Sila-Consensus-Core/v7/proto/silaengine/v1"
	"github.com/sila-chain/Sila-Consensus-Core/v7/runtime/version"
	"google.golang.org/protobuf/proto"
)

// Proto converts the signed beacon block to a protobuf object.
func (b *SignedBeaconBlock) Proto() (proto.Message, error) { // nolint:gocognit
	if b == nil {
		return nil, errNilBlock
	}

	blockMessage, err := b.block.Proto()
	if err != nil {
		return nil, err
	}

	switch b.version {
	case version.Phase0:
		var block *sila.BeaconBlock
		if blockMessage != nil {
			var ok bool
			block, ok = blockMessage.(*sila.BeaconBlock)
			if !ok {
				return nil, errIncorrectBlockVersion
			}
		}
		return &sila.SignedBeaconBlock{
			Block:     block,
			Signature: b.signature[:],
		}, nil
	case version.Altair:
		var block *sila.BeaconBlockAltair
		if blockMessage != nil {
			var ok bool
			block, ok = blockMessage.(*sila.BeaconBlockAltair)
			if !ok {
				return nil, errIncorrectBlockVersion
			}
		}
		return &sila.SignedBeaconBlockAltair{
			Block:     block,
			Signature: b.signature[:],
		}, nil
	case version.Bellatrix:
		if b.IsBlinded() {
			var block *sila.BlindedBeaconBlockBellatrix
			if blockMessage != nil {
				var ok bool
				block, ok = blockMessage.(*sila.BlindedBeaconBlockBellatrix)
				if !ok {
					return nil, errIncorrectBlockVersion
				}
			}
			return &sila.SignedBlindedBeaconBlockBellatrix{
				Block:     block,
				Signature: b.signature[:],
			}, nil
		}
		var block *sila.BeaconBlockBellatrix
		if blockMessage != nil {
			var ok bool
			block, ok = blockMessage.(*sila.BeaconBlockBellatrix)
			if !ok {
				return nil, errIncorrectBlockVersion
			}
		}
		return &sila.SignedBeaconBlockBellatrix{
			Block:     block,
			Signature: b.signature[:],
		}, nil
	case version.Capella:
		if b.IsBlinded() {
			var block *sila.BlindedBeaconBlockCapella
			if blockMessage != nil {
				var ok bool
				block, ok = blockMessage.(*sila.BlindedBeaconBlockCapella)
				if !ok {
					return nil, errIncorrectBlockVersion
				}
			}
			return &sila.SignedBlindedBeaconBlockCapella{
				Block:     block,
				Signature: b.signature[:],
			}, nil
		}
		var block *sila.BeaconBlockCapella
		if blockMessage != nil {
			var ok bool
			block, ok = blockMessage.(*sila.BeaconBlockCapella)
			if !ok {
				return nil, errIncorrectBlockVersion
			}
		}
		return &sila.SignedBeaconBlockCapella{
			Block:     block,
			Signature: b.signature[:],
		}, nil
	case version.Deneb:
		if b.IsBlinded() {
			var block *sila.BlindedBeaconBlockDeneb
			if blockMessage != nil {
				var ok bool
				block, ok = blockMessage.(*sila.BlindedBeaconBlockDeneb)
				if !ok {
					return nil, errIncorrectBlockVersion
				}
			}
			return &sila.SignedBlindedBeaconBlockDeneb{
				Message:   block,
				Signature: b.signature[:],
			}, nil
		}
		var block *sila.BeaconBlockDeneb
		if blockMessage != nil {
			var ok bool
			block, ok = blockMessage.(*sila.BeaconBlockDeneb)
			if !ok {
				return nil, errIncorrectBlockVersion
			}
		}
		return &sila.SignedBeaconBlockDeneb{
			Block:     block,
			Signature: b.signature[:],
		}, nil
	case version.Electra:
		if b.IsBlinded() {
			var block *sila.BlindedBeaconBlockElectra
			if blockMessage != nil {
				var ok bool
				block, ok = blockMessage.(*sila.BlindedBeaconBlockElectra)
				if !ok {
					return nil, errIncorrectBlockVersion
				}
			}
			return &sila.SignedBlindedBeaconBlockElectra{
				Message:   block,
				Signature: b.signature[:],
			}, nil
		}
		var block *sila.BeaconBlockElectra
		if blockMessage != nil {
			var ok bool
			block, ok = blockMessage.(*sila.BeaconBlockElectra)
			if !ok {
				return nil, errIncorrectBlockVersion
			}
		}
		return &sila.SignedBeaconBlockElectra{
			Block:     block,
			Signature: b.signature[:],
		}, nil
	case version.Fulu:
		if b.IsBlinded() {
			var block *sila.BlindedBeaconBlockFulu
			if blockMessage != nil {
				var ok bool
				block, ok = blockMessage.(*sila.BlindedBeaconBlockFulu)
				if !ok {
					return nil, errIncorrectBlockVersion
				}
			}
			return &sila.SignedBlindedBeaconBlockFulu{
				Message:   block,
				Signature: b.signature[:],
			}, nil
		}
		var block *sila.BeaconBlockElectra
		if blockMessage != nil {
			var ok bool
			block, ok = blockMessage.(*sila.BeaconBlockElectra)
			if !ok {
				return nil, errIncorrectBlockVersion
			}
		}
		return &sila.SignedBeaconBlockFulu{
			Block:     block,
			Signature: b.signature[:],
		}, nil
	case version.Gloas:
		var block *sila.BeaconBlockGloas
		if blockMessage != nil {
			var ok bool
			block, ok = blockMessage.(*sila.BeaconBlockGloas)
			if !ok {
				return nil, errIncorrectBlockVersion
			}
		}
		return &sila.SignedBeaconBlockGloas{
			Block:     block,
			Signature: b.signature[:],
		}, nil
	default:
		return nil, errors.New("unsupported signed beacon block version")
	}
}

// Proto converts the beacon block to a protobuf object.
func (b *BeaconBlock) Proto() (proto.Message, error) { // nolint:gocognit
	if b == nil {
		return nil, nil
	}

	bodyMessage, err := b.body.Proto()
	if err != nil {
		return nil, err
	}

	switch b.version {
	case version.Phase0:
		var body *sila.BeaconBlockBody
		if bodyMessage != nil {
			var ok bool
			body, ok = bodyMessage.(*sila.BeaconBlockBody)
			if !ok {
				return nil, errIncorrectBodyVersion
			}
		}
		return &sila.BeaconBlock{
			Slot:          b.slot,
			ProposerIndex: b.proposerIndex,
			ParentRoot:    b.parentRoot[:],
			StateRoot:     b.stateRoot[:],
			Body:          body,
		}, nil
	case version.Altair:
		var body *sila.BeaconBlockBodyAltair
		if bodyMessage != nil {
			var ok bool
			body, ok = bodyMessage.(*sila.BeaconBlockBodyAltair)
			if !ok {
				return nil, errIncorrectBodyVersion
			}
		}
		return &sila.BeaconBlockAltair{
			Slot:          b.slot,
			ProposerIndex: b.proposerIndex,
			ParentRoot:    b.parentRoot[:],
			StateRoot:     b.stateRoot[:],
			Body:          body,
		}, nil
	case version.Bellatrix:
		if b.IsBlinded() {
			var body *sila.BlindedBeaconBlockBodyBellatrix
			if bodyMessage != nil {
				var ok bool
				body, ok = bodyMessage.(*sila.BlindedBeaconBlockBodyBellatrix)
				if !ok {
					return nil, errIncorrectBodyVersion
				}
			}
			return &sila.BlindedBeaconBlockBellatrix{
				Slot:          b.slot,
				ProposerIndex: b.proposerIndex,
				ParentRoot:    b.parentRoot[:],
				StateRoot:     b.stateRoot[:],
				Body:          body,
			}, nil
		}
		var body *sila.BeaconBlockBodyBellatrix
		if bodyMessage != nil {
			var ok bool
			body, ok = bodyMessage.(*sila.BeaconBlockBodyBellatrix)
			if !ok {
				return nil, errIncorrectBodyVersion
			}
		}
		return &sila.BeaconBlockBellatrix{
			Slot:          b.slot,
			ProposerIndex: b.proposerIndex,
			ParentRoot:    b.parentRoot[:],
			StateRoot:     b.stateRoot[:],
			Body:          body,
		}, nil
	case version.Capella:
		if b.IsBlinded() {
			var body *sila.BlindedBeaconBlockBodyCapella
			if bodyMessage != nil {
				var ok bool
				body, ok = bodyMessage.(*sila.BlindedBeaconBlockBodyCapella)
				if !ok {
					return nil, errIncorrectBodyVersion
				}
			}
			return &sila.BlindedBeaconBlockCapella{
				Slot:          b.slot,
				ProposerIndex: b.proposerIndex,
				ParentRoot:    b.parentRoot[:],
				StateRoot:     b.stateRoot[:],
				Body:          body,
			}, nil
		}
		var body *sila.BeaconBlockBodyCapella
		if bodyMessage != nil {
			var ok bool
			body, ok = bodyMessage.(*sila.BeaconBlockBodyCapella)
			if !ok {
				return nil, errIncorrectBodyVersion
			}
		}
		return &sila.BeaconBlockCapella{
			Slot:          b.slot,
			ProposerIndex: b.proposerIndex,
			ParentRoot:    b.parentRoot[:],
			StateRoot:     b.stateRoot[:],
			Body:          body,
		}, nil
	case version.Deneb:
		if b.IsBlinded() {
			var body *sila.BlindedBeaconBlockBodyDeneb
			if bodyMessage != nil {
				var ok bool
				body, ok = bodyMessage.(*sila.BlindedBeaconBlockBodyDeneb)
				if !ok {
					return nil, errIncorrectBodyVersion
				}
			}
			return &sila.BlindedBeaconBlockDeneb{
				Slot:          b.slot,
				ProposerIndex: b.proposerIndex,
				ParentRoot:    b.parentRoot[:],
				StateRoot:     b.stateRoot[:],
				Body:          body,
			}, nil
		}
		var body *sila.BeaconBlockBodyDeneb
		if bodyMessage != nil {
			var ok bool
			body, ok = bodyMessage.(*sila.BeaconBlockBodyDeneb)
			if !ok {
				return nil, errIncorrectBodyVersion
			}
		}
		return &sila.BeaconBlockDeneb{
			Slot:          b.slot,
			ProposerIndex: b.proposerIndex,
			ParentRoot:    b.parentRoot[:],
			StateRoot:     b.stateRoot[:],
			Body:          body,
		}, nil
	case version.Electra:
		if b.IsBlinded() {
			var body *sila.BlindedBeaconBlockBodyElectra
			if bodyMessage != nil {
				var ok bool
				body, ok = bodyMessage.(*sila.BlindedBeaconBlockBodyElectra)
				if !ok {
					return nil, errIncorrectBodyVersion
				}
			}
			return &sila.BlindedBeaconBlockElectra{
				Slot:          b.slot,
				ProposerIndex: b.proposerIndex,
				ParentRoot:    b.parentRoot[:],
				StateRoot:     b.stateRoot[:],
				Body:          body,
			}, nil
		}
		var body *sila.BeaconBlockBodyElectra
		if bodyMessage != nil {
			var ok bool
			body, ok = bodyMessage.(*sila.BeaconBlockBodyElectra)
			if !ok {
				return nil, errIncorrectBodyVersion
			}
		}
		return &sila.BeaconBlockElectra{
			Slot:          b.slot,
			ProposerIndex: b.proposerIndex,
			ParentRoot:    b.parentRoot[:],
			StateRoot:     b.stateRoot[:],
			Body:          body,
		}, nil
	case version.Fulu:
		if b.IsBlinded() {
			var body *sila.BlindedBeaconBlockBodyElectra
			if bodyMessage != nil {
				var ok bool
				body, ok = bodyMessage.(*sila.BlindedBeaconBlockBodyElectra)
				if !ok {
					return nil, errIncorrectBodyVersion
				}
			}
			return &sila.BlindedBeaconBlockFulu{
				Slot:          b.slot,
				ProposerIndex: b.proposerIndex,
				ParentRoot:    b.parentRoot[:],
				StateRoot:     b.stateRoot[:],
				Body:          body,
			}, nil
		}
		var body *sila.BeaconBlockBodyElectra
		if bodyMessage != nil {
			var ok bool
			body, ok = bodyMessage.(*sila.BeaconBlockBodyElectra)
			if !ok {
				return nil, errIncorrectBodyVersion
			}
		}
		return &sila.BeaconBlockElectra{
			Slot:          b.slot,
			ProposerIndex: b.proposerIndex,
			ParentRoot:    b.parentRoot[:],
			StateRoot:     b.stateRoot[:],
			Body:          body,
		}, nil
	case version.Gloas:
		var body *sila.BeaconBlockBodyGloas
		if bodyMessage != nil {
			var ok bool
			body, ok = bodyMessage.(*sila.BeaconBlockBodyGloas)
			if !ok {
				return nil, errIncorrectBodyVersion
			}
		}
		return &sila.BeaconBlockGloas{
			Slot:          b.slot,
			ProposerIndex: b.proposerIndex,
			ParentRoot:    b.parentRoot[:],
			StateRoot:     b.stateRoot[:],
			Body:          body,
		}, nil
	default:
		return nil, fmt.Errorf("unsupported beacon block version: %s", version.String(b.version))
	}
}

// Proto converts the beacon block body to a protobuf object.
// nolint:gocognit
func (b *BeaconBlockBody) Proto() (proto.Message, error) {
	if b == nil {
		return nil, nil
	}

	switch b.version {
	case version.Phase0:
		return &sila.BeaconBlockBody{
			RandaoReveal:      b.randaoReveal[:],
			SilaData:          b.silaexecData,
			Graffiti:          b.graffiti[:],
			ProposerSlashings: b.proposerSlashings,
			AttesterSlashings: b.attesterSlashings,
			Attestations:      b.attestations,
			Deposits:          b.deposits,
			VoluntaryExits:    b.voluntaryExits,
		}, nil
	case version.Altair:
		return &sila.BeaconBlockBodyAltair{
			RandaoReveal:      b.randaoReveal[:],
			SilaData:          b.silaexecData,
			Graffiti:          b.graffiti[:],
			ProposerSlashings: b.proposerSlashings,
			AttesterSlashings: b.attesterSlashings,
			Attestations:      b.attestations,
			Deposits:          b.deposits,
			VoluntaryExits:    b.voluntaryExits,
			SyncAggregate:     b.syncAggregate,
		}, nil
	case version.Bellatrix:
		if b.IsBlinded() {
			var ph *silaenginev1.SilaPayloadHeader
			var ok bool
			if b.silaPayloadHeader != nil {
				ph, ok = b.silaPayloadHeader.Proto().(*silaenginev1.SilaPayloadHeader)
				if !ok {
					return nil, errPayloadHeaderWrongType
				}
			}
			return &sila.BlindedBeaconBlockBodyBellatrix{
				RandaoReveal:      b.randaoReveal[:],
				SilaData:          b.silaexecData,
				Graffiti:          b.graffiti[:],
				ProposerSlashings: b.proposerSlashings,
				AttesterSlashings: b.attesterSlashings,
				Attestations:      b.attestations,
				Deposits:          b.deposits,
				VoluntaryExits:    b.voluntaryExits,
				SyncAggregate:     b.syncAggregate,
				SilaPayloadHeader: ph,
			}, nil
		}
		var p *silaenginev1.SilaPayload
		var ok bool
		if b.silaPayload != nil {
			p, ok = b.silaPayload.Proto().(*silaenginev1.SilaPayload)
			if !ok {
				return nil, errPayloadWrongType
			}
		}
		return &sila.BeaconBlockBodyBellatrix{
			RandaoReveal:      b.randaoReveal[:],
			SilaData:          b.silaexecData,
			Graffiti:          b.graffiti[:],
			ProposerSlashings: b.proposerSlashings,
			AttesterSlashings: b.attesterSlashings,
			Attestations:      b.attestations,
			Deposits:          b.deposits,
			VoluntaryExits:    b.voluntaryExits,
			SyncAggregate:     b.syncAggregate,
			SilaPayload:       p,
		}, nil
	case version.Capella:
		if b.IsBlinded() {
			var ph *silaenginev1.SilaPayloadHeaderCapella
			var ok bool
			if b.silaPayloadHeader != nil {
				ph, ok = b.silaPayloadHeader.Proto().(*silaenginev1.SilaPayloadHeaderCapella)
				if !ok {
					return nil, errPayloadHeaderWrongType
				}
			}
			return &sila.BlindedBeaconBlockBodyCapella{
				RandaoReveal:      b.randaoReveal[:],
				SilaData:          b.silaexecData,
				Graffiti:          b.graffiti[:],
				ProposerSlashings: b.proposerSlashings,
				AttesterSlashings: b.attesterSlashings,
				Attestations:      b.attestations,
				Deposits:          b.deposits,
				VoluntaryExits:    b.voluntaryExits,
				SyncAggregate:     b.syncAggregate,
				SilaPayloadHeader: ph,
				BlsToSilaChanges:  b.blsToSilaChanges,
			}, nil
		}
		var p *silaenginev1.SilaPayloadCapella
		var ok bool
		if b.silaPayload != nil {
			p, ok = b.silaPayload.Proto().(*silaenginev1.SilaPayloadCapella)
			if !ok {
				return nil, errPayloadWrongType
			}
		}
		return &sila.BeaconBlockBodyCapella{
			RandaoReveal:      b.randaoReveal[:],
			SilaData:          b.silaexecData,
			Graffiti:          b.graffiti[:],
			ProposerSlashings: b.proposerSlashings,
			AttesterSlashings: b.attesterSlashings,
			Attestations:      b.attestations,
			Deposits:          b.deposits,
			VoluntaryExits:    b.voluntaryExits,
			SyncAggregate:     b.syncAggregate,
			SilaPayload:       p,
			BlsToSilaChanges:  b.blsToSilaChanges,
		}, nil
	case version.Deneb:
		if b.IsBlinded() {
			var ph *silaenginev1.SilaPayloadHeaderDeneb
			var ok bool
			if b.silaPayloadHeader != nil {
				ph, ok = b.silaPayloadHeader.Proto().(*silaenginev1.SilaPayloadHeaderDeneb)
				if !ok {
					return nil, errPayloadHeaderWrongType
				}
			}
			return &sila.BlindedBeaconBlockBodyDeneb{
				RandaoReveal:       b.randaoReveal[:],
				SilaData:           b.silaexecData,
				Graffiti:           b.graffiti[:],
				ProposerSlashings:  b.proposerSlashings,
				AttesterSlashings:  b.attesterSlashings,
				Attestations:       b.attestations,
				Deposits:           b.deposits,
				VoluntaryExits:     b.voluntaryExits,
				SyncAggregate:      b.syncAggregate,
				SilaPayloadHeader:  ph,
				BlsToSilaChanges:   b.blsToSilaChanges,
				BlobKzgCommitments: b.blobKzgCommitments,
			}, nil
		}
		var p *silaenginev1.SilaPayloadDeneb
		var ok bool
		if b.silaPayload != nil {
			p, ok = b.silaPayload.Proto().(*silaenginev1.SilaPayloadDeneb)
			if !ok {
				return nil, errPayloadWrongType
			}
		}
		return &sila.BeaconBlockBodyDeneb{
			RandaoReveal:       b.randaoReveal[:],
			SilaData:           b.silaexecData,
			Graffiti:           b.graffiti[:],
			ProposerSlashings:  b.proposerSlashings,
			AttesterSlashings:  b.attesterSlashings,
			Attestations:       b.attestations,
			Deposits:           b.deposits,
			VoluntaryExits:     b.voluntaryExits,
			SyncAggregate:      b.syncAggregate,
			SilaPayload:        p,
			BlsToSilaChanges:   b.blsToSilaChanges,
			BlobKzgCommitments: b.blobKzgCommitments,
		}, nil
	case version.Electra:
		if b.IsBlinded() {
			var ph *silaenginev1.SilaPayloadHeaderDeneb
			var ok bool
			if b.silaPayloadHeader != nil {
				ph, ok = b.silaPayloadHeader.Proto().(*silaenginev1.SilaPayloadHeaderDeneb)
				if !ok {
					return nil, errPayloadHeaderWrongType
				}
			}
			return &sila.BlindedBeaconBlockBodyElectra{
				RandaoReveal:       b.randaoReveal[:],
				SilaData:           b.silaexecData,
				Graffiti:           b.graffiti[:],
				ProposerSlashings:  b.proposerSlashings,
				AttesterSlashings:  b.attesterSlashingsElectra,
				Attestations:       b.attestationsElectra,
				Deposits:           b.deposits,
				VoluntaryExits:     b.voluntaryExits,
				SyncAggregate:      b.syncAggregate,
				SilaPayloadHeader:  ph,
				BlsToSilaChanges:   b.blsToSilaChanges,
				BlobKzgCommitments: b.blobKzgCommitments,
				SilaRequests:       b.silaRequests,
			}, nil
		}
		var p *silaenginev1.SilaPayloadDeneb
		var ok bool
		if b.silaPayload != nil {
			p, ok = b.silaPayload.Proto().(*silaenginev1.SilaPayloadDeneb)
			if !ok {
				return nil, errPayloadWrongType
			}
		}
		return &sila.BeaconBlockBodyElectra{
			RandaoReveal:       b.randaoReveal[:],
			SilaData:           b.silaexecData,
			Graffiti:           b.graffiti[:],
			ProposerSlashings:  b.proposerSlashings,
			AttesterSlashings:  b.attesterSlashingsElectra,
			Attestations:       b.attestationsElectra,
			Deposits:           b.deposits,
			VoluntaryExits:     b.voluntaryExits,
			SyncAggregate:      b.syncAggregate,
			SilaPayload:        p,
			BlsToSilaChanges:   b.blsToSilaChanges,
			BlobKzgCommitments: b.blobKzgCommitments,
			SilaRequests:       b.silaRequests,
		}, nil
	case version.Fulu:
		if b.IsBlinded() {
			var ph *silaenginev1.SilaPayloadHeaderDeneb
			var ok bool
			if b.silaPayloadHeader != nil {
				ph, ok = b.silaPayloadHeader.Proto().(*silaenginev1.SilaPayloadHeaderDeneb)
				if !ok {
					return nil, errPayloadHeaderWrongType
				}
			}
			return &sila.BlindedBeaconBlockBodyElectra{
				RandaoReveal:       b.randaoReveal[:],
				SilaData:           b.silaexecData,
				Graffiti:           b.graffiti[:],
				ProposerSlashings:  b.proposerSlashings,
				AttesterSlashings:  b.attesterSlashingsElectra,
				Attestations:       b.attestationsElectra,
				Deposits:           b.deposits,
				VoluntaryExits:     b.voluntaryExits,
				SyncAggregate:      b.syncAggregate,
				SilaPayloadHeader:  ph,
				BlsToSilaChanges:   b.blsToSilaChanges,
				BlobKzgCommitments: b.blobKzgCommitments,
				SilaRequests:       b.silaRequests,
			}, nil
		}
		var p *silaenginev1.SilaPayloadDeneb
		var ok bool
		if b.silaPayload != nil {
			p, ok = b.silaPayload.Proto().(*silaenginev1.SilaPayloadDeneb)
			if !ok {
				return nil, errPayloadWrongType
			}
		}
		return &sila.BeaconBlockBodyElectra{
			RandaoReveal:       b.randaoReveal[:],
			SilaData:           b.silaexecData,
			Graffiti:           b.graffiti[:],
			ProposerSlashings:  b.proposerSlashings,
			AttesterSlashings:  b.attesterSlashingsElectra,
			Attestations:       b.attestationsElectra,
			Deposits:           b.deposits,
			VoluntaryExits:     b.voluntaryExits,
			SyncAggregate:      b.syncAggregate,
			SilaPayload:        p,
			BlsToSilaChanges:   b.blsToSilaChanges,
			BlobKzgCommitments: b.blobKzgCommitments,
			SilaRequests:       b.silaRequests,
		}, nil
	case version.Gloas:
		return &sila.BeaconBlockBodyGloas{
			RandaoReveal:         b.randaoReveal[:],
			SilaData:             b.silaexecData,
			Graffiti:             b.graffiti[:],
			ProposerSlashings:    b.proposerSlashings,
			AttesterSlashings:    b.attesterSlashingsElectra,
			Attestations:         b.attestationsElectra,
			Deposits:             b.deposits,
			VoluntaryExits:       b.voluntaryExits,
			SyncAggregate:        b.syncAggregate,
			BlsToSilaChanges:     b.blsToSilaChanges,
			SignedSilaPayloadBid: b.signedSilaPayloadBid,
			PayloadAttestations:  b.payloadAttestations,
			ParentSilaRequests:   b.parentSilaRequests,
		}, nil
	default:
		return nil, errors.New("unsupported beacon block body version")
	}
}

// ----------------------------------------------------------------------------
// Phase 0
// ----------------------------------------------------------------------------

func initSignedBlockFromProtoPhase0(pb *sila.SignedBeaconBlock) (*SignedBeaconBlock, error) {
	if pb == nil {
		return nil, errNilBlock
	}

	block, err := initBlockFromProtoPhase0(pb.Block)
	if err != nil {
		return nil, err
	}
	b := &SignedBeaconBlock{
		version:   version.Phase0,
		block:     block,
		signature: bytesutil.ToBytes96(pb.Signature),
	}
	return b, nil
}

func initBlockFromProtoPhase0(pb *sila.BeaconBlock) (*BeaconBlock, error) {
	if pb == nil {
		return nil, errNilBlock
	}

	body, err := initBlockBodyFromProtoPhase0(pb.Body)
	if err != nil {
		return nil, err
	}
	b := &BeaconBlock{
		version:       version.Phase0,
		slot:          pb.Slot,
		proposerIndex: pb.ProposerIndex,
		parentRoot:    bytesutil.ToBytes32(pb.ParentRoot),
		stateRoot:     bytesutil.ToBytes32(pb.StateRoot),
		body:          body,
	}
	return b, nil
}

func initBlockBodyFromProtoPhase0(pb *sila.BeaconBlockBody) (*BeaconBlockBody, error) {
	if pb == nil {
		return nil, errNilBlockBody
	}

	b := &BeaconBlockBody{
		version:           version.Phase0,
		randaoReveal:      bytesutil.ToBytes96(pb.RandaoReveal),
		silaexecData:      pb.SilaData,
		graffiti:          bytesutil.ToBytes32(pb.Graffiti),
		proposerSlashings: pb.ProposerSlashings,
		attesterSlashings: pb.AttesterSlashings,
		attestations:      pb.Attestations,
		deposits:          pb.Deposits,
		voluntaryExits:    pb.VoluntaryExits,
	}
	return b, nil
}

// ----------------------------------------------------------------------------
// Altair
// ----------------------------------------------------------------------------

func initSignedBlockFromProtoAltair(pb *sila.SignedBeaconBlockAltair) (*SignedBeaconBlock, error) {
	if pb == nil {
		return nil, errNilBlock
	}

	block, err := initBlockFromProtoAltair(pb.Block)
	if err != nil {
		return nil, err
	}
	b := &SignedBeaconBlock{
		version:   version.Altair,
		block:     block,
		signature: bytesutil.ToBytes96(pb.Signature),
	}
	return b, nil
}

func initBlockFromProtoAltair(pb *sila.BeaconBlockAltair) (*BeaconBlock, error) {
	if pb == nil {
		return nil, errNilBlock
	}

	body, err := initBlockBodyFromProtoAltair(pb.Body)
	if err != nil {
		return nil, err
	}
	b := &BeaconBlock{
		version:       version.Altair,
		slot:          pb.Slot,
		proposerIndex: pb.ProposerIndex,
		parentRoot:    bytesutil.ToBytes32(pb.ParentRoot),
		stateRoot:     bytesutil.ToBytes32(pb.StateRoot),
		body:          body,
	}
	return b, nil
}

func initBlockBodyFromProtoAltair(pb *sila.BeaconBlockBodyAltair) (*BeaconBlockBody, error) {
	if pb == nil {
		return nil, errNilBlockBody
	}

	b := &BeaconBlockBody{
		version:           version.Altair,
		randaoReveal:      bytesutil.ToBytes96(pb.RandaoReveal),
		silaexecData:      pb.SilaData,
		graffiti:          bytesutil.ToBytes32(pb.Graffiti),
		proposerSlashings: pb.ProposerSlashings,
		attesterSlashings: pb.AttesterSlashings,
		attestations:      pb.Attestations,
		deposits:          pb.Deposits,
		voluntaryExits:    pb.VoluntaryExits,
		syncAggregate:     pb.SyncAggregate,
	}
	return b, nil
}

// ----------------------------------------------------------------------------
// Bellatrix
// ----------------------------------------------------------------------------

func initSignedBlockFromProtoBellatrix(pb *sila.SignedBeaconBlockBellatrix) (*SignedBeaconBlock, error) {
	if pb == nil {
		return nil, errNilBlock
	}

	block, err := initBlockFromProtoBellatrix(pb.Block)
	if err != nil {
		return nil, err
	}
	b := &SignedBeaconBlock{
		version:   version.Bellatrix,
		block:     block,
		signature: bytesutil.ToBytes96(pb.Signature),
	}
	return b, nil
}

func initBlindedSignedBlockFromProtoBellatrix(pb *sila.SignedBlindedBeaconBlockBellatrix) (*SignedBeaconBlock, error) {
	if pb == nil {
		return nil, errNilBlock
	}

	block, err := initBlindedBlockFromProtoBellatrix(pb.Block)
	if err != nil {
		return nil, err
	}
	b := &SignedBeaconBlock{
		version:   version.Bellatrix,
		block:     block,
		signature: bytesutil.ToBytes96(pb.Signature),
	}
	return b, nil
}

func initBlockFromProtoBellatrix(pb *sila.BeaconBlockBellatrix) (*BeaconBlock, error) {
	if pb == nil {
		return nil, errNilBlock
	}

	body, err := initBlockBodyFromProtoBellatrix(pb.Body)
	if err != nil {
		return nil, err
	}
	b := &BeaconBlock{
		version:       version.Bellatrix,
		slot:          pb.Slot,
		proposerIndex: pb.ProposerIndex,
		parentRoot:    bytesutil.ToBytes32(pb.ParentRoot),
		stateRoot:     bytesutil.ToBytes32(pb.StateRoot),
		body:          body,
	}
	return b, nil
}

func initBlindedBlockFromProtoBellatrix(pb *sila.BlindedBeaconBlockBellatrix) (*BeaconBlock, error) {
	if pb == nil {
		return nil, errNilBlock
	}

	body, err := initBlindedBlockBodyFromProtoBellatrix(pb.Body)
	if err != nil {
		return nil, err
	}
	b := &BeaconBlock{
		version:       version.Bellatrix,
		slot:          pb.Slot,
		proposerIndex: pb.ProposerIndex,
		parentRoot:    bytesutil.ToBytes32(pb.ParentRoot),
		stateRoot:     bytesutil.ToBytes32(pb.StateRoot),
		body:          body,
	}
	return b, nil
}

func initBlockBodyFromProtoBellatrix(pb *sila.BeaconBlockBodyBellatrix) (*BeaconBlockBody, error) {
	if pb == nil {
		return nil, errNilBlockBody
	}

	p, err := WrappedSilaPayload(pb.SilaPayload)
	// We allow the payload to be nil
	if err != nil && !errors.Is(err, consensus_types.ErrNilObjectWrapped) {
		return nil, err
	}
	b := &BeaconBlockBody{
		version:           version.Bellatrix,
		randaoReveal:      bytesutil.ToBytes96(pb.RandaoReveal),
		silaexecData:      pb.SilaData,
		graffiti:          bytesutil.ToBytes32(pb.Graffiti),
		proposerSlashings: pb.ProposerSlashings,
		attesterSlashings: pb.AttesterSlashings,
		attestations:      pb.Attestations,
		deposits:          pb.Deposits,
		voluntaryExits:    pb.VoluntaryExits,
		syncAggregate:     pb.SyncAggregate,
		silaPayload:       p,
	}
	return b, nil
}

func initBlindedBlockBodyFromProtoBellatrix(pb *sila.BlindedBeaconBlockBodyBellatrix) (*BeaconBlockBody, error) {
	if pb == nil {
		return nil, errNilBlockBody
	}

	ph, err := WrappedSilaPayloadHeader(pb.SilaPayloadHeader)
	// We allow the payload to be nil
	if err != nil && !errors.Is(err, consensus_types.ErrNilObjectWrapped) {
		return nil, err
	}
	b := &BeaconBlockBody{
		version:           version.Bellatrix,
		randaoReveal:      bytesutil.ToBytes96(pb.RandaoReveal),
		silaexecData:      pb.SilaData,
		graffiti:          bytesutil.ToBytes32(pb.Graffiti),
		proposerSlashings: pb.ProposerSlashings,
		attesterSlashings: pb.AttesterSlashings,
		attestations:      pb.Attestations,
		deposits:          pb.Deposits,
		voluntaryExits:    pb.VoluntaryExits,
		syncAggregate:     pb.SyncAggregate,
		silaPayloadHeader: ph,
	}
	return b, nil
}

// ----------------------------------------------------------------------------
// Capella
// ----------------------------------------------------------------------------

func initSignedBlockFromProtoCapella(pb *sila.SignedBeaconBlockCapella) (*SignedBeaconBlock, error) {
	if pb == nil {
		return nil, errNilBlock
	}

	block, err := initBlockFromProtoCapella(pb.Block)
	if err != nil {
		return nil, err
	}
	b := &SignedBeaconBlock{
		version:   version.Capella,
		block:     block,
		signature: bytesutil.ToBytes96(pb.Signature),
	}
	return b, nil
}

func initBlindedSignedBlockFromProtoCapella(pb *sila.SignedBlindedBeaconBlockCapella) (*SignedBeaconBlock, error) {
	if pb == nil {
		return nil, errNilBlock
	}

	block, err := initBlindedBlockFromProtoCapella(pb.Block)
	if err != nil {
		return nil, err
	}
	b := &SignedBeaconBlock{
		version:   version.Capella,
		block:     block,
		signature: bytesutil.ToBytes96(pb.Signature),
	}
	return b, nil
}

func initBlockFromProtoCapella(pb *sila.BeaconBlockCapella) (*BeaconBlock, error) {
	if pb == nil {
		return nil, errNilBlock
	}

	body, err := initBlockBodyFromProtoCapella(pb.Body)
	if err != nil {
		return nil, err
	}
	b := &BeaconBlock{
		version:       version.Capella,
		slot:          pb.Slot,
		proposerIndex: pb.ProposerIndex,
		parentRoot:    bytesutil.ToBytes32(pb.ParentRoot),
		stateRoot:     bytesutil.ToBytes32(pb.StateRoot),
		body:          body,
	}
	return b, nil
}

func initBlindedBlockFromProtoCapella(pb *sila.BlindedBeaconBlockCapella) (*BeaconBlock, error) {
	if pb == nil {
		return nil, errNilBlock
	}

	body, err := initBlindedBlockBodyFromProtoCapella(pb.Body)
	if err != nil {
		return nil, err
	}
	b := &BeaconBlock{
		version:       version.Capella,
		slot:          pb.Slot,
		proposerIndex: pb.ProposerIndex,
		parentRoot:    bytesutil.ToBytes32(pb.ParentRoot),
		stateRoot:     bytesutil.ToBytes32(pb.StateRoot),
		body:          body,
	}
	return b, nil
}

func initBlockBodyFromProtoCapella(pb *sila.BeaconBlockBodyCapella) (*BeaconBlockBody, error) {
	if pb == nil {
		return nil, errNilBlockBody
	}

	p, err := WrappedSilaPayloadCapella(pb.SilaPayload)
	// We allow the payload to be nil
	if err != nil && !errors.Is(err, consensus_types.ErrNilObjectWrapped) {
		return nil, err
	}
	b := &BeaconBlockBody{
		version:           version.Capella,
		randaoReveal:      bytesutil.ToBytes96(pb.RandaoReveal),
		silaexecData:      pb.SilaData,
		graffiti:          bytesutil.ToBytes32(pb.Graffiti),
		proposerSlashings: pb.ProposerSlashings,
		attesterSlashings: pb.AttesterSlashings,
		attestations:      pb.Attestations,
		deposits:          pb.Deposits,
		voluntaryExits:    pb.VoluntaryExits,
		syncAggregate:     pb.SyncAggregate,
		silaPayload:       p,
		blsToSilaChanges:  pb.BlsToSilaChanges,
	}
	return b, nil
}

func initBlindedBlockBodyFromProtoCapella(pb *sila.BlindedBeaconBlockBodyCapella) (*BeaconBlockBody, error) {
	if pb == nil {
		return nil, errNilBlockBody
	}

	ph, err := WrappedSilaPayloadHeaderCapella(pb.SilaPayloadHeader)
	// We allow the payload to be nil
	if err != nil && !errors.Is(err, consensus_types.ErrNilObjectWrapped) {
		return nil, err
	}
	b := &BeaconBlockBody{
		version:           version.Capella,
		randaoReveal:      bytesutil.ToBytes96(pb.RandaoReveal),
		silaexecData:      pb.SilaData,
		graffiti:          bytesutil.ToBytes32(pb.Graffiti),
		proposerSlashings: pb.ProposerSlashings,
		attesterSlashings: pb.AttesterSlashings,
		attestations:      pb.Attestations,
		deposits:          pb.Deposits,
		voluntaryExits:    pb.VoluntaryExits,
		syncAggregate:     pb.SyncAggregate,
		silaPayloadHeader: ph,
		blsToSilaChanges:  pb.BlsToSilaChanges,
	}
	return b, nil
}

// ----------------------------------------------------------------------------
// Deneb
// ----------------------------------------------------------------------------

func initSignedBlockFromProtoDeneb(pb *sila.SignedBeaconBlockDeneb) (*SignedBeaconBlock, error) {
	if pb == nil {
		return nil, errNilBlock
	}

	block, err := initBlockFromProtoDeneb(pb.Block)
	if err != nil {
		return nil, err
	}
	b := &SignedBeaconBlock{
		version:   version.Deneb,
		block:     block,
		signature: bytesutil.ToBytes96(pb.Signature),
	}
	return b, nil
}

func initBlindedSignedBlockFromProtoDeneb(pb *sila.SignedBlindedBeaconBlockDeneb) (*SignedBeaconBlock, error) {
	if pb == nil {
		return nil, errNilBlock
	}

	block, err := initBlindedBlockFromProtoDeneb(pb.Message)
	if err != nil {
		return nil, err
	}
	b := &SignedBeaconBlock{
		version:   version.Deneb,
		block:     block,
		signature: bytesutil.ToBytes96(pb.Signature),
	}
	return b, nil
}

func initBlockFromProtoDeneb(pb *sila.BeaconBlockDeneb) (*BeaconBlock, error) {
	if pb == nil {
		return nil, errNilBlock
	}

	body, err := initBlockBodyFromProtoDeneb(pb.Body)
	if err != nil {
		return nil, err
	}
	b := &BeaconBlock{
		version:       version.Deneb,
		slot:          pb.Slot,
		proposerIndex: pb.ProposerIndex,
		parentRoot:    bytesutil.ToBytes32(pb.ParentRoot),
		stateRoot:     bytesutil.ToBytes32(pb.StateRoot),
		body:          body,
	}
	return b, nil
}

func initBlindedBlockFromProtoDeneb(pb *sila.BlindedBeaconBlockDeneb) (*BeaconBlock, error) {
	if pb == nil {
		return nil, errNilBlock
	}

	body, err := initBlindedBlockBodyFromProtoDeneb(pb.Body)
	if err != nil {
		return nil, err
	}
	b := &BeaconBlock{
		version:       version.Deneb,
		slot:          pb.Slot,
		proposerIndex: pb.ProposerIndex,
		parentRoot:    bytesutil.ToBytes32(pb.ParentRoot),
		stateRoot:     bytesutil.ToBytes32(pb.StateRoot),
		body:          body,
	}
	return b, nil
}

func initBlockBodyFromProtoDeneb(pb *sila.BeaconBlockBodyDeneb) (*BeaconBlockBody, error) {
	if pb == nil {
		return nil, errNilBlockBody
	}

	p, err := WrappedSilaPayloadDeneb(pb.SilaPayload)
	// We allow the payload to be nil
	if err != nil && !errors.Is(err, consensus_types.ErrNilObjectWrapped) {
		return nil, err
	}
	b := &BeaconBlockBody{
		version:            version.Deneb,
		randaoReveal:       bytesutil.ToBytes96(pb.RandaoReveal),
		silaexecData:       pb.SilaData,
		graffiti:           bytesutil.ToBytes32(pb.Graffiti),
		proposerSlashings:  pb.ProposerSlashings,
		attesterSlashings:  pb.AttesterSlashings,
		attestations:       pb.Attestations,
		deposits:           pb.Deposits,
		voluntaryExits:     pb.VoluntaryExits,
		syncAggregate:      pb.SyncAggregate,
		silaPayload:        p,
		blsToSilaChanges:   pb.BlsToSilaChanges,
		blobKzgCommitments: pb.BlobKzgCommitments,
	}
	return b, nil
}

func initBlindedBlockBodyFromProtoDeneb(pb *sila.BlindedBeaconBlockBodyDeneb) (*BeaconBlockBody, error) {
	if pb == nil {
		return nil, errNilBlockBody
	}

	ph, err := WrappedSilaPayloadHeaderDeneb(pb.SilaPayloadHeader)
	// We allow the payload to be nil
	if err != nil && !errors.Is(err, consensus_types.ErrNilObjectWrapped) {
		return nil, err
	}
	b := &BeaconBlockBody{
		version:            version.Deneb,
		randaoReveal:       bytesutil.ToBytes96(pb.RandaoReveal),
		silaexecData:       pb.SilaData,
		graffiti:           bytesutil.ToBytes32(pb.Graffiti),
		proposerSlashings:  pb.ProposerSlashings,
		attesterSlashings:  pb.AttesterSlashings,
		attestations:       pb.Attestations,
		deposits:           pb.Deposits,
		voluntaryExits:     pb.VoluntaryExits,
		syncAggregate:      pb.SyncAggregate,
		silaPayloadHeader:  ph,
		blsToSilaChanges:   pb.BlsToSilaChanges,
		blobKzgCommitments: pb.BlobKzgCommitments,
	}
	return b, nil
}

// ----------------------------------------------------------------------------
// Electra
// ----------------------------------------------------------------------------

func initSignedBlockFromProtoElectra(pb *sila.SignedBeaconBlockElectra) (*SignedBeaconBlock, error) {
	if pb == nil {
		return nil, errNilBlock
	}

	block, err := initBlockFromProtoElectra(pb.Block)
	if err != nil {
		return nil, err
	}
	b := &SignedBeaconBlock{
		version:   version.Electra,
		block:     block,
		signature: bytesutil.ToBytes96(pb.Signature),
	}
	return b, nil
}

func initBlindedSignedBlockFromProtoElectra(pb *sila.SignedBlindedBeaconBlockElectra) (*SignedBeaconBlock, error) {
	if pb == nil {
		return nil, errNilBlock
	}

	block, err := initBlindedBlockFromProtoElectra(pb.Message)
	if err != nil {
		return nil, err
	}
	b := &SignedBeaconBlock{
		version:   version.Electra,
		block:     block,
		signature: bytesutil.ToBytes96(pb.Signature),
	}
	return b, nil
}

func initBlockFromProtoElectra(pb *sila.BeaconBlockElectra) (*BeaconBlock, error) {
	if pb == nil {
		return nil, errNilBlock
	}

	body, err := initBlockBodyFromProtoElectra(pb.Body)
	if err != nil {
		return nil, err
	}
	b := &BeaconBlock{
		version:       version.Electra,
		slot:          pb.Slot,
		proposerIndex: pb.ProposerIndex,
		parentRoot:    bytesutil.ToBytes32(pb.ParentRoot),
		stateRoot:     bytesutil.ToBytes32(pb.StateRoot),
		body:          body,
	}
	return b, nil
}

func initBlindedBlockFromProtoElectra(pb *sila.BlindedBeaconBlockElectra) (*BeaconBlock, error) {
	if pb == nil {
		return nil, errNilBlock
	}

	body, err := initBlindedBlockBodyFromProtoElectra(pb.Body)
	if err != nil {
		return nil, err
	}
	b := &BeaconBlock{
		version:       version.Electra,
		slot:          pb.Slot,
		proposerIndex: pb.ProposerIndex,
		parentRoot:    bytesutil.ToBytes32(pb.ParentRoot),
		stateRoot:     bytesutil.ToBytes32(pb.StateRoot),
		body:          body,
	}
	return b, nil
}

func initBlockBodyFromProtoElectra(pb *sila.BeaconBlockBodyElectra) (*BeaconBlockBody, error) {
	if pb == nil {
		return nil, errNilBlockBody
	}

	p, err := WrappedSilaPayloadDeneb(pb.SilaPayload)
	// We allow the payload to be nil
	if err != nil && !errors.Is(err, consensus_types.ErrNilObjectWrapped) {
		return nil, err
	}
	er := pb.SilaRequests
	if er == nil {
		er = &silaenginev1.SilaRequests{}
	}
	b := &BeaconBlockBody{
		version:                  version.Electra,
		randaoReveal:             bytesutil.ToBytes96(pb.RandaoReveal),
		silaexecData:             pb.SilaData,
		graffiti:                 bytesutil.ToBytes32(pb.Graffiti),
		proposerSlashings:        pb.ProposerSlashings,
		attesterSlashingsElectra: pb.AttesterSlashings,
		attestationsElectra:      pb.Attestations,
		deposits:                 pb.Deposits,
		voluntaryExits:           pb.VoluntaryExits,
		syncAggregate:            pb.SyncAggregate,
		silaPayload:              p,
		blsToSilaChanges:         pb.BlsToSilaChanges,
		blobKzgCommitments:       pb.BlobKzgCommitments,
		silaRequests:             er,
	}
	return b, nil
}

func initBlindedBlockBodyFromProtoElectra(pb *sila.BlindedBeaconBlockBodyElectra) (*BeaconBlockBody, error) {
	if pb == nil {
		return nil, errNilBlockBody
	}

	ph, err := WrappedSilaPayloadHeaderDeneb(pb.SilaPayloadHeader)
	// We allow the payload to be nil
	if err != nil && !errors.Is(err, consensus_types.ErrNilObjectWrapped) {
		return nil, err
	}
	er := pb.SilaRequests
	if er == nil {
		er = &silaenginev1.SilaRequests{}
	}
	b := &BeaconBlockBody{
		version:                  version.Electra,
		randaoReveal:             bytesutil.ToBytes96(pb.RandaoReveal),
		silaexecData:             pb.SilaData,
		graffiti:                 bytesutil.ToBytes32(pb.Graffiti),
		proposerSlashings:        pb.ProposerSlashings,
		attesterSlashingsElectra: pb.AttesterSlashings,
		attestationsElectra:      pb.Attestations,
		deposits:                 pb.Deposits,
		voluntaryExits:           pb.VoluntaryExits,
		syncAggregate:            pb.SyncAggregate,
		silaPayloadHeader:        ph,
		blsToSilaChanges:         pb.BlsToSilaChanges,
		blobKzgCommitments:       pb.BlobKzgCommitments,
		silaRequests:             er,
	}
	return b, nil
}

// ----------------------------------------------------------------------------
// Fulu
// ----------------------------------------------------------------------------

func initSignedBlockFromProtoFulu(pb *sila.SignedBeaconBlockFulu) (*SignedBeaconBlock, error) {
	if pb == nil {
		return nil, errNilBlock
	}

	block, err := initBlockFromProtoFulu(pb.Block)
	if err != nil {
		return nil, err
	}
	b := &SignedBeaconBlock{
		version:   version.Fulu,
		block:     block,
		signature: bytesutil.ToBytes96(pb.Signature),
	}
	return b, nil
}

func initBlindedSignedBlockFromProtoFulu(pb *sila.SignedBlindedBeaconBlockFulu) (*SignedBeaconBlock, error) {
	if pb == nil {
		return nil, errNilBlock
	}

	block, err := initBlindedBlockFromProtoFulu(pb.Message)
	if err != nil {
		return nil, err
	}
	b := &SignedBeaconBlock{
		version:   version.Fulu,
		block:     block,
		signature: bytesutil.ToBytes96(pb.Signature),
	}
	return b, nil
}

func initBlockFromProtoFulu(pb *sila.BeaconBlockElectra) (*BeaconBlock, error) {
	if pb == nil {
		return nil, errNilBlock
	}

	body, err := initBlockBodyFromProtoFulu(pb.Body)
	if err != nil {
		return nil, err
	}
	b := &BeaconBlock{
		version:       version.Fulu,
		slot:          pb.Slot,
		proposerIndex: pb.ProposerIndex,
		parentRoot:    bytesutil.ToBytes32(pb.ParentRoot),
		stateRoot:     bytesutil.ToBytes32(pb.StateRoot),
		body:          body,
	}
	return b, nil
}

func initBlindedBlockFromProtoFulu(pb *sila.BlindedBeaconBlockFulu) (*BeaconBlock, error) {
	if pb == nil {
		return nil, errNilBlock
	}

	body, err := initBlindedBlockBodyFromProtoFulu(pb.Body)
	if err != nil {
		return nil, err
	}
	b := &BeaconBlock{
		version:       version.Fulu,
		slot:          pb.Slot,
		proposerIndex: pb.ProposerIndex,
		parentRoot:    bytesutil.ToBytes32(pb.ParentRoot),
		stateRoot:     bytesutil.ToBytes32(pb.StateRoot),
		body:          body,
	}
	return b, nil
}

func initBlockBodyFromProtoFulu(pb *sila.BeaconBlockBodyElectra) (*BeaconBlockBody, error) {
	if pb == nil {
		return nil, errNilBlockBody
	}

	p, err := WrappedSilaPayloadDeneb(pb.SilaPayload)
	// We allow the payload to be nil
	if err != nil && !errors.Is(err, consensus_types.ErrNilObjectWrapped) {
		return nil, err
	}
	er := pb.SilaRequests
	if er == nil {
		er = &silaenginev1.SilaRequests{}
	}
	b := &BeaconBlockBody{
		version:                  version.Fulu,
		randaoReveal:             bytesutil.ToBytes96(pb.RandaoReveal),
		silaexecData:             pb.SilaData,
		graffiti:                 bytesutil.ToBytes32(pb.Graffiti),
		proposerSlashings:        pb.ProposerSlashings,
		attesterSlashingsElectra: pb.AttesterSlashings,
		attestationsElectra:      pb.Attestations,
		deposits:                 pb.Deposits,
		voluntaryExits:           pb.VoluntaryExits,
		syncAggregate:            pb.SyncAggregate,
		silaPayload:              p,
		blsToSilaChanges:         pb.BlsToSilaChanges,
		blobKzgCommitments:       pb.BlobKzgCommitments,
		silaRequests:             er,
	}
	return b, nil
}

func initBlindedBlockBodyFromProtoFulu(pb *sila.BlindedBeaconBlockBodyElectra) (*BeaconBlockBody, error) {
	if pb == nil {
		return nil, errNilBlockBody
	}

	ph, err := WrappedSilaPayloadHeaderDeneb(pb.SilaPayloadHeader)
	// We allow the payload to be nil
	if err != nil && !errors.Is(err, consensus_types.ErrNilObjectWrapped) {
		return nil, err
	}
	er := pb.SilaRequests
	if er == nil {
		er = &silaenginev1.SilaRequests{}
	}
	b := &BeaconBlockBody{
		version:                  version.Fulu,
		randaoReveal:             bytesutil.ToBytes96(pb.RandaoReveal),
		silaexecData:             pb.SilaData,
		graffiti:                 bytesutil.ToBytes32(pb.Graffiti),
		proposerSlashings:        pb.ProposerSlashings,
		attesterSlashingsElectra: pb.AttesterSlashings,
		attestationsElectra:      pb.Attestations,
		deposits:                 pb.Deposits,
		voluntaryExits:           pb.VoluntaryExits,
		syncAggregate:            pb.SyncAggregate,
		silaPayloadHeader:        ph,
		blsToSilaChanges:         pb.BlsToSilaChanges,
		blobKzgCommitments:       pb.BlobKzgCommitments,
		silaRequests:             er,
	}
	return b, nil
}

// ----------------------------------------------------------------------------
// Gloas
// ----------------------------------------------------------------------------

func initSignedBlockFromProtoGloas(pb *sila.SignedBeaconBlockGloas) (*SignedBeaconBlock, error) {
	if pb == nil {
		return nil, errNilBlock
	}

	block, err := initBlockFromProtoGloas(pb.Block)
	if err != nil {
		return nil, err
	}
	b := &SignedBeaconBlock{
		version:   version.Gloas,
		block:     block,
		signature: bytesutil.ToBytes96(pb.Signature),
	}
	return b, nil
}

func initBlockFromProtoGloas(pb *sila.BeaconBlockGloas) (*BeaconBlock, error) {
	if pb == nil {
		return nil, errNilBlock
	}

	body, err := initBlockBodyFromProtoGloas(pb.Body)
	if err != nil {
		return nil, err
	}
	b := &BeaconBlock{
		version:       version.Gloas,
		slot:          pb.Slot,
		proposerIndex: pb.ProposerIndex,
		parentRoot:    bytesutil.ToBytes32(pb.ParentRoot),
		stateRoot:     bytesutil.ToBytes32(pb.StateRoot),
		body:          body,
	}
	return b, nil
}

func initBlockBodyFromProtoGloas(pb *sila.BeaconBlockBodyGloas) (*BeaconBlockBody, error) {
	if pb == nil {
		return nil, errNilBlockBody
	}

	per := pb.ParentSilaRequests
	if per == nil {
		per = &silaenginev1.SilaRequests{}
	}
	b := &BeaconBlockBody{
		version:                  version.Gloas,
		randaoReveal:             bytesutil.ToBytes96(pb.RandaoReveal),
		silaexecData:             pb.SilaData,
		graffiti:                 bytesutil.ToBytes32(pb.Graffiti),
		proposerSlashings:        pb.ProposerSlashings,
		attesterSlashingsElectra: pb.AttesterSlashings,
		attestationsElectra:      pb.Attestations,
		deposits:                 pb.Deposits,
		voluntaryExits:           pb.VoluntaryExits,
		syncAggregate:            pb.SyncAggregate,
		blsToSilaChanges:         pb.BlsToSilaChanges,
		signedSilaPayloadBid:     pb.SignedSilaPayloadBid,
		payloadAttestations:      pb.PayloadAttestations,
		parentSilaRequests:       per,
	}
	return b, nil
}
