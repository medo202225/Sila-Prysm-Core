package blocks

import (
	"testing"

	"github.com/sila-chain/Sila-Consensus-Core/v7/config/params"
	"github.com/sila-chain/Sila-Consensus-Core/v7/consensus-types/primitives"
	"github.com/sila-chain/Sila-Consensus-Core/v7/encoding/bytesutil"
	sila "github.com/sila-chain/Sila-Consensus-Core/v7/proto/sila/v1alpha1"
	silaenginev1 "github.com/sila-chain/Sila-Consensus-Core/v7/proto/silaengine/v1"
	"github.com/sila-chain/Sila-Consensus-Core/v7/runtime/version"
	"github.com/sila-chain/Sila-Consensus-Core/v7/testing/assert"
	"github.com/sila-chain/Sila-Consensus-Core/v7/testing/require"
	"github.com/sila-chain/go-bitfield"
)

type fields struct {
	root                     [32]byte
	sig                      [96]byte
	deposits                 []*sila.Deposit
	atts                     []*sila.Attestation
	attsElectra              []*sila.AttestationElectra
	proposerSlashings        []*sila.ProposerSlashing
	attesterSlashings        []*sila.AttesterSlashing
	attesterSlashingsElectra []*sila.AttesterSlashingElectra
	voluntaryExits           []*sila.SignedVoluntaryExit
	syncAggregate            *sila.SyncAggregate
	execPayload              *silaenginev1.SilaPayload
	execPayloadHeader        *silaenginev1.SilaPayloadHeader
	execPayloadCapella       *silaenginev1.SilaPayloadCapella
	execPayloadHeaderCapella *silaenginev1.SilaPayloadHeaderCapella
	execPayloadDeneb         *silaenginev1.SilaPayloadDeneb
	execPayloadHeaderDeneb   *silaenginev1.SilaPayloadHeaderDeneb
	blsToSilaChanges         []*sila.SignedBLSToSilaChange
	kzgCommitments           [][]byte
	execRequests             *silaenginev1.SilaRequests
}

func Test_SignedBeaconBlock_Proto(t *testing.T) {
	f := getFields()

	t.Run("Phase0", func(t *testing.T) {
		expectedBlock := &sila.SignedBeaconBlock{
			Block: &sila.BeaconBlock{
				Slot:          128,
				ProposerIndex: 128,
				ParentRoot:    f.root[:],
				StateRoot:     f.root[:],
				Body:          bodyPbPhase0(),
			},
			Signature: f.sig[:],
		}
		block := &SignedBeaconBlock{
			version: version.Phase0,
			block: &BeaconBlock{
				version:       version.Phase0,
				slot:          128,
				proposerIndex: 128,
				parentRoot:    f.root,
				stateRoot:     f.root,
				body:          bodyPhase0(),
			},
			signature: f.sig,
		}

		result, err := block.Proto()
		require.NoError(t, err)
		resultBlock, ok := result.(*sila.SignedBeaconBlock)
		require.Equal(t, true, ok)
		resultHTR, err := resultBlock.HashTreeRoot()
		require.NoError(t, err)
		expectedHTR, err := expectedBlock.HashTreeRoot()
		require.NoError(t, err)
		assert.DeepEqual(t, expectedHTR, resultHTR)
	})
	t.Run("Altair", func(t *testing.T) {
		expectedBlock := &sila.SignedBeaconBlockAltair{
			Block: &sila.BeaconBlockAltair{
				Slot:          128,
				ProposerIndex: 128,
				ParentRoot:    f.root[:],
				StateRoot:     f.root[:],
				Body:          bodyPbAltair(),
			},
			Signature: f.sig[:],
		}
		block := &SignedBeaconBlock{
			version: version.Altair,
			block: &BeaconBlock{
				version:       version.Altair,
				slot:          128,
				proposerIndex: 128,
				parentRoot:    f.root,
				stateRoot:     f.root,
				body:          bodyAltair(),
			},
			signature: f.sig,
		}

		result, err := block.Proto()
		require.NoError(t, err)
		resultBlock, ok := result.(*sila.SignedBeaconBlockAltair)
		require.Equal(t, true, ok)
		resultHTR, err := resultBlock.HashTreeRoot()
		require.NoError(t, err)
		expectedHTR, err := expectedBlock.HashTreeRoot()
		require.NoError(t, err)
		assert.DeepEqual(t, expectedHTR, resultHTR)
	})
	t.Run("Bellatrix", func(t *testing.T) {
		expectedBlock := &sila.SignedBeaconBlockBellatrix{
			Block: &sila.BeaconBlockBellatrix{
				Slot:          128,
				ProposerIndex: 128,
				ParentRoot:    f.root[:],
				StateRoot:     f.root[:],
				Body:          bodyPbBellatrix(),
			},
			Signature: f.sig[:],
		}
		block := &SignedBeaconBlock{
			version: version.Bellatrix,
			block: &BeaconBlock{
				version:       version.Bellatrix,
				slot:          128,
				proposerIndex: 128,
				parentRoot:    f.root,
				stateRoot:     f.root,
				body:          bodyBellatrix(t),
			},
			signature: f.sig,
		}

		result, err := block.Proto()
		require.NoError(t, err)
		resultBlock, ok := result.(*sila.SignedBeaconBlockBellatrix)
		require.Equal(t, true, ok)
		resultHTR, err := resultBlock.HashTreeRoot()
		require.NoError(t, err)
		expectedHTR, err := expectedBlock.HashTreeRoot()
		require.NoError(t, err)
		assert.DeepEqual(t, expectedHTR, resultHTR)
	})
	t.Run("BellatrixBlind", func(t *testing.T) {
		expectedBlock := &sila.SignedBlindedBeaconBlockBellatrix{
			Block: &sila.BlindedBeaconBlockBellatrix{
				Slot:          128,
				ProposerIndex: 128,
				ParentRoot:    f.root[:],
				StateRoot:     f.root[:],
				Body:          bodyPbBlindedBellatrix(),
			},
			Signature: f.sig[:],
		}
		block := &SignedBeaconBlock{
			version: version.Bellatrix,
			block: &BeaconBlock{
				version:       version.Bellatrix,
				slot:          128,
				proposerIndex: 128,
				parentRoot:    f.root,
				stateRoot:     f.root,
				body:          bodyBlindedBellatrix(t),
			},
			signature: f.sig,
		}

		result, err := block.Proto()
		require.NoError(t, err)
		resultBlock, ok := result.(*sila.SignedBlindedBeaconBlockBellatrix)
		require.Equal(t, true, ok)
		resultHTR, err := resultBlock.HashTreeRoot()
		require.NoError(t, err)
		expectedHTR, err := expectedBlock.HashTreeRoot()
		require.NoError(t, err)
		assert.DeepEqual(t, expectedHTR, resultHTR)
	})
	t.Run("Capella", func(t *testing.T) {
		expectedBlock := &sila.SignedBeaconBlockCapella{
			Block: &sila.BeaconBlockCapella{
				Slot:          128,
				ProposerIndex: 128,
				ParentRoot:    f.root[:],
				StateRoot:     f.root[:],
				Body:          bodyPbCapella(),
			},
			Signature: f.sig[:],
		}
		block := &SignedBeaconBlock{
			version: version.Capella,
			block: &BeaconBlock{
				version:       version.Capella,
				slot:          128,
				proposerIndex: 128,
				parentRoot:    f.root,
				stateRoot:     f.root,
				body:          bodyCapella(t),
			},
			signature: f.sig,
		}

		result, err := block.Proto()
		require.NoError(t, err)
		resultBlock, ok := result.(*sila.SignedBeaconBlockCapella)
		require.Equal(t, true, ok)
		resultHTR, err := resultBlock.HashTreeRoot()
		require.NoError(t, err)
		expectedHTR, err := expectedBlock.HashTreeRoot()
		require.NoError(t, err)
		assert.DeepEqual(t, expectedHTR, resultHTR)
	})
	t.Run("CapellaBlind", func(t *testing.T) {
		expectedBlock := &sila.SignedBlindedBeaconBlockCapella{
			Block: &sila.BlindedBeaconBlockCapella{
				Slot:          128,
				ProposerIndex: 128,
				ParentRoot:    f.root[:],
				StateRoot:     f.root[:],
				Body:          bodyPbBlindedCapella(),
			},
			Signature: f.sig[:],
		}
		block := &SignedBeaconBlock{
			version: version.Capella,
			block: &BeaconBlock{
				version:       version.Capella,
				slot:          128,
				proposerIndex: 128,
				parentRoot:    f.root,
				stateRoot:     f.root,
				body:          bodyBlindedCapella(t),
			},
			signature: f.sig,
		}

		result, err := block.Proto()
		require.NoError(t, err)
		resultBlock, ok := result.(*sila.SignedBlindedBeaconBlockCapella)
		require.Equal(t, true, ok)
		resultHTR, err := resultBlock.HashTreeRoot()
		require.NoError(t, err)
		expectedHTR, err := expectedBlock.HashTreeRoot()
		require.NoError(t, err)
		assert.DeepEqual(t, expectedHTR, resultHTR)
	})
	t.Run("Deneb", func(t *testing.T) {
		expectedBlock := &sila.SignedBeaconBlockDeneb{
			Block: &sila.BeaconBlockDeneb{
				Slot:          128,
				ProposerIndex: 128,
				ParentRoot:    f.root[:],
				StateRoot:     f.root[:],
				Body:          bodyPbDeneb(),
			},
			Signature: f.sig[:],
		}
		block := &SignedBeaconBlock{
			version: version.Deneb,
			block: &BeaconBlock{
				version:       version.Deneb,
				slot:          128,
				proposerIndex: 128,
				parentRoot:    f.root,
				stateRoot:     f.root,
				body:          bodyDeneb(t),
			},
			signature: f.sig,
		}

		result, err := block.Proto()
		require.NoError(t, err)
		resultBlock, ok := result.(*sila.SignedBeaconBlockDeneb)
		require.Equal(t, true, ok)
		resultHTR, err := resultBlock.HashTreeRoot()
		require.NoError(t, err)
		expectedHTR, err := expectedBlock.HashTreeRoot()
		require.NoError(t, err)
		assert.DeepEqual(t, expectedHTR, resultHTR)
	})
	t.Run("DenebBlind", func(t *testing.T) {
		expectedBlock := &sila.SignedBlindedBeaconBlockDeneb{
			Message: &sila.BlindedBeaconBlockDeneb{
				Slot:          128,
				ProposerIndex: 128,
				ParentRoot:    f.root[:],
				StateRoot:     f.root[:],
				Body:          bodyPbBlindedDeneb(),
			},
			Signature: f.sig[:],
		}
		block := &SignedBeaconBlock{
			version: version.Deneb,
			block: &BeaconBlock{
				version:       version.Deneb,
				slot:          128,
				proposerIndex: 128,
				parentRoot:    f.root,
				stateRoot:     f.root,
				body:          bodyBlindedDeneb(t),
			},
			signature: f.sig,
		}

		result, err := block.Proto()
		require.NoError(t, err)
		resultBlock, ok := result.(*sila.SignedBlindedBeaconBlockDeneb)
		require.Equal(t, true, ok)
		resultHTR, err := resultBlock.HashTreeRoot()
		require.NoError(t, err)
		expectedHTR, err := expectedBlock.HashTreeRoot()
		require.NoError(t, err)
		assert.DeepEqual(t, expectedHTR, resultHTR)
	})
	t.Run("Electra", func(t *testing.T) {
		expectedBlock := &sila.SignedBeaconBlockElectra{
			Block: &sila.BeaconBlockElectra{
				Slot:          128,
				ProposerIndex: 128,
				ParentRoot:    f.root[:],
				StateRoot:     f.root[:],
				Body:          bodyPbElectra(),
			},
			Signature: f.sig[:],
		}
		block := &SignedBeaconBlock{
			version: version.Electra,
			block: &BeaconBlock{
				version:       version.Electra,
				slot:          128,
				proposerIndex: 128,
				parentRoot:    f.root,
				stateRoot:     f.root,
				body:          bodyElectra(t),
			},
			signature: f.sig,
		}

		result, err := block.Proto()
		require.NoError(t, err)
		resultBlock, ok := result.(*sila.SignedBeaconBlockElectra)
		require.Equal(t, true, ok)
		resultHTR, err := resultBlock.HashTreeRoot()
		require.NoError(t, err)
		expectedHTR, err := expectedBlock.HashTreeRoot()
		require.NoError(t, err)
		assert.DeepEqual(t, expectedHTR, resultHTR)
	})
	t.Run("ElectraBlind", func(t *testing.T) {
		expectedBlock := &sila.SignedBlindedBeaconBlockElectra{
			Message: &sila.BlindedBeaconBlockElectra{
				Slot:          128,
				ProposerIndex: 128,
				ParentRoot:    f.root[:],
				StateRoot:     f.root[:],
				Body:          bodyPbBlindedElectra(),
			},
			Signature: f.sig[:],
		}
		block := &SignedBeaconBlock{
			version: version.Electra,
			block: &BeaconBlock{
				version:       version.Electra,
				slot:          128,
				proposerIndex: 128,
				parentRoot:    f.root,
				stateRoot:     f.root,
				body:          bodyBlindedElectra(t),
			},
			signature: f.sig,
		}

		result, err := block.Proto()
		require.NoError(t, err)
		resultBlock, ok := result.(*sila.SignedBlindedBeaconBlockElectra)
		require.Equal(t, true, ok)
		resultHTR, err := resultBlock.HashTreeRoot()
		require.NoError(t, err)
		expectedHTR, err := expectedBlock.HashTreeRoot()
		require.NoError(t, err)
		assert.DeepEqual(t, expectedHTR, resultHTR)
	})
}

func Test_BeaconBlock_Proto(t *testing.T) {
	f := getFields()

	t.Run("Phase0", func(t *testing.T) {
		expectedBlock := &sila.BeaconBlock{
			Slot:          128,
			ProposerIndex: 128,
			ParentRoot:    f.root[:],
			StateRoot:     f.root[:],
			Body:          bodyPbPhase0(),
		}
		block := &BeaconBlock{
			version:       version.Phase0,
			slot:          128,
			proposerIndex: 128,
			parentRoot:    f.root,
			stateRoot:     f.root,
			body:          bodyPhase0(),
		}

		result, err := block.Proto()
		require.NoError(t, err)
		resultBlock, ok := result.(*sila.BeaconBlock)
		require.Equal(t, true, ok)
		resultHTR, err := resultBlock.HashTreeRoot()
		require.NoError(t, err)
		expectedHTR, err := expectedBlock.HashTreeRoot()
		require.NoError(t, err)
		assert.DeepEqual(t, expectedHTR, resultHTR)
	})
	t.Run("Altair", func(t *testing.T) {
		expectedBlock := &sila.BeaconBlockAltair{
			Slot:          128,
			ProposerIndex: 128,
			ParentRoot:    f.root[:],
			StateRoot:     f.root[:],
			Body:          bodyPbAltair(),
		}
		block := &BeaconBlock{
			version:       version.Altair,
			slot:          128,
			proposerIndex: 128,
			parentRoot:    f.root,
			stateRoot:     f.root,
			body:          bodyAltair(),
		}

		result, err := block.Proto()
		require.NoError(t, err)
		resultBlock, ok := result.(*sila.BeaconBlockAltair)
		require.Equal(t, true, ok)
		resultHTR, err := resultBlock.HashTreeRoot()
		require.NoError(t, err)
		expectedHTR, err := expectedBlock.HashTreeRoot()
		require.NoError(t, err)
		assert.DeepEqual(t, expectedHTR, resultHTR)
	})
	t.Run("Bellatrix", func(t *testing.T) {
		expectedBlock := &sila.BeaconBlockBellatrix{
			Slot:          128,
			ProposerIndex: 128,
			ParentRoot:    f.root[:],
			StateRoot:     f.root[:],
			Body:          bodyPbBellatrix(),
		}
		block := &BeaconBlock{
			version:       version.Bellatrix,
			slot:          128,
			proposerIndex: 128,
			parentRoot:    f.root,
			stateRoot:     f.root,
			body:          bodyBellatrix(t),
		}

		result, err := block.Proto()
		require.NoError(t, err)
		resultBlock, ok := result.(*sila.BeaconBlockBellatrix)
		require.Equal(t, true, ok)
		resultHTR, err := resultBlock.HashTreeRoot()
		require.NoError(t, err)
		expectedHTR, err := expectedBlock.HashTreeRoot()
		require.NoError(t, err)
		assert.DeepEqual(t, expectedHTR, resultHTR)
	})
	t.Run("BellatrixBlind", func(t *testing.T) {
		expectedBlock := &sila.BlindedBeaconBlockBellatrix{
			Slot:          128,
			ProposerIndex: 128,
			ParentRoot:    f.root[:],
			StateRoot:     f.root[:],
			Body:          bodyPbBlindedBellatrix(),
		}
		block := &BeaconBlock{
			version:       version.Bellatrix,
			slot:          128,
			proposerIndex: 128,
			parentRoot:    f.root,
			stateRoot:     f.root,
			body:          bodyBlindedBellatrix(t),
		}

		result, err := block.Proto()
		require.NoError(t, err)
		resultBlock, ok := result.(*sila.BlindedBeaconBlockBellatrix)
		require.Equal(t, true, ok)
		resultHTR, err := resultBlock.HashTreeRoot()
		require.NoError(t, err)
		expectedHTR, err := expectedBlock.HashTreeRoot()
		require.NoError(t, err)
		assert.DeepEqual(t, expectedHTR, resultHTR)
	})
	t.Run("Capella", func(t *testing.T) {
		expectedBlock := &sila.BeaconBlockCapella{
			Slot:          128,
			ProposerIndex: 128,
			ParentRoot:    f.root[:],
			StateRoot:     f.root[:],
			Body:          bodyPbCapella(),
		}
		block := &BeaconBlock{
			version:       version.Capella,
			slot:          128,
			proposerIndex: 128,
			parentRoot:    f.root,
			stateRoot:     f.root,
			body:          bodyCapella(t),
		}

		result, err := block.Proto()
		require.NoError(t, err)
		resultBlock, ok := result.(*sila.BeaconBlockCapella)
		require.Equal(t, true, ok)
		resultHTR, err := resultBlock.HashTreeRoot()
		require.NoError(t, err)
		expectedHTR, err := expectedBlock.HashTreeRoot()
		require.NoError(t, err)
		assert.DeepEqual(t, expectedHTR, resultHTR)
	})
	t.Run("CapellaBlind", func(t *testing.T) {
		expectedBlock := &sila.BlindedBeaconBlockCapella{
			Slot:          128,
			ProposerIndex: 128,
			ParentRoot:    f.root[:],
			StateRoot:     f.root[:],
			Body:          bodyPbBlindedCapella(),
		}
		block := &BeaconBlock{
			version:       version.Capella,
			slot:          128,
			proposerIndex: 128,
			parentRoot:    f.root,
			stateRoot:     f.root,
			body:          bodyBlindedCapella(t),
		}

		result, err := block.Proto()
		require.NoError(t, err)
		resultBlock, ok := result.(*sila.BlindedBeaconBlockCapella)
		require.Equal(t, true, ok)
		resultHTR, err := resultBlock.HashTreeRoot()
		require.NoError(t, err)
		expectedHTR, err := expectedBlock.HashTreeRoot()
		require.NoError(t, err)
		assert.DeepEqual(t, expectedHTR, resultHTR)
	})
	t.Run("Deneb", func(t *testing.T) {
		expectedBlock := &sila.BeaconBlockDeneb{
			Slot:          128,
			ProposerIndex: 128,
			ParentRoot:    f.root[:],
			StateRoot:     f.root[:],
			Body:          bodyPbDeneb(),
		}
		block := &BeaconBlock{
			version:       version.Deneb,
			slot:          128,
			proposerIndex: 128,
			parentRoot:    f.root,
			stateRoot:     f.root,
			body:          bodyDeneb(t),
		}

		result, err := block.Proto()
		require.NoError(t, err)
		resultBlock, ok := result.(*sila.BeaconBlockDeneb)
		require.Equal(t, true, ok)
		resultHTR, err := resultBlock.HashTreeRoot()
		require.NoError(t, err)
		expectedHTR, err := expectedBlock.HashTreeRoot()
		require.NoError(t, err)
		assert.DeepEqual(t, expectedHTR, resultHTR)
	})
	t.Run("DenebBlind", func(t *testing.T) {
		expectedBlock := &sila.BlindedBeaconBlockDeneb{
			Slot:          128,
			ProposerIndex: 128,
			ParentRoot:    f.root[:],
			StateRoot:     f.root[:],
			Body:          bodyPbBlindedDeneb(),
		}
		block := &BeaconBlock{
			version:       version.Deneb,
			slot:          128,
			proposerIndex: 128,
			parentRoot:    f.root,
			stateRoot:     f.root,
			body:          bodyBlindedDeneb(t),
		}

		result, err := block.Proto()
		require.NoError(t, err)
		resultBlock, ok := result.(*sila.BlindedBeaconBlockDeneb)
		require.Equal(t, true, ok)
		resultHTR, err := resultBlock.HashTreeRoot()
		require.NoError(t, err)
		expectedHTR, err := expectedBlock.HashTreeRoot()
		require.NoError(t, err)
		assert.DeepEqual(t, expectedHTR, resultHTR)
	})
	t.Run("Electra", func(t *testing.T) {
		expectedBlock := &sila.BeaconBlockElectra{
			Slot:          128,
			ProposerIndex: 128,
			ParentRoot:    f.root[:],
			StateRoot:     f.root[:],
			Body:          bodyPbElectra(),
		}
		block := &BeaconBlock{
			version:       version.Electra,
			slot:          128,
			proposerIndex: 128,
			parentRoot:    f.root,
			stateRoot:     f.root,
			body:          bodyElectra(t),
		}

		result, err := block.Proto()
		require.NoError(t, err)
		resultBlock, ok := result.(*sila.BeaconBlockElectra)
		require.Equal(t, true, ok)
		resultHTR, err := resultBlock.HashTreeRoot()
		require.NoError(t, err)
		expectedHTR, err := expectedBlock.HashTreeRoot()
		require.NoError(t, err)
		assert.DeepEqual(t, expectedHTR, resultHTR)
	})
	t.Run("ElectraBlind", func(t *testing.T) {
		expectedBlock := &sila.BlindedBeaconBlockElectra{
			Slot:          128,
			ProposerIndex: 128,
			ParentRoot:    f.root[:],
			StateRoot:     f.root[:],
			Body:          bodyPbBlindedElectra(),
		}
		block := &BeaconBlock{
			version:       version.Electra,
			slot:          128,
			proposerIndex: 128,
			parentRoot:    f.root,
			stateRoot:     f.root,
			body:          bodyBlindedElectra(t),
		}

		result, err := block.Proto()
		require.NoError(t, err)
		resultBlock, ok := result.(*sila.BlindedBeaconBlockElectra)
		require.Equal(t, true, ok)
		resultHTR, err := resultBlock.HashTreeRoot()
		require.NoError(t, err)
		expectedHTR, err := expectedBlock.HashTreeRoot()
		require.NoError(t, err)
		assert.DeepEqual(t, expectedHTR, resultHTR)
	})
}

func Test_BeaconBlockBody_Proto(t *testing.T) {
	t.Run("Phase0", func(t *testing.T) {
		expectedBody := bodyPbPhase0()
		body := bodyPhase0()

		result, err := body.Proto()
		require.NoError(t, err)
		resultBlock, ok := result.(*sila.BeaconBlockBody)
		require.Equal(t, true, ok)
		resultHTR, err := resultBlock.HashTreeRoot()
		require.NoError(t, err)
		expectedHTR, err := expectedBody.HashTreeRoot()
		require.NoError(t, err)
		assert.DeepEqual(t, expectedHTR, resultHTR)
	})
	t.Run("Altair", func(t *testing.T) {
		expectedBody := bodyPbAltair()
		body := bodyAltair()
		result, err := body.Proto()
		require.NoError(t, err)
		resultBlock, ok := result.(*sila.BeaconBlockBodyAltair)
		require.Equal(t, true, ok)
		resultHTR, err := resultBlock.HashTreeRoot()
		require.NoError(t, err)
		expectedHTR, err := expectedBody.HashTreeRoot()
		require.NoError(t, err)
		assert.DeepEqual(t, expectedHTR, resultHTR)
	})
	t.Run("Bellatrix", func(t *testing.T) {
		expectedBody := bodyPbBellatrix()
		body := bodyBellatrix(t)
		result, err := body.Proto()
		require.NoError(t, err)
		resultBlock, ok := result.(*sila.BeaconBlockBodyBellatrix)
		require.Equal(t, true, ok)
		resultHTR, err := resultBlock.HashTreeRoot()
		require.NoError(t, err)
		expectedHTR, err := expectedBody.HashTreeRoot()
		require.NoError(t, err)
		assert.DeepEqual(t, expectedHTR, resultHTR)
	})
	t.Run("BellatrixBlind", func(t *testing.T) {
		expectedBody := bodyPbBlindedBellatrix()
		body := bodyBlindedBellatrix(t)
		result, err := body.Proto()
		require.NoError(t, err)
		resultBlock, ok := result.(*sila.BlindedBeaconBlockBodyBellatrix)
		require.Equal(t, true, ok)
		resultHTR, err := resultBlock.HashTreeRoot()
		require.NoError(t, err)
		expectedHTR, err := expectedBody.HashTreeRoot()
		require.NoError(t, err)
		assert.DeepEqual(t, expectedHTR, resultHTR)
	})
	t.Run("Capella", func(t *testing.T) {
		expectedBody := bodyPbCapella()
		body := bodyCapella(t)
		result, err := body.Proto()
		require.NoError(t, err)
		resultBlock, ok := result.(*sila.BeaconBlockBodyCapella)
		require.Equal(t, true, ok)
		resultHTR, err := resultBlock.HashTreeRoot()
		require.NoError(t, err)
		expectedHTR, err := expectedBody.HashTreeRoot()
		require.NoError(t, err)
		assert.DeepEqual(t, expectedHTR, resultHTR)
	})
	t.Run("CapellaBlind", func(t *testing.T) {
		expectedBody := bodyPbBlindedCapella()
		body := bodyBlindedCapella(t)
		result, err := body.Proto()
		require.NoError(t, err)
		resultBlock, ok := result.(*sila.BlindedBeaconBlockBodyCapella)
		require.Equal(t, true, ok)
		resultHTR, err := resultBlock.HashTreeRoot()
		require.NoError(t, err)
		expectedHTR, err := expectedBody.HashTreeRoot()
		require.NoError(t, err)
		assert.DeepEqual(t, expectedHTR, resultHTR)
	})
	t.Run("Deneb", func(t *testing.T) {
		expectedBody := bodyPbDeneb()
		body := bodyDeneb(t)
		result, err := body.Proto()
		require.NoError(t, err)
		resultBlock, ok := result.(*sila.BeaconBlockBodyDeneb)
		require.Equal(t, true, ok)
		resultHTR, err := resultBlock.HashTreeRoot()
		require.NoError(t, err)
		expectedHTR, err := expectedBody.HashTreeRoot()
		require.NoError(t, err)
		assert.DeepEqual(t, expectedHTR, resultHTR)
	})
	t.Run("DenebBlind", func(t *testing.T) {
		expectedBody := bodyPbBlindedDeneb()
		body := bodyBlindedDeneb(t)
		result, err := body.Proto()
		require.NoError(t, err)
		resultBlock, ok := result.(*sila.BlindedBeaconBlockBodyDeneb)
		require.Equal(t, true, ok)
		resultHTR, err := resultBlock.HashTreeRoot()
		require.NoError(t, err)
		expectedHTR, err := expectedBody.HashTreeRoot()
		require.NoError(t, err)
		assert.DeepEqual(t, expectedHTR, resultHTR)
	})
	t.Run("Electra", func(t *testing.T) {
		expectedBody := bodyPbElectra()
		body := bodyElectra(t)
		result, err := body.Proto()
		require.NoError(t, err)
		resultBlock, ok := result.(*sila.BeaconBlockBodyElectra)
		require.Equal(t, true, ok)
		resultHTR, err := resultBlock.HashTreeRoot()
		require.NoError(t, err)
		expectedHTR, err := expectedBody.HashTreeRoot()
		require.NoError(t, err)
		assert.DeepEqual(t, expectedHTR, resultHTR)
	})
	t.Run("ElectraBlind", func(t *testing.T) {
		expectedBody := bodyPbBlindedElectra()
		body := bodyBlindedElectra(t)
		result, err := body.Proto()
		require.NoError(t, err)
		resultBlock, ok := result.(*sila.BlindedBeaconBlockBodyElectra)
		require.Equal(t, true, ok)
		resultHTR, err := resultBlock.HashTreeRoot()
		require.NoError(t, err)
		expectedHTR, err := expectedBody.HashTreeRoot()
		require.NoError(t, err)
		assert.DeepEqual(t, expectedHTR, resultHTR)
	})
	t.Run("Bellatrix - wrong payload type", func(t *testing.T) {
		body := bodyBellatrix(t)
		body.silaPayload = &silaPayloadHeader{}
		_, err := body.Proto()
		require.ErrorIs(t, err, errPayloadWrongType)
	})
	t.Run("BellatrixBlind - wrong payload type", func(t *testing.T) {
		body := bodyBlindedBellatrix(t)
		body.silaPayloadHeader = &silaPayload{}
		_, err := body.Proto()
		require.ErrorIs(t, err, errPayloadHeaderWrongType)
	})
	t.Run("Capella - wrong payload type", func(t *testing.T) {
		body := bodyCapella(t)
		body.silaPayload = &silaPayloadHeaderCapella{}
		_, err := body.Proto()
		require.ErrorIs(t, err, errPayloadWrongType)
	})
	t.Run("CapellaBlind - wrong payload type", func(t *testing.T) {
		body := bodyBlindedCapella(t)
		body.silaPayloadHeader = &silaPayloadCapella{}
		_, err := body.Proto()
		require.ErrorIs(t, err, errPayloadHeaderWrongType)
	})
	t.Run("Deneb - wrong payload type", func(t *testing.T) {
		body := bodyDeneb(t)
		body.silaPayload = &silaPayloadHeaderDeneb{}
		_, err := body.Proto()
		require.ErrorIs(t, err, errPayloadWrongType)
	})
	t.Run("DenebBlind - wrong payload type", func(t *testing.T) {
		body := bodyBlindedDeneb(t)
		body.silaPayloadHeader = &silaPayloadDeneb{}
		_, err := body.Proto()
		require.ErrorIs(t, err, errPayloadHeaderWrongType)
	})
	t.Run("Electra - wrong payload type", func(t *testing.T) {
		body := bodyElectra(t)
		body.silaPayload = &silaPayloadHeaderDeneb{}
		_, err := body.Proto()
		require.ErrorIs(t, err, errPayloadWrongType)
	})
	t.Run("ElectraBlind - wrong payload type", func(t *testing.T) {
		body := bodyBlindedElectra(t)
		body.silaPayloadHeader = &silaPayloadDeneb{}
		_, err := body.Proto()
		require.ErrorIs(t, err, errPayloadHeaderWrongType)
	})
}

func Test_initSignedBlockFromProtoPhase0(t *testing.T) {
	f := getFields()
	expectedBlock := &sila.SignedBeaconBlock{
		Block: &sila.BeaconBlock{
			Slot:          128,
			ProposerIndex: 128,
			ParentRoot:    f.root[:],
			StateRoot:     f.root[:],
			Body:          bodyPbPhase0(),
		},
		Signature: f.sig[:],
	}
	resultBlock, err := initSignedBlockFromProtoPhase0(expectedBlock)
	require.NoError(t, err)
	resultHTR, err := resultBlock.block.HashTreeRoot()
	require.NoError(t, err)
	expectedHTR, err := expectedBlock.Block.HashTreeRoot()
	require.NoError(t, err)
	assert.DeepEqual(t, expectedHTR, resultHTR)
	assert.DeepEqual(t, expectedBlock.Signature, resultBlock.signature[:])
}

func Test_initSignedBlockFromProtoAltair(t *testing.T) {
	f := getFields()
	expectedBlock := &sila.SignedBeaconBlockAltair{
		Block: &sila.BeaconBlockAltair{
			Slot:          128,
			ProposerIndex: 128,
			ParentRoot:    f.root[:],
			StateRoot:     f.root[:],
			Body:          bodyPbAltair(),
		},
		Signature: f.sig[:],
	}
	resultBlock, err := initSignedBlockFromProtoAltair(expectedBlock)
	require.NoError(t, err)
	resultHTR, err := resultBlock.block.HashTreeRoot()
	require.NoError(t, err)
	expectedHTR, err := expectedBlock.Block.HashTreeRoot()
	require.NoError(t, err)
	assert.DeepEqual(t, expectedHTR, resultHTR)
	assert.DeepEqual(t, expectedBlock.Signature, resultBlock.signature[:])
}

func Test_initSignedBlockFromProtoBellatrix(t *testing.T) {
	f := getFields()
	expectedBlock := &sila.SignedBeaconBlockBellatrix{
		Block: &sila.BeaconBlockBellatrix{
			Slot:          128,
			ProposerIndex: 128,
			ParentRoot:    f.root[:],
			StateRoot:     f.root[:],
			Body:          bodyPbBellatrix(),
		},
		Signature: f.sig[:],
	}
	resultBlock, err := initSignedBlockFromProtoBellatrix(expectedBlock)
	require.NoError(t, err)
	resultHTR, err := resultBlock.block.HashTreeRoot()
	require.NoError(t, err)
	expectedHTR, err := expectedBlock.Block.HashTreeRoot()
	require.NoError(t, err)
	assert.DeepEqual(t, expectedHTR, resultHTR)
	assert.DeepEqual(t, expectedBlock.Signature, resultBlock.signature[:])
}

func Test_initBlindedSignedBlockFromProtoBellatrix(t *testing.T) {
	f := getFields()
	expectedBlock := &sila.SignedBlindedBeaconBlockBellatrix{
		Block: &sila.BlindedBeaconBlockBellatrix{
			Slot:          128,
			ProposerIndex: 128,
			ParentRoot:    f.root[:],
			StateRoot:     f.root[:],
			Body:          bodyPbBlindedBellatrix(),
		},
		Signature: f.sig[:],
	}
	resultBlock, err := initBlindedSignedBlockFromProtoBellatrix(expectedBlock)
	require.NoError(t, err)
	resultHTR, err := resultBlock.block.HashTreeRoot()
	require.NoError(t, err)
	expectedHTR, err := expectedBlock.Block.HashTreeRoot()
	require.NoError(t, err)
	assert.DeepEqual(t, expectedHTR, resultHTR)
	assert.DeepEqual(t, expectedBlock.Signature, resultBlock.signature[:])
}

func Test_initSignedBlockFromProtoCapella(t *testing.T) {
	f := getFields()
	expectedBlock := &sila.SignedBeaconBlockCapella{
		Block: &sila.BeaconBlockCapella{
			Slot:          128,
			ProposerIndex: 128,
			ParentRoot:    f.root[:],
			StateRoot:     f.root[:],
			Body:          bodyPbCapella(),
		},
		Signature: f.sig[:],
	}
	resultBlock, err := initSignedBlockFromProtoCapella(expectedBlock)
	require.NoError(t, err)
	resultHTR, err := resultBlock.block.HashTreeRoot()
	require.NoError(t, err)
	expectedHTR, err := expectedBlock.Block.HashTreeRoot()
	require.NoError(t, err)
	assert.DeepEqual(t, expectedHTR, resultHTR)
	assert.DeepEqual(t, expectedBlock.Signature, resultBlock.signature[:])
}

func Test_initBlindedSignedBlockFromProtoCapella(t *testing.T) {
	f := getFields()
	expectedBlock := &sila.SignedBlindedBeaconBlockCapella{
		Block: &sila.BlindedBeaconBlockCapella{
			Slot:          128,
			ProposerIndex: 128,
			ParentRoot:    f.root[:],
			StateRoot:     f.root[:],
			Body:          bodyPbBlindedCapella(),
		},
		Signature: f.sig[:],
	}
	resultBlock, err := initBlindedSignedBlockFromProtoCapella(expectedBlock)
	require.NoError(t, err)
	resultHTR, err := resultBlock.block.HashTreeRoot()
	require.NoError(t, err)
	expectedHTR, err := expectedBlock.Block.HashTreeRoot()
	require.NoError(t, err)
	assert.DeepEqual(t, expectedHTR, resultHTR)
	assert.DeepEqual(t, expectedBlock.Signature, resultBlock.signature[:])
}

func Test_initSignedBlockFromProtoDeneb(t *testing.T) {
	f := getFields()
	expectedBlock := &sila.SignedBeaconBlockDeneb{
		Block: &sila.BeaconBlockDeneb{
			Slot:          128,
			ProposerIndex: 128,
			ParentRoot:    f.root[:],
			StateRoot:     f.root[:],
			Body:          bodyPbDeneb(),
		},
		Signature: f.sig[:],
	}
	resultBlock, err := initSignedBlockFromProtoDeneb(expectedBlock)
	require.NoError(t, err)
	resultHTR, err := resultBlock.block.HashTreeRoot()
	require.NoError(t, err)
	expectedHTR, err := expectedBlock.Block.HashTreeRoot()
	require.NoError(t, err)
	assert.DeepEqual(t, expectedHTR, resultHTR)
	assert.DeepEqual(t, expectedBlock.Signature, resultBlock.signature[:])
}

func Test_initBlindedSignedBlockFromProtoDeneb(t *testing.T) {
	f := getFields()
	expectedBlock := &sila.SignedBlindedBeaconBlockDeneb{
		Message: &sila.BlindedBeaconBlockDeneb{
			Slot:          128,
			ProposerIndex: 128,
			ParentRoot:    f.root[:],
			StateRoot:     f.root[:],
			Body:          bodyPbBlindedDeneb(),
		},
		Signature: f.sig[:],
	}
	resultBlock, err := initBlindedSignedBlockFromProtoDeneb(expectedBlock)
	require.NoError(t, err)
	resultHTR, err := resultBlock.block.HashTreeRoot()
	require.NoError(t, err)
	expectedHTR, err := expectedBlock.Message.HashTreeRoot()
	require.NoError(t, err)
	assert.DeepEqual(t, expectedHTR, resultHTR)
	assert.DeepEqual(t, expectedBlock.Signature, resultBlock.signature[:])
}

func Test_initSignedBlockFromProtoElectra(t *testing.T) {
	f := getFields()
	expectedBlock := &sila.SignedBeaconBlockElectra{
		Block: &sila.BeaconBlockElectra{
			Slot:          128,
			ProposerIndex: 128,
			ParentRoot:    f.root[:],
			StateRoot:     f.root[:],
			Body:          bodyPbElectra(),
		},
		Signature: f.sig[:],
	}
	resultBlock, err := initSignedBlockFromProtoElectra(expectedBlock)
	require.NoError(t, err)
	resultHTR, err := resultBlock.block.HashTreeRoot()
	require.NoError(t, err)
	expectedHTR, err := expectedBlock.Block.HashTreeRoot()
	require.NoError(t, err)
	assert.DeepEqual(t, expectedHTR, resultHTR)
	assert.DeepEqual(t, expectedBlock.Signature, resultBlock.signature[:])
}

func Test_initBlindedSignedBlockFromProtoElectra(t *testing.T) {
	f := getFields()
	expectedBlock := &sila.SignedBlindedBeaconBlockElectra{
		Message: &sila.BlindedBeaconBlockElectra{
			Slot:          128,
			ProposerIndex: 128,
			ParentRoot:    f.root[:],
			StateRoot:     f.root[:],
			Body:          bodyPbBlindedElectra(),
		},
		Signature: f.sig[:],
	}
	resultBlock, err := initBlindedSignedBlockFromProtoElectra(expectedBlock)
	require.NoError(t, err)
	resultHTR, err := resultBlock.block.HashTreeRoot()
	require.NoError(t, err)
	expectedHTR, err := expectedBlock.Message.HashTreeRoot()
	require.NoError(t, err)
	assert.DeepEqual(t, expectedHTR, resultHTR)
	assert.DeepEqual(t, expectedBlock.Signature, resultBlock.signature[:])
}

func Test_initBlockFromProtoPhase0(t *testing.T) {
	f := getFields()
	expectedBlock := &sila.BeaconBlock{
		Slot:          128,
		ProposerIndex: 128,
		ParentRoot:    f.root[:],
		StateRoot:     f.root[:],
		Body:          bodyPbPhase0(),
	}
	resultBlock, err := initBlockFromProtoPhase0(expectedBlock)
	require.NoError(t, err)
	resultHTR, err := resultBlock.HashTreeRoot()
	require.NoError(t, err)
	expectedHTR, err := expectedBlock.HashTreeRoot()
	require.NoError(t, err)
	assert.DeepEqual(t, expectedHTR, resultHTR)
}

func Test_initBlockFromProtoAltair(t *testing.T) {
	f := getFields()
	expectedBlock := &sila.BeaconBlockAltair{
		Slot:          128,
		ProposerIndex: 128,
		ParentRoot:    f.root[:],
		StateRoot:     f.root[:],
		Body:          bodyPbAltair(),
	}
	resultBlock, err := initBlockFromProtoAltair(expectedBlock)
	require.NoError(t, err)
	resultHTR, err := resultBlock.HashTreeRoot()
	require.NoError(t, err)
	expectedHTR, err := expectedBlock.HashTreeRoot()
	require.NoError(t, err)
	assert.DeepEqual(t, expectedHTR, resultHTR)
}

func Test_initBlockFromProtoBellatrix(t *testing.T) {
	f := getFields()
	expectedBlock := &sila.BeaconBlockBellatrix{
		Slot:          128,
		ProposerIndex: 128,
		ParentRoot:    f.root[:],
		StateRoot:     f.root[:],
		Body:          bodyPbBellatrix(),
	}
	resultBlock, err := initBlockFromProtoBellatrix(expectedBlock)
	require.NoError(t, err)
	resultHTR, err := resultBlock.HashTreeRoot()
	require.NoError(t, err)
	expectedHTR, err := expectedBlock.HashTreeRoot()
	require.NoError(t, err)
	assert.DeepEqual(t, expectedHTR, resultHTR)
}

func Test_initBlockFromProtoBlindedBellatrix(t *testing.T) {
	f := getFields()
	expectedBlock := &sila.BlindedBeaconBlockBellatrix{
		Slot:          128,
		ProposerIndex: 128,
		ParentRoot:    f.root[:],
		StateRoot:     f.root[:],
		Body:          bodyPbBlindedBellatrix(),
	}
	resultBlock, err := initBlindedBlockFromProtoBellatrix(expectedBlock)
	require.NoError(t, err)
	resultHTR, err := resultBlock.HashTreeRoot()
	require.NoError(t, err)
	expectedHTR, err := expectedBlock.HashTreeRoot()
	require.NoError(t, err)
	assert.DeepEqual(t, expectedHTR, resultHTR)
}

func Test_initBlockFromProtoCapella(t *testing.T) {
	f := getFields()
	expectedBlock := &sila.BeaconBlockCapella{
		Slot:          128,
		ProposerIndex: 128,
		ParentRoot:    f.root[:],
		StateRoot:     f.root[:],
		Body:          bodyPbCapella(),
	}
	resultBlock, err := initBlockFromProtoCapella(expectedBlock)
	require.NoError(t, err)
	resultHTR, err := resultBlock.HashTreeRoot()
	require.NoError(t, err)
	expectedHTR, err := expectedBlock.HashTreeRoot()
	require.NoError(t, err)
	assert.DeepEqual(t, expectedHTR, resultHTR)
}

func Test_initBlockFromProtoBlindedCapella(t *testing.T) {
	f := getFields()
	expectedBlock := &sila.BlindedBeaconBlockCapella{
		Slot:          128,
		ProposerIndex: 128,
		ParentRoot:    f.root[:],
		StateRoot:     f.root[:],
		Body:          bodyPbBlindedCapella(),
	}
	resultBlock, err := initBlindedBlockFromProtoCapella(expectedBlock)
	require.NoError(t, err)
	resultHTR, err := resultBlock.HashTreeRoot()
	require.NoError(t, err)
	expectedHTR, err := expectedBlock.HashTreeRoot()
	require.NoError(t, err)
	assert.DeepEqual(t, expectedHTR, resultHTR)
}

func Test_initBlockFromProtoDeneb(t *testing.T) {
	f := getFields()
	expectedBlock := &sila.BeaconBlockDeneb{
		Slot:          128,
		ProposerIndex: 128,
		ParentRoot:    f.root[:],
		StateRoot:     f.root[:],
		Body:          bodyPbDeneb(),
	}
	resultBlock, err := initBlockFromProtoDeneb(expectedBlock)
	require.NoError(t, err)
	resultHTR, err := resultBlock.HashTreeRoot()
	require.NoError(t, err)
	expectedHTR, err := expectedBlock.HashTreeRoot()
	require.NoError(t, err)
	assert.DeepEqual(t, expectedHTR, resultHTR)
}

func Test_initBlockFromProtoBlindedDeneb(t *testing.T) {
	f := getFields()
	expectedBlock := &sila.BlindedBeaconBlockDeneb{
		Slot:          128,
		ProposerIndex: 128,
		ParentRoot:    f.root[:],
		StateRoot:     f.root[:],
		Body:          bodyPbBlindedDeneb(),
	}
	resultBlock, err := initBlindedBlockFromProtoDeneb(expectedBlock)
	require.NoError(t, err)
	resultHTR, err := resultBlock.HashTreeRoot()
	require.NoError(t, err)
	expectedHTR, err := expectedBlock.HashTreeRoot()
	require.NoError(t, err)
	assert.DeepEqual(t, expectedHTR, resultHTR)
}

func Test_initBlockFromProtoElectra(t *testing.T) {
	f := getFields()
	expectedBlock := &sila.BeaconBlockElectra{
		Slot:          128,
		ProposerIndex: 128,
		ParentRoot:    f.root[:],
		StateRoot:     f.root[:],
		Body:          bodyPbElectra(),
	}
	resultBlock, err := initBlockFromProtoElectra(expectedBlock)
	require.NoError(t, err)
	resultHTR, err := resultBlock.HashTreeRoot()
	require.NoError(t, err)
	expectedHTR, err := expectedBlock.HashTreeRoot()
	require.NoError(t, err)
	assert.DeepEqual(t, expectedHTR, resultHTR)
}

func Test_initBlockFromProtoBlindedElectra(t *testing.T) {
	f := getFields()
	expectedBlock := &sila.BlindedBeaconBlockElectra{
		Slot:          128,
		ProposerIndex: 128,
		ParentRoot:    f.root[:],
		StateRoot:     f.root[:],
		Body:          bodyPbBlindedElectra(),
	}
	resultBlock, err := initBlindedBlockFromProtoElectra(expectedBlock)
	require.NoError(t, err)
	resultHTR, err := resultBlock.HashTreeRoot()
	require.NoError(t, err)
	expectedHTR, err := expectedBlock.HashTreeRoot()
	require.NoError(t, err)
	assert.DeepEqual(t, expectedHTR, resultHTR)
}

func Test_initBlockBodyFromProtoPhase0(t *testing.T) {
	expectedBody := bodyPbPhase0()
	resultBody, err := initBlockBodyFromProtoPhase0(expectedBody)
	require.NoError(t, err)
	resultHTR, err := resultBody.HashTreeRoot()
	require.NoError(t, err)
	expectedHTR, err := expectedBody.HashTreeRoot()
	require.NoError(t, err)
	assert.DeepEqual(t, expectedHTR, resultHTR)
}

func Test_initBlockBodyFromProtoAltair(t *testing.T) {
	expectedBody := bodyPbAltair()
	resultBody, err := initBlockBodyFromProtoAltair(expectedBody)
	require.NoError(t, err)
	resultHTR, err := resultBody.HashTreeRoot()
	require.NoError(t, err)
	expectedHTR, err := expectedBody.HashTreeRoot()
	require.NoError(t, err)
	assert.DeepEqual(t, expectedHTR, resultHTR)
}

func Test_initBlockBodyFromProtoBellatrix(t *testing.T) {
	expectedBody := bodyPbBellatrix()
	resultBody, err := initBlockBodyFromProtoBellatrix(expectedBody)
	require.NoError(t, err)
	resultHTR, err := resultBody.HashTreeRoot()
	require.NoError(t, err)
	expectedHTR, err := expectedBody.HashTreeRoot()
	require.NoError(t, err)
	assert.DeepEqual(t, expectedHTR, resultHTR)
}

func Test_initBlockBodyFromProtoBlindedBellatrix(t *testing.T) {
	expectedBody := bodyPbBlindedBellatrix()
	resultBody, err := initBlindedBlockBodyFromProtoBellatrix(expectedBody)
	require.NoError(t, err)
	resultHTR, err := resultBody.HashTreeRoot()
	require.NoError(t, err)
	expectedHTR, err := expectedBody.HashTreeRoot()
	require.NoError(t, err)
	assert.DeepEqual(t, expectedHTR, resultHTR)
}

func Test_initBlockBodyFromProtoCapella(t *testing.T) {
	expectedBody := bodyPbCapella()
	resultBody, err := initBlockBodyFromProtoCapella(expectedBody)
	require.NoError(t, err)
	resultHTR, err := resultBody.HashTreeRoot()
	require.NoError(t, err)
	expectedHTR, err := expectedBody.HashTreeRoot()
	require.NoError(t, err)
	assert.DeepEqual(t, expectedHTR, resultHTR)
}

func Test_initBlockBodyFromProtoBlindedCapella(t *testing.T) {
	expectedBody := bodyPbBlindedCapella()
	resultBody, err := initBlindedBlockBodyFromProtoCapella(expectedBody)
	require.NoError(t, err)
	resultHTR, err := resultBody.HashTreeRoot()
	require.NoError(t, err)
	expectedHTR, err := expectedBody.HashTreeRoot()
	require.NoError(t, err)
	assert.DeepEqual(t, expectedHTR, resultHTR)
}

func Test_initBlockBodyFromProtoDeneb(t *testing.T) {
	expectedBody := bodyPbDeneb()
	resultBody, err := initBlockBodyFromProtoDeneb(expectedBody)
	require.NoError(t, err)
	resultHTR, err := resultBody.HashTreeRoot()
	require.NoError(t, err)
	expectedHTR, err := expectedBody.HashTreeRoot()
	require.NoError(t, err)
	assert.DeepEqual(t, expectedHTR, resultHTR)
}

func Test_initBlockBodyFromProtoBlindedDeneb(t *testing.T) {
	expectedBody := bodyPbBlindedDeneb()
	resultBody, err := initBlindedBlockBodyFromProtoDeneb(expectedBody)
	require.NoError(t, err)
	resultHTR, err := resultBody.HashTreeRoot()
	require.NoError(t, err)
	expectedHTR, err := expectedBody.HashTreeRoot()
	require.NoError(t, err)
	assert.DeepEqual(t, expectedHTR, resultHTR)
}

func Test_initBlockBodyFromProtoElectra(t *testing.T) {
	expectedBody := bodyPbElectra()
	resultBody, err := initBlockBodyFromProtoElectra(expectedBody)
	require.NoError(t, err)
	resultHTR, err := resultBody.HashTreeRoot()
	require.NoError(t, err)
	expectedHTR, err := expectedBody.HashTreeRoot()
	require.NoError(t, err)
	assert.DeepEqual(t, expectedHTR, resultHTR)
}

func Test_initBlockBodyFromProtoBlindedElectra(t *testing.T) {
	expectedBody := bodyPbBlindedElectra()
	resultBody, err := initBlindedBlockBodyFromProtoElectra(expectedBody)
	require.NoError(t, err)
	resultHTR, err := resultBody.HashTreeRoot()
	require.NoError(t, err)
	expectedHTR, err := expectedBody.HashTreeRoot()
	require.NoError(t, err)
	assert.DeepEqual(t, expectedHTR, resultHTR)
}

func bodyPbPhase0() *sila.BeaconBlockBody {
	f := getFields()
	return &sila.BeaconBlockBody{
		RandaoReveal: f.sig[:],
		SilaData: &sila.SilaData{
			DepositRoot:  f.root[:],
			DepositCount: 128,
			BlockHash:    f.root[:],
		},
		Graffiti:          f.root[:],
		ProposerSlashings: f.proposerSlashings,
		AttesterSlashings: f.attesterSlashings,
		Attestations:      f.atts,
		Deposits:          f.deposits,
		VoluntaryExits:    f.voluntaryExits,
	}
}

func bodyPbAltair() *sila.BeaconBlockBodyAltair {
	f := getFields()
	return &sila.BeaconBlockBodyAltair{
		RandaoReveal: f.sig[:],
		SilaData: &sila.SilaData{
			DepositRoot:  f.root[:],
			DepositCount: 128,
			BlockHash:    f.root[:],
		},
		Graffiti:          f.root[:],
		ProposerSlashings: f.proposerSlashings,
		AttesterSlashings: f.attesterSlashings,
		Attestations:      f.atts,
		Deposits:          f.deposits,
		VoluntaryExits:    f.voluntaryExits,
		SyncAggregate:     f.syncAggregate,
	}
}

func bodyPbBellatrix() *sila.BeaconBlockBodyBellatrix {
	f := getFields()
	return &sila.BeaconBlockBodyBellatrix{
		RandaoReveal: f.sig[:],
		SilaData: &sila.SilaData{
			DepositRoot:  f.root[:],
			DepositCount: 128,
			BlockHash:    f.root[:],
		},
		Graffiti:          f.root[:],
		ProposerSlashings: f.proposerSlashings,
		AttesterSlashings: f.attesterSlashings,
		Attestations:      f.atts,
		Deposits:          f.deposits,
		VoluntaryExits:    f.voluntaryExits,
		SyncAggregate:     f.syncAggregate,
		SilaPayload:       f.execPayload,
	}
}

func bodyPbBlindedBellatrix() *sila.BlindedBeaconBlockBodyBellatrix {
	f := getFields()
	return &sila.BlindedBeaconBlockBodyBellatrix{
		RandaoReveal: f.sig[:],
		SilaData: &sila.SilaData{
			DepositRoot:  f.root[:],
			DepositCount: 128,
			BlockHash:    f.root[:],
		},
		Graffiti:          f.root[:],
		ProposerSlashings: f.proposerSlashings,
		AttesterSlashings: f.attesterSlashings,
		Attestations:      f.atts,
		Deposits:          f.deposits,
		VoluntaryExits:    f.voluntaryExits,
		SyncAggregate:     f.syncAggregate,
		SilaPayloadHeader: f.execPayloadHeader,
	}
}

func bodyPbCapella() *sila.BeaconBlockBodyCapella {
	f := getFields()
	return &sila.BeaconBlockBodyCapella{
		RandaoReveal: f.sig[:],
		SilaData: &sila.SilaData{
			DepositRoot:  f.root[:],
			DepositCount: 128,
			BlockHash:    f.root[:],
		},
		Graffiti:          f.root[:],
		ProposerSlashings: f.proposerSlashings,
		AttesterSlashings: f.attesterSlashings,
		Attestations:      f.atts,
		Deposits:          f.deposits,
		VoluntaryExits:    f.voluntaryExits,
		SyncAggregate:     f.syncAggregate,
		SilaPayload:       f.execPayloadCapella,
		BlsToSilaChanges:  f.blsToSilaChanges,
	}
}

func bodyPbBlindedCapella() *sila.BlindedBeaconBlockBodyCapella {
	f := getFields()
	return &sila.BlindedBeaconBlockBodyCapella{
		RandaoReveal: f.sig[:],
		SilaData: &sila.SilaData{
			DepositRoot:  f.root[:],
			DepositCount: 128,
			BlockHash:    f.root[:],
		},
		Graffiti:          f.root[:],
		ProposerSlashings: f.proposerSlashings,
		AttesterSlashings: f.attesterSlashings,
		Attestations:      f.atts,
		Deposits:          f.deposits,
		VoluntaryExits:    f.voluntaryExits,
		SyncAggregate:     f.syncAggregate,
		SilaPayloadHeader: f.execPayloadHeaderCapella,
		BlsToSilaChanges:  f.blsToSilaChanges,
	}
}

func bodyPbDeneb() *sila.BeaconBlockBodyDeneb {
	f := getFields()
	return &sila.BeaconBlockBodyDeneb{
		RandaoReveal: f.sig[:],
		SilaData: &sila.SilaData{
			DepositRoot:  f.root[:],
			DepositCount: 128,
			BlockHash:    f.root[:],
		},
		Graffiti:           f.root[:],
		ProposerSlashings:  f.proposerSlashings,
		AttesterSlashings:  f.attesterSlashings,
		Attestations:       f.atts,
		Deposits:           f.deposits,
		VoluntaryExits:     f.voluntaryExits,
		SyncAggregate:      f.syncAggregate,
		SilaPayload:        f.execPayloadDeneb,
		BlsToSilaChanges:   f.blsToSilaChanges,
		BlobKzgCommitments: f.kzgCommitments,
	}
}

func bodyPbBlindedDeneb() *sila.BlindedBeaconBlockBodyDeneb {
	f := getFields()
	return &sila.BlindedBeaconBlockBodyDeneb{
		RandaoReveal: f.sig[:],
		SilaData: &sila.SilaData{
			DepositRoot:  f.root[:],
			DepositCount: 128,
			BlockHash:    f.root[:],
		},
		Graffiti:           f.root[:],
		ProposerSlashings:  f.proposerSlashings,
		AttesterSlashings:  f.attesterSlashings,
		Attestations:       f.atts,
		Deposits:           f.deposits,
		VoluntaryExits:     f.voluntaryExits,
		SyncAggregate:      f.syncAggregate,
		SilaPayloadHeader:  f.execPayloadHeaderDeneb,
		BlsToSilaChanges:   f.blsToSilaChanges,
		BlobKzgCommitments: f.kzgCommitments,
	}
}

func bodyPbElectra() *sila.BeaconBlockBodyElectra {
	f := getFields()
	return &sila.BeaconBlockBodyElectra{
		RandaoReveal: f.sig[:],
		SilaData: &sila.SilaData{
			DepositRoot:  f.root[:],
			DepositCount: 128,
			BlockHash:    f.root[:],
		},
		Graffiti:           f.root[:],
		ProposerSlashings:  f.proposerSlashings,
		AttesterSlashings:  f.attesterSlashingsElectra,
		Attestations:       f.attsElectra,
		Deposits:           f.deposits,
		VoluntaryExits:     f.voluntaryExits,
		SyncAggregate:      f.syncAggregate,
		SilaPayload:        f.execPayloadDeneb,
		BlsToSilaChanges:   f.blsToSilaChanges,
		BlobKzgCommitments: f.kzgCommitments,
		SilaRequests:       f.execRequests,
	}
}

func bodyPbBlindedElectra() *sila.BlindedBeaconBlockBodyElectra {
	f := getFields()
	return &sila.BlindedBeaconBlockBodyElectra{
		RandaoReveal: f.sig[:],
		SilaData: &sila.SilaData{
			DepositRoot:  f.root[:],
			DepositCount: 128,
			BlockHash:    f.root[:],
		},
		Graffiti:           f.root[:],
		ProposerSlashings:  f.proposerSlashings,
		AttesterSlashings:  f.attesterSlashingsElectra,
		Attestations:       f.attsElectra,
		Deposits:           f.deposits,
		VoluntaryExits:     f.voluntaryExits,
		SyncAggregate:      f.syncAggregate,
		SilaPayloadHeader:  f.execPayloadHeaderDeneb,
		BlsToSilaChanges:   f.blsToSilaChanges,
		BlobKzgCommitments: f.kzgCommitments,
		SilaRequests:       f.execRequests,
	}
}

func bodyPhase0() *BeaconBlockBody {
	f := getFields()
	return &BeaconBlockBody{
		version:      version.Phase0,
		randaoReveal: f.sig,
		silaexecData: &sila.SilaData{
			DepositRoot:  f.root[:],
			DepositCount: 128,
			BlockHash:    f.root[:],
		},
		graffiti:          f.root,
		proposerSlashings: f.proposerSlashings,
		attesterSlashings: f.attesterSlashings,
		attestations:      f.atts,
		deposits:          f.deposits,
		voluntaryExits:    f.voluntaryExits,
	}
}

func bodyAltair() *BeaconBlockBody {
	f := getFields()
	return &BeaconBlockBody{
		version:      version.Altair,
		randaoReveal: f.sig,
		silaexecData: &sila.SilaData{
			DepositRoot:  f.root[:],
			DepositCount: 128,
			BlockHash:    f.root[:],
		},
		graffiti:          f.root,
		proposerSlashings: f.proposerSlashings,
		attesterSlashings: f.attesterSlashings,
		attestations:      f.atts,
		deposits:          f.deposits,
		voluntaryExits:    f.voluntaryExits,
		syncAggregate:     f.syncAggregate,
	}
}

func bodyBellatrix(t *testing.T) *BeaconBlockBody {
	f := getFields()
	p, err := WrappedSilaPayload(f.execPayload)
	require.NoError(t, err)
	return &BeaconBlockBody{
		version:      version.Bellatrix,
		randaoReveal: f.sig,
		silaexecData: &sila.SilaData{
			DepositRoot:  f.root[:],
			DepositCount: 128,
			BlockHash:    f.root[:],
		},
		graffiti:          f.root,
		proposerSlashings: f.proposerSlashings,
		attesterSlashings: f.attesterSlashings,
		attestations:      f.atts,
		deposits:          f.deposits,
		voluntaryExits:    f.voluntaryExits,
		syncAggregate:     f.syncAggregate,
		silaPayload:       p,
	}
}

func bodyBlindedBellatrix(t *testing.T) *BeaconBlockBody {
	f := getFields()
	ph, err := WrappedSilaPayloadHeader(f.execPayloadHeader)
	require.NoError(t, err)
	return &BeaconBlockBody{
		version:      version.Bellatrix,
		randaoReveal: f.sig,
		silaexecData: &sila.SilaData{
			DepositRoot:  f.root[:],
			DepositCount: 128,
			BlockHash:    f.root[:],
		},
		graffiti:          f.root,
		proposerSlashings: f.proposerSlashings,
		attesterSlashings: f.attesterSlashings,
		attestations:      f.atts,
		deposits:          f.deposits,
		voluntaryExits:    f.voluntaryExits,
		syncAggregate:     f.syncAggregate,
		silaPayloadHeader: ph,
	}
}

func bodyCapella(t *testing.T) *BeaconBlockBody {
	f := getFields()
	p, err := WrappedSilaPayloadCapella(f.execPayloadCapella)
	require.NoError(t, err)
	return &BeaconBlockBody{
		version:      version.Capella,
		randaoReveal: f.sig,
		silaexecData: &sila.SilaData{
			DepositRoot:  f.root[:],
			DepositCount: 128,
			BlockHash:    f.root[:],
		},
		graffiti:          f.root,
		proposerSlashings: f.proposerSlashings,
		attesterSlashings: f.attesterSlashings,
		attestations:      f.atts,
		deposits:          f.deposits,
		voluntaryExits:    f.voluntaryExits,
		syncAggregate:     f.syncAggregate,
		silaPayload:       p,
		blsToSilaChanges:  f.blsToSilaChanges,
	}
}

func bodyBlindedCapella(t *testing.T) *BeaconBlockBody {
	f := getFields()
	ph, err := WrappedSilaPayloadHeaderCapella(f.execPayloadHeaderCapella)
	require.NoError(t, err)
	return &BeaconBlockBody{
		version:      version.Capella,
		randaoReveal: f.sig,
		silaexecData: &sila.SilaData{
			DepositRoot:  f.root[:],
			DepositCount: 128,
			BlockHash:    f.root[:],
		},
		graffiti:          f.root,
		proposerSlashings: f.proposerSlashings,
		attesterSlashings: f.attesterSlashings,
		attestations:      f.atts,
		deposits:          f.deposits,
		voluntaryExits:    f.voluntaryExits,
		syncAggregate:     f.syncAggregate,
		silaPayloadHeader: ph,
		blsToSilaChanges:  f.blsToSilaChanges,
	}
}

func bodyDeneb(t *testing.T) *BeaconBlockBody {
	f := getFields()
	p, err := WrappedSilaPayloadDeneb(f.execPayloadDeneb)
	require.NoError(t, err)
	return &BeaconBlockBody{
		version:      version.Deneb,
		randaoReveal: f.sig,
		silaexecData: &sila.SilaData{
			DepositRoot:  f.root[:],
			DepositCount: 128,
			BlockHash:    f.root[:],
		},
		graffiti:           f.root,
		proposerSlashings:  f.proposerSlashings,
		attesterSlashings:  f.attesterSlashings,
		attestations:       f.atts,
		deposits:           f.deposits,
		voluntaryExits:     f.voluntaryExits,
		syncAggregate:      f.syncAggregate,
		silaPayload:        p,
		blsToSilaChanges:   f.blsToSilaChanges,
		blobKzgCommitments: f.kzgCommitments,
	}
}

func bodyBlindedDeneb(t *testing.T) *BeaconBlockBody {
	f := getFields()
	ph, err := WrappedSilaPayloadHeaderDeneb(f.execPayloadHeaderDeneb)
	require.NoError(t, err)
	return &BeaconBlockBody{
		version:      version.Deneb,
		randaoReveal: f.sig,
		silaexecData: &sila.SilaData{
			DepositRoot:  f.root[:],
			DepositCount: 128,
			BlockHash:    f.root[:],
		},
		graffiti:           f.root,
		proposerSlashings:  f.proposerSlashings,
		attesterSlashings:  f.attesterSlashings,
		attestations:       f.atts,
		deposits:           f.deposits,
		voluntaryExits:     f.voluntaryExits,
		syncAggregate:      f.syncAggregate,
		silaPayloadHeader:  ph,
		blsToSilaChanges:   f.blsToSilaChanges,
		blobKzgCommitments: f.kzgCommitments,
	}
}

func bodyElectra(t *testing.T) *BeaconBlockBody {
	f := getFields()
	p, err := WrappedSilaPayloadDeneb(f.execPayloadDeneb)
	require.NoError(t, err)
	return &BeaconBlockBody{
		version:      version.Electra,
		randaoReveal: f.sig,
		silaexecData: &sila.SilaData{
			DepositRoot:  f.root[:],
			DepositCount: 128,
			BlockHash:    f.root[:],
		},
		graffiti:                 f.root,
		proposerSlashings:        f.proposerSlashings,
		attesterSlashingsElectra: f.attesterSlashingsElectra,
		attestationsElectra:      f.attsElectra,
		deposits:                 f.deposits,
		voluntaryExits:           f.voluntaryExits,
		syncAggregate:            f.syncAggregate,
		silaPayload:              p,
		blsToSilaChanges:         f.blsToSilaChanges,
		blobKzgCommitments:       f.kzgCommitments,
		silaRequests:             f.execRequests,
	}
}

func bodyBlindedElectra(t *testing.T) *BeaconBlockBody {
	f := getFields()
	ph, err := WrappedSilaPayloadHeaderDeneb(f.execPayloadHeaderDeneb)
	require.NoError(t, err)
	return &BeaconBlockBody{
		version:      version.Electra,
		randaoReveal: f.sig,
		silaexecData: &sila.SilaData{
			DepositRoot:  f.root[:],
			DepositCount: 128,
			BlockHash:    f.root[:],
		},
		graffiti:                 f.root,
		proposerSlashings:        f.proposerSlashings,
		attesterSlashingsElectra: f.attesterSlashingsElectra,
		attestationsElectra:      f.attsElectra,
		deposits:                 f.deposits,
		voluntaryExits:           f.voluntaryExits,
		syncAggregate:            f.syncAggregate,
		silaPayloadHeader:        ph,
		blsToSilaChanges:         f.blsToSilaChanges,
		blobKzgCommitments:       f.kzgCommitments,
		silaRequests:             f.execRequests,
	}
}

func TestSignedBeaconBlockProtoGloas(t *testing.T) {
	payload := []*sila.PayloadAttestation{{Signature: []byte{0x01}}}
	bid := &sila.SignedSilaPayloadBid{Signature: []byte{0x02}}
	sb := &SignedBeaconBlock{
		version: version.Gloas,
		block: &BeaconBlock{
			version: version.Gloas,
			body: &BeaconBlockBody{
				version:              version.Gloas,
				payloadAttestations:  payload,
				signedSilaPayloadBid: bid,
			},
		},
	}

	msg, err := sb.Proto()
	require.NoError(t, err)
	gloas, ok := msg.(*sila.SignedBeaconBlockGloas)
	require.Equal(t, true, ok)
	require.DeepEqual(t, payload, gloas.Block.Body.PayloadAttestations)
	require.DeepEqual(t, bid, gloas.Block.Body.SignedSilaPayloadBid)
}

func TestInitSignedBlockFromProtoGloas(t *testing.T) {
	bits := bitfield.NewBitvector512()
	bits.SetBitAt(0, true)
	pb := &sila.SignedBeaconBlockGloas{
		Block: &sila.BeaconBlockGloas{
			Body: &sila.BeaconBlockBodyGloas{
				PayloadAttestations: []*sila.PayloadAttestation{
					{
						AggregationBits: bits,
						Signature:       []byte{0x01},
					},
				},
				SignedSilaPayloadBid: &sila.SignedSilaPayloadBid{Signature: []byte{0x02}},
			},
		},
		Signature: []byte{0x03},
	}

	sb, err := initSignedBlockFromProtoGloas(pb)
	require.NoError(t, err)
	require.Equal(t, version.Gloas, sb.Version())

	gotPayload, err := sb.Block().Body().PayloadAttestations()
	require.NoError(t, err)
	require.Equal(t, 1, len(gotPayload))
	require.DeepEqual(t, pb.Block.Body.PayloadAttestations, gotPayload)

	gotBid, err := sb.Block().Body().SignedSilaPayloadBid()
	require.NoError(t, err)
	require.DeepEqual(t, pb.Block.Body.SignedSilaPayloadBid, gotBid)
}

func getFields() fields {
	b20 := make([]byte, 20)
	b48 := make([]byte, 48)
	b256 := make([]byte, 256)
	var root [32]byte
	var sig [96]byte
	b20[0], b20[5], b20[10] = 'q', 'u', 'x'
	b48[0], b48[5], b48[10] = 'b', 'a', 'r'
	b256[0], b256[5], b256[10] = 'x', 'y', 'z'
	root[0], root[5], root[10] = 'a', 'b', 'c'
	sig[0], sig[5], sig[10] = 'd', 'e', 'f'
	deposits := make([]*sila.Deposit, 16)
	for i := range deposits {
		deposits[i] = &sila.Deposit{}
		deposits[i].Proof = make([][]byte, 33)
		for j := range deposits[i].Proof {
			deposits[i].Proof[j] = root[:]
		}
		deposits[i].Data = &sila.Deposit_Data{
			PublicKey:             b48,
			WithdrawalCredentials: root[:],
			Amount:                128,
			Signature:             sig[:],
		}
	}

	attBits := bitfield.NewBitlist(1)
	committeeBits := primitives.NewAttestationCommitteeBits()
	atts := make([]*sila.Attestation, params.BeaconConfig().MaxAttestations)
	for i := range atts {
		atts[i] = &sila.Attestation{}
		atts[i].Signature = sig[:]
		atts[i].AggregationBits = attBits
		atts[i].Data = &sila.AttestationData{
			Slot:            128,
			CommitteeIndex:  128,
			BeaconBlockRoot: root[:],
			Source: &sila.Checkpoint{
				Epoch: 128,
				Root:  root[:],
			},
			Target: &sila.Checkpoint{
				Epoch: 128,
				Root:  root[:],
			},
		}
	}
	attsElectra := make([]*sila.AttestationElectra, params.BeaconConfig().MaxAttestationsElectra)
	for i := range attsElectra {
		attsElectra[i] = &sila.AttestationElectra{}
		attsElectra[i].Signature = sig[:]
		attsElectra[i].AggregationBits = attBits
		attsElectra[i].CommitteeBits = committeeBits
		attsElectra[i].Data = &sila.AttestationData{
			Slot:            128,
			CommitteeIndex:  128,
			BeaconBlockRoot: root[:],
			Source: &sila.Checkpoint{
				Epoch: 128,
				Root:  root[:],
			},
			Target: &sila.Checkpoint{
				Epoch: 128,
				Root:  root[:],
			},
		}
	}

	proposerSlashing := &sila.ProposerSlashing{
		Header_1: &sila.SignedBeaconBlockHeader{
			Header: &sila.BeaconBlockHeader{
				Slot:          128,
				ProposerIndex: 128,
				ParentRoot:    root[:],
				StateRoot:     root[:],
				BodyRoot:      root[:],
			},
			Signature: sig[:],
		},
		Header_2: &sila.SignedBeaconBlockHeader{
			Header: &sila.BeaconBlockHeader{
				Slot:          128,
				ProposerIndex: 128,
				ParentRoot:    root[:],
				StateRoot:     root[:],
				BodyRoot:      root[:],
			},
			Signature: sig[:],
		},
	}
	attesterSlashing := &sila.AttesterSlashing{
		Attestation_1: &sila.IndexedAttestation{
			AttestingIndices: []uint64{1, 2, 8},
			Data: &sila.AttestationData{
				Slot:            128,
				CommitteeIndex:  128,
				BeaconBlockRoot: root[:],
				Source: &sila.Checkpoint{
					Epoch: 128,
					Root:  root[:],
				},
				Target: &sila.Checkpoint{
					Epoch: 128,
					Root:  root[:],
				},
			},
			Signature: sig[:],
		},
		Attestation_2: &sila.IndexedAttestation{
			AttestingIndices: []uint64{1, 2, 8},
			Data: &sila.AttestationData{
				Slot:            128,
				CommitteeIndex:  128,
				BeaconBlockRoot: root[:],
				Source: &sila.Checkpoint{
					Epoch: 128,
					Root:  root[:],
				},
				Target: &sila.Checkpoint{
					Epoch: 128,
					Root:  root[:],
				},
			},
			Signature: sig[:],
		},
	}
	attesterSlashingElectra := &sila.AttesterSlashingElectra{
		Attestation_1: &sila.IndexedAttestationElectra{
			AttestingIndices: []uint64{1, 2, 8},
			Data: &sila.AttestationData{
				Slot:            128,
				CommitteeIndex:  128,
				BeaconBlockRoot: root[:],
				Source: &sila.Checkpoint{
					Epoch: 128,
					Root:  root[:],
				},
				Target: &sila.Checkpoint{
					Epoch: 128,
					Root:  root[:],
				},
			},
			Signature: sig[:],
		},
		Attestation_2: &sila.IndexedAttestationElectra{
			AttestingIndices: []uint64{1, 2, 8},
			Data: &sila.AttestationData{
				Slot:            128,
				CommitteeIndex:  128,
				BeaconBlockRoot: root[:],
				Source: &sila.Checkpoint{
					Epoch: 128,
					Root:  root[:],
				},
				Target: &sila.Checkpoint{
					Epoch: 128,
					Root:  root[:],
				},
			},
			Signature: sig[:],
		},
	}
	voluntaryExit := &sila.SignedVoluntaryExit{
		Exit: &sila.VoluntaryExit{
			Epoch:          128,
			ValidatorIndex: 128,
		},
		Signature: sig[:],
	}
	syncCommitteeBits := bitfield.NewBitvector512()
	syncCommitteeBits.SetBitAt(1, true)
	syncCommitteeBits.SetBitAt(2, true)
	syncCommitteeBits.SetBitAt(8, true)
	syncAggregate := &sila.SyncAggregate{
		SyncCommitteeBits:      syncCommitteeBits,
		SyncCommitteeSignature: sig[:],
	}
	execPayload := &silaenginev1.SilaPayload{
		ParentHash:    root[:],
		FeeRecipient:  b20,
		StateRoot:     root[:],
		ReceiptsRoot:  root[:],
		LogsBloom:     b256,
		PrevRandao:    root[:],
		BlockNumber:   128,
		GasLimit:      128,
		GasUsed:       128,
		Timestamp:     128,
		ExtraData:     root[:],
		BaseFeePerGas: root[:],
		BlockHash:     root[:],
		Transactions: [][]byte{
			[]byte("transaction1"),
			[]byte("transaction2"),
			[]byte("transaction8"),
		},
	}
	execPayloadHeader := &silaenginev1.SilaPayloadHeader{
		ParentHash:       root[:],
		FeeRecipient:     b20,
		StateRoot:        root[:],
		ReceiptsRoot:     root[:],
		LogsBloom:        b256,
		PrevRandao:       root[:],
		BlockNumber:      128,
		GasLimit:         128,
		GasUsed:          128,
		Timestamp:        128,
		ExtraData:        root[:],
		BaseFeePerGas:    root[:],
		BlockHash:        root[:],
		TransactionsRoot: root[:],
	}
	execPayloadCapella := &silaenginev1.SilaPayloadCapella{
		ParentHash:    root[:],
		FeeRecipient:  b20,
		StateRoot:     root[:],
		ReceiptsRoot:  root[:],
		LogsBloom:     b256,
		PrevRandao:    root[:],
		BlockNumber:   128,
		GasLimit:      128,
		GasUsed:       128,
		Timestamp:     128,
		ExtraData:     root[:],
		BaseFeePerGas: root[:],
		BlockHash:     root[:],
		Transactions: [][]byte{
			[]byte("transaction1"),
			[]byte("transaction2"),
			[]byte("transaction8"),
		},
		Withdrawals: []*silaenginev1.Withdrawal{
			{
				Index:   128,
				Address: b20,
				Amount:  128,
			},
		},
	}
	execPayloadHeaderCapella := &silaenginev1.SilaPayloadHeaderCapella{
		ParentHash:       root[:],
		FeeRecipient:     b20,
		StateRoot:        root[:],
		ReceiptsRoot:     root[:],
		LogsBloom:        b256,
		PrevRandao:       root[:],
		BlockNumber:      128,
		GasLimit:         128,
		GasUsed:          128,
		Timestamp:        128,
		ExtraData:        root[:],
		BaseFeePerGas:    root[:],
		BlockHash:        root[:],
		TransactionsRoot: root[:],
		WithdrawalsRoot:  root[:],
	}
	blsToSilaChanges := []*sila.SignedBLSToSilaChange{{
		Message: &sila.BLSToSilaChange{
			ValidatorIndex: 128,
			FromBlsPubkey:  b48,
			ToSilaAddress:  b20,
		},
		Signature: sig[:],
	}}

	execPayloadDeneb := &silaenginev1.SilaPayloadDeneb{
		ParentHash:    root[:],
		FeeRecipient:  b20,
		StateRoot:     root[:],
		ReceiptsRoot:  root[:],
		LogsBloom:     b256,
		PrevRandao:    root[:],
		BlockNumber:   128,
		GasLimit:      128,
		GasUsed:       128,
		Timestamp:     128,
		ExtraData:     root[:],
		BaseFeePerGas: root[:],
		BlockHash:     root[:],
		Transactions: [][]byte{
			[]byte("transaction1"),
			[]byte("transaction2"),
			[]byte("transaction8"),
		},
		Withdrawals: []*silaenginev1.Withdrawal{
			{
				Index:   128,
				Address: b20,
				Amount:  128,
			},
		},
		BlobGasUsed:   128,
		ExcessBlobGas: 128,
	}
	execPayloadHeaderDeneb := &silaenginev1.SilaPayloadHeaderDeneb{
		ParentHash:       root[:],
		FeeRecipient:     b20,
		StateRoot:        root[:],
		ReceiptsRoot:     root[:],
		LogsBloom:        b256,
		PrevRandao:       root[:],
		BlockNumber:      128,
		GasLimit:         128,
		GasUsed:          128,
		Timestamp:        128,
		ExtraData:        root[:],
		BaseFeePerGas:    root[:],
		BlockHash:        root[:],
		TransactionsRoot: root[:],
		WithdrawalsRoot:  root[:],
		BlobGasUsed:      128,
		ExcessBlobGas:    128,
	}

	kzgCommitments := [][]byte{
		bytesutil.PadTo([]byte{123}, 48),
		bytesutil.PadTo([]byte{223}, 48),
		bytesutil.PadTo([]byte{183}, 48),
		bytesutil.PadTo([]byte{143}, 48),
	}

	execRequests := &silaenginev1.SilaRequests{
		Deposits: []*silaenginev1.DepositRequest{{
			Pubkey:                b48,
			WithdrawalCredentials: root[:],
			Amount:                128,
			Signature:             sig[:],
			Index:                 128,
		}},
		Withdrawals: []*silaenginev1.WithdrawalRequest{{
			SourceAddress:   b20,
			ValidatorPubkey: b48,
			Amount:          128,
		}},
		Consolidations: []*silaenginev1.ConsolidationRequest{{
			SourceAddress: b20,
			SourcePubkey:  b48,
			TargetPubkey:  b48,
		}},
	}

	return fields{
		root:                     root,
		sig:                      sig,
		deposits:                 deposits,
		atts:                     atts,
		attsElectra:              attsElectra,
		proposerSlashings:        []*sila.ProposerSlashing{proposerSlashing},
		attesterSlashings:        []*sila.AttesterSlashing{attesterSlashing},
		attesterSlashingsElectra: []*sila.AttesterSlashingElectra{attesterSlashingElectra},
		voluntaryExits:           []*sila.SignedVoluntaryExit{voluntaryExit},
		syncAggregate:            syncAggregate,
		execPayload:              execPayload,
		execPayloadHeader:        execPayloadHeader,
		execPayloadCapella:       execPayloadCapella,
		execPayloadHeaderCapella: execPayloadHeaderCapella,
		execPayloadDeneb:         execPayloadDeneb,
		execPayloadHeaderDeneb:   execPayloadHeaderDeneb,
		blsToSilaChanges:         blsToSilaChanges,
		kzgCommitments:           kzgCommitments,
		execRequests:             execRequests,
	}
}
