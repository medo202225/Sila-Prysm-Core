package blocks

import (
	"bytes"
	"errors"
	"testing"

	fieldparams "github.com/sila-chain/Sila-Consensus-Core/v7/config/fieldparams"
	"github.com/sila-chain/Sila-Consensus-Core/v7/encoding/bytesutil"
	sila "github.com/sila-chain/Sila-Consensus-Core/v7/proto/sila/v1alpha1"
	silaenginev1 "github.com/sila-chain/Sila-Consensus-Core/v7/proto/silaengine/v1"
	"github.com/sila-chain/Sila-Consensus-Core/v7/runtime/version"
	"github.com/sila-chain/Sila-Consensus-Core/v7/testing/assert"
	"github.com/sila-chain/Sila-Consensus-Core/v7/testing/require"
)

func Test_NewSignedBeaconBlock(t *testing.T) {
	t.Run("GenericSignedBeaconBlock_Phase0", func(t *testing.T) {
		pb := &sila.GenericSignedBeaconBlock_Phase0{
			Phase0: &sila.SignedBeaconBlock{
				Block: &sila.BeaconBlock{
					Body: &sila.BeaconBlockBody{}}}}
		b, err := NewSignedBeaconBlock(pb)
		require.NoError(t, err)
		assert.Equal(t, version.Phase0, b.Version())
	})
	t.Run("SignedBeaconBlock", func(t *testing.T) {
		pb := &sila.SignedBeaconBlock{
			Block: &sila.BeaconBlock{
				Body: &sila.BeaconBlockBody{}}}
		b, err := NewSignedBeaconBlock(pb)
		require.NoError(t, err)
		assert.Equal(t, version.Phase0, b.Version())
	})
	t.Run("GenericSignedBeaconBlock_Altair", func(t *testing.T) {
		pb := &sila.GenericSignedBeaconBlock_Altair{
			Altair: &sila.SignedBeaconBlockAltair{
				Block: &sila.BeaconBlockAltair{
					Body: &sila.BeaconBlockBodyAltair{}}}}
		b, err := NewSignedBeaconBlock(pb)
		require.NoError(t, err)
		assert.Equal(t, version.Altair, b.Version())
	})
	t.Run("SignedBeaconBlockAltair", func(t *testing.T) {
		pb := &sila.SignedBeaconBlockAltair{
			Block: &sila.BeaconBlockAltair{
				Body: &sila.BeaconBlockBodyAltair{}}}
		b, err := NewSignedBeaconBlock(pb)
		require.NoError(t, err)
		assert.Equal(t, version.Altair, b.Version())
	})
	t.Run("GenericSignedBeaconBlock_Bellatrix", func(t *testing.T) {
		pb := &sila.GenericSignedBeaconBlock_Bellatrix{
			Bellatrix: &sila.SignedBeaconBlockBellatrix{
				Block: &sila.BeaconBlockBellatrix{
					Body: &sila.BeaconBlockBodyBellatrix{}}}}
		b, err := NewSignedBeaconBlock(pb)
		require.NoError(t, err)
		assert.Equal(t, version.Bellatrix, b.Version())
	})
	t.Run("SignedBeaconBlockBellatrix", func(t *testing.T) {
		pb := &sila.SignedBeaconBlockBellatrix{
			Block: &sila.BeaconBlockBellatrix{
				Body: &sila.BeaconBlockBodyBellatrix{}}}
		b, err := NewSignedBeaconBlock(pb)
		require.NoError(t, err)
		assert.Equal(t, version.Bellatrix, b.Version())
	})
	t.Run("GenericSignedBeaconBlock_BlindedBellatrix", func(t *testing.T) {
		pb := &sila.GenericSignedBeaconBlock_BlindedBellatrix{
			BlindedBellatrix: &sila.SignedBlindedBeaconBlockBellatrix{
				Block: &sila.BlindedBeaconBlockBellatrix{
					Body: &sila.BlindedBeaconBlockBodyBellatrix{}}}}
		b, err := NewSignedBeaconBlock(pb)
		require.NoError(t, err)
		assert.Equal(t, version.Bellatrix, b.Version())
		assert.Equal(t, true, b.IsBlinded())
	})
	t.Run("SignedBlindedBeaconBlockBellatrix", func(t *testing.T) {
		pb := &sila.SignedBlindedBeaconBlockBellatrix{
			Block: &sila.BlindedBeaconBlockBellatrix{
				Body: &sila.BlindedBeaconBlockBodyBellatrix{}}}
		b, err := NewSignedBeaconBlock(pb)
		require.NoError(t, err)
		assert.Equal(t, version.Bellatrix, b.Version())
		assert.Equal(t, true, b.IsBlinded())
	})
	t.Run("GenericSignedBeaconBlock_Capella", func(t *testing.T) {
		pb := &sila.GenericSignedBeaconBlock_Capella{
			Capella: &sila.SignedBeaconBlockCapella{
				Block: &sila.BeaconBlockCapella{
					Body: &sila.BeaconBlockBodyCapella{}}}}
		b, err := NewSignedBeaconBlock(pb)
		require.NoError(t, err)
		assert.Equal(t, version.Capella, b.Version())
	})
	t.Run("SignedBeaconBlockCapella", func(t *testing.T) {
		pb := &sila.SignedBeaconBlockCapella{
			Block: &sila.BeaconBlockCapella{
				Body: &sila.BeaconBlockBodyCapella{}}}
		b, err := NewSignedBeaconBlock(pb)
		require.NoError(t, err)
		assert.Equal(t, version.Capella, b.Version())
	})
	t.Run("GenericSignedBeaconBlock_BlindedCapella", func(t *testing.T) {
		pb := &sila.GenericSignedBeaconBlock_BlindedCapella{
			BlindedCapella: &sila.SignedBlindedBeaconBlockCapella{
				Block: &sila.BlindedBeaconBlockCapella{
					Body: &sila.BlindedBeaconBlockBodyCapella{}}}}
		b, err := NewSignedBeaconBlock(pb)
		require.NoError(t, err)
		assert.Equal(t, version.Capella, b.Version())
		assert.Equal(t, true, b.IsBlinded())
	})
	t.Run("SignedBlindedBeaconBlockCapella", func(t *testing.T) {
		pb := &sila.SignedBlindedBeaconBlockCapella{
			Block: &sila.BlindedBeaconBlockCapella{
				Body: &sila.BlindedBeaconBlockBodyCapella{}}}
		b, err := NewSignedBeaconBlock(pb)
		require.NoError(t, err)
		assert.Equal(t, version.Capella, b.Version())
		assert.Equal(t, true, b.IsBlinded())
	})
	t.Run("GenericSignedBeaconBlock_Deneb", func(t *testing.T) {
		pb := &sila.GenericSignedBeaconBlock_Deneb{
			Deneb: &sila.SignedBeaconBlockContentsDeneb{
				Block: &sila.SignedBeaconBlockDeneb{Block: &sila.BeaconBlockDeneb{
					Body: &sila.BeaconBlockBodyDeneb{},
				}},
			},
		}
		b, err := NewSignedBeaconBlock(pb)
		require.NoError(t, err)
		assert.Equal(t, version.Deneb, b.Version())
	})
	t.Run("SignedBeaconBlockDeneb", func(t *testing.T) {
		pb := &sila.SignedBeaconBlockDeneb{
			Block: &sila.BeaconBlockDeneb{
				Body: &sila.BeaconBlockBodyDeneb{}}}
		b, err := NewSignedBeaconBlock(pb)
		require.NoError(t, err)
		assert.Equal(t, version.Deneb, b.Version())
	})
	t.Run("SignedBlindedBeaconBlockDeneb", func(t *testing.T) {
		pb := &sila.SignedBlindedBeaconBlockDeneb{
			Message: &sila.BlindedBeaconBlockDeneb{
				Body: &sila.BlindedBeaconBlockBodyDeneb{}}}
		b, err := NewSignedBeaconBlock(pb)
		require.NoError(t, err)
		assert.Equal(t, version.Deneb, b.Version())
		assert.Equal(t, true, b.IsBlinded())
	})
	t.Run("GenericSignedBeaconBlock_BlindedDeneb", func(t *testing.T) {
		pb := &sila.GenericSignedBeaconBlock_BlindedDeneb{
			BlindedDeneb: &sila.SignedBlindedBeaconBlockDeneb{
				Message: &sila.BlindedBeaconBlockDeneb{
					Body: &sila.BlindedBeaconBlockBodyDeneb{},
				}}}
		b, err := NewSignedBeaconBlock(pb)
		require.NoError(t, err)
		assert.Equal(t, version.Deneb, b.Version())
		assert.Equal(t, true, b.IsBlinded())
	})
	t.Run("SignedBeaconBlockGloas", func(t *testing.T) {
		pb := &sila.SignedBeaconBlockGloas{
			Block: &sila.BeaconBlockGloas{
				Body: &sila.BeaconBlockBodyGloas{},
			},
			Signature: []byte("sig"),
		}
		b, err := NewSignedBeaconBlock(pb)
		require.NoError(t, err)
		assert.Equal(t, version.Gloas, b.Version())
		assert.Equal(t, false, b.IsBlinded())
	})
	t.Run("GenericSignedBeaconBlock_Gloas", func(t *testing.T) {
		pb := &sila.GenericSignedBeaconBlock_Gloas{
			Gloas: &sila.SignedBeaconBlockGloas{
				Block: &sila.BeaconBlockGloas{
					Body: &sila.BeaconBlockBodyGloas{},
				},
				Signature: []byte("sig"),
			},
		}
		b, err := NewSignedBeaconBlock(pb)
		require.NoError(t, err)
		assert.Equal(t, version.Gloas, b.Version())
		assert.Equal(t, false, b.IsBlinded())
	})
	t.Run("nil", func(t *testing.T) {
		_, err := NewSignedBeaconBlock(nil)
		assert.ErrorContains(t, "received nil object", err)
	})
	t.Run("unsupported type", func(t *testing.T) {
		_, err := NewSignedBeaconBlock(&bytes.Reader{})
		assert.ErrorContains(t, "unable to create block from type *bytes.Reader", err)
	})
}

func Test_NewBeaconBlock(t *testing.T) {
	t.Run("GenericBeaconBlock_Phase0", func(t *testing.T) {
		pb := &sila.GenericBeaconBlock_Phase0{Phase0: &sila.BeaconBlock{Body: &sila.BeaconBlockBody{}}}
		b, err := NewBeaconBlock(pb)
		require.NoError(t, err)
		assert.Equal(t, version.Phase0, b.Version())
	})
	t.Run("BeaconBlock", func(t *testing.T) {
		pb := &sila.BeaconBlock{Body: &sila.BeaconBlockBody{}}
		b, err := NewBeaconBlock(pb)
		require.NoError(t, err)
		assert.Equal(t, version.Phase0, b.Version())
	})
	t.Run("GenericBeaconBlock_Altair", func(t *testing.T) {
		pb := &sila.GenericBeaconBlock_Altair{Altair: &sila.BeaconBlockAltair{Body: &sila.BeaconBlockBodyAltair{}}}
		b, err := NewBeaconBlock(pb)
		require.NoError(t, err)
		assert.Equal(t, version.Altair, b.Version())
	})
	t.Run("BeaconBlockAltair", func(t *testing.T) {
		pb := &sila.BeaconBlockAltair{Body: &sila.BeaconBlockBodyAltair{}}
		b, err := NewBeaconBlock(pb)
		require.NoError(t, err)
		assert.Equal(t, version.Altair, b.Version())
	})
	t.Run("GenericBeaconBlock_Bellatrix", func(t *testing.T) {
		pb := &sila.GenericBeaconBlock_Bellatrix{Bellatrix: &sila.BeaconBlockBellatrix{Body: &sila.BeaconBlockBodyBellatrix{}}}
		b, err := NewBeaconBlock(pb)
		require.NoError(t, err)
		assert.Equal(t, version.Bellatrix, b.Version())
	})
	t.Run("BeaconBlockBellatrix", func(t *testing.T) {
		pb := &sila.BeaconBlockBellatrix{Body: &sila.BeaconBlockBodyBellatrix{}}
		b, err := NewBeaconBlock(pb)
		require.NoError(t, err)
		assert.Equal(t, version.Bellatrix, b.Version())
	})
	t.Run("GenericBeaconBlock_BlindedBellatrix", func(t *testing.T) {
		pb := &sila.GenericBeaconBlock_BlindedBellatrix{BlindedBellatrix: &sila.BlindedBeaconBlockBellatrix{Body: &sila.BlindedBeaconBlockBodyBellatrix{}}}
		b, err := NewBeaconBlock(pb)
		require.NoError(t, err)
		assert.Equal(t, version.Bellatrix, b.Version())
		assert.Equal(t, true, b.IsBlinded())
	})
	t.Run("BlindedBeaconBlockBellatrix", func(t *testing.T) {
		pb := &sila.BlindedBeaconBlockBellatrix{Body: &sila.BlindedBeaconBlockBodyBellatrix{}}
		b, err := NewBeaconBlock(pb)
		require.NoError(t, err)
		assert.Equal(t, version.Bellatrix, b.Version())
		assert.Equal(t, true, b.IsBlinded())
	})
	t.Run("GenericBeaconBlock_Capella", func(t *testing.T) {
		pb := &sila.GenericBeaconBlock_Capella{Capella: &sila.BeaconBlockCapella{Body: &sila.BeaconBlockBodyCapella{}}}
		b, err := NewBeaconBlock(pb)
		require.NoError(t, err)
		assert.Equal(t, version.Capella, b.Version())
	})
	t.Run("BeaconBlockCapella", func(t *testing.T) {
		pb := &sila.BeaconBlockCapella{Body: &sila.BeaconBlockBodyCapella{}}
		b, err := NewBeaconBlock(pb)
		require.NoError(t, err)
		assert.Equal(t, version.Capella, b.Version())
	})
	t.Run("GenericBeaconBlock_BlindedCapella", func(t *testing.T) {
		pb := &sila.GenericBeaconBlock_BlindedCapella{BlindedCapella: &sila.BlindedBeaconBlockCapella{Body: &sila.BlindedBeaconBlockBodyCapella{}}}
		b, err := NewBeaconBlock(pb)
		require.NoError(t, err)
		assert.Equal(t, version.Capella, b.Version())
		assert.Equal(t, true, b.IsBlinded())
	})
	t.Run("BlindedBeaconBlockCapella", func(t *testing.T) {
		pb := &sila.BlindedBeaconBlockCapella{Body: &sila.BlindedBeaconBlockBodyCapella{}}
		b, err := NewBeaconBlock(pb)
		require.NoError(t, err)
		assert.Equal(t, version.Capella, b.Version())
		assert.Equal(t, true, b.IsBlinded())
	})
	t.Run("GenericBeaconBlock_Deneb", func(t *testing.T) {
		pb := &sila.GenericBeaconBlock_Deneb{Deneb: &sila.BeaconBlockContentsDeneb{Block: &sila.BeaconBlockDeneb{
			Body: &sila.BeaconBlockBodyDeneb{},
		}}}
		b, err := NewBeaconBlock(pb)
		require.NoError(t, err)
		assert.Equal(t, version.Deneb, b.Version())
	})
	t.Run("BeaconBlockDeneb", func(t *testing.T) {
		pb := &sila.BeaconBlockDeneb{Body: &sila.BeaconBlockBodyDeneb{}}
		b, err := NewBeaconBlock(pb)
		require.NoError(t, err)
		assert.Equal(t, version.Deneb, b.Version())
	})
	t.Run("BlindedBeaconBlockDeneb", func(t *testing.T) {
		pb := &sila.BlindedBeaconBlockDeneb{Body: &sila.BlindedBeaconBlockBodyDeneb{}}
		b, err := NewBeaconBlock(pb)
		require.NoError(t, err)
		assert.Equal(t, version.Deneb, b.Version())
		assert.Equal(t, true, b.IsBlinded())
	})
	t.Run("GenericBeaconBlock_BlindedDeneb", func(t *testing.T) {
		pb := &sila.GenericBeaconBlock_BlindedDeneb{BlindedDeneb: &sila.BlindedBeaconBlockDeneb{Body: &sila.BlindedBeaconBlockBodyDeneb{}}}
		b, err := NewBeaconBlock(pb)
		require.NoError(t, err)
		assert.Equal(t, version.Deneb, b.Version())
		assert.Equal(t, true, b.IsBlinded())
	})
	t.Run("BeaconBlockGloas", func(t *testing.T) {
		pb := &sila.BeaconBlockGloas{Body: &sila.BeaconBlockBodyGloas{}}
		b, err := NewBeaconBlock(pb)
		require.NoError(t, err)
		assert.Equal(t, version.Gloas, b.Version())
		assert.Equal(t, false, b.IsBlinded())
	})
	t.Run("nil", func(t *testing.T) {
		_, err := NewBeaconBlock(nil)
		assert.ErrorContains(t, "received nil object", err)
	})
	t.Run("unsupported type", func(t *testing.T) {
		_, err := NewBeaconBlock(&bytes.Reader{})
		assert.ErrorContains(t, "unable to create block from type *bytes.Reader", err)
	})
}

func Test_NewBeaconBlockBody(t *testing.T) {
	t.Run("BeaconBlockBody", func(t *testing.T) {
		pb := &sila.BeaconBlockBody{}
		i, err := NewBeaconBlockBody(pb)
		require.NoError(t, err)
		b, ok := i.(*BeaconBlockBody)
		require.Equal(t, true, ok)
		assert.Equal(t, version.Phase0, b.version)
	})
	t.Run("BeaconBlockBodyAltair", func(t *testing.T) {
		pb := &sila.BeaconBlockBodyAltair{}
		i, err := NewBeaconBlockBody(pb)
		require.NoError(t, err)
		b, ok := i.(*BeaconBlockBody)
		require.Equal(t, true, ok)
		assert.Equal(t, version.Altair, b.version)
	})
	t.Run("BeaconBlockBodyBellatrix", func(t *testing.T) {
		pb := &sila.BeaconBlockBodyBellatrix{}
		i, err := NewBeaconBlockBody(pb)
		require.NoError(t, err)
		b, ok := i.(*BeaconBlockBody)
		require.Equal(t, true, ok)
		assert.Equal(t, version.Bellatrix, b.version)
	})
	t.Run("BlindedBeaconBlockBodyBellatrix", func(t *testing.T) {
		pb := &sila.BlindedBeaconBlockBodyBellatrix{}
		i, err := NewBeaconBlockBody(pb)
		require.NoError(t, err)
		b, ok := i.(*BeaconBlockBody)
		require.Equal(t, true, ok)
		assert.Equal(t, version.Bellatrix, b.version)
		assert.Equal(t, true, b.IsBlinded())
	})
	t.Run("BeaconBlockBodyCapella", func(t *testing.T) {
		pb := &sila.BeaconBlockBodyCapella{}
		i, err := NewBeaconBlockBody(pb)
		require.NoError(t, err)
		b, ok := i.(*BeaconBlockBody)
		require.Equal(t, true, ok)
		assert.Equal(t, version.Capella, b.version)
	})
	t.Run("BlindedBeaconBlockBodyCapella", func(t *testing.T) {
		pb := &sila.BlindedBeaconBlockBodyCapella{}
		i, err := NewBeaconBlockBody(pb)
		require.NoError(t, err)
		b, ok := i.(*BeaconBlockBody)
		require.Equal(t, true, ok)
		assert.Equal(t, version.Capella, b.version)
		assert.Equal(t, true, b.IsBlinded())
	})
	t.Run("BeaconBlockBodyDeneb", func(t *testing.T) {
		pb := &sila.BeaconBlockBodyDeneb{}
		i, err := NewBeaconBlockBody(pb)
		require.NoError(t, err)
		b, ok := i.(*BeaconBlockBody)
		require.Equal(t, true, ok)
		assert.Equal(t, version.Deneb, b.version)
	})
	t.Run("BlindedBeaconBlockBodyDeneb", func(t *testing.T) {
		pb := &sila.BlindedBeaconBlockBodyDeneb{}
		i, err := NewBeaconBlockBody(pb)
		require.NoError(t, err)
		b, ok := i.(*BeaconBlockBody)
		require.Equal(t, true, ok)
		assert.Equal(t, version.Deneb, b.version)
		assert.Equal(t, true, b.IsBlinded())
	})
	t.Run("BeaconBlockBodyGloas", func(t *testing.T) {
		pb := &sila.BeaconBlockBodyGloas{}
		i, err := NewBeaconBlockBody(pb)
		require.NoError(t, err)
		b, ok := i.(*BeaconBlockBody)
		require.Equal(t, true, ok)
		assert.Equal(t, version.Gloas, b.version)
		assert.Equal(t, false, b.IsBlinded())
	})
	t.Run("nil", func(t *testing.T) {
		_, err := NewBeaconBlockBody(nil)
		assert.ErrorContains(t, "received nil object", err)
	})
	t.Run("unsupported type", func(t *testing.T) {
		_, err := NewBeaconBlockBody(&bytes.Reader{})
		assert.ErrorContains(t, "unable to create block body from type *bytes.Reader", err)
	})
}

func Test_BuildSignedBeaconBlock(t *testing.T) {
	sig := bytesutil.ToBytes96([]byte("signature"))
	t.Run("Phase0", func(t *testing.T) {
		b := &BeaconBlock{version: version.Phase0, body: &BeaconBlockBody{version: version.Phase0}}
		sb, err := BuildSignedBeaconBlock(b, sig[:])
		require.NoError(t, err)
		assert.DeepEqual(t, sig, sb.Signature())
		assert.Equal(t, version.Phase0, sb.Version())
	})
	t.Run("Altair", func(t *testing.T) {
		b := &BeaconBlock{version: version.Altair, body: &BeaconBlockBody{version: version.Altair}}
		sb, err := BuildSignedBeaconBlock(b, sig[:])
		require.NoError(t, err)
		assert.DeepEqual(t, sig, sb.Signature())
		assert.Equal(t, version.Altair, sb.Version())
	})
	t.Run("Bellatrix", func(t *testing.T) {
		b := &BeaconBlock{version: version.Bellatrix, body: &BeaconBlockBody{version: version.Bellatrix}}
		sb, err := BuildSignedBeaconBlock(b, sig[:])
		require.NoError(t, err)
		assert.DeepEqual(t, sig, sb.Signature())
		assert.Equal(t, version.Bellatrix, sb.Version())
	})
	t.Run("BellatrixBlind", func(t *testing.T) {
		b := &BeaconBlock{version: version.Bellatrix, body: &BeaconBlockBody{version: version.Bellatrix}}
		sb, err := BuildSignedBeaconBlock(b, sig[:])
		require.NoError(t, err)
		assert.DeepEqual(t, sig, sb.Signature())
		assert.Equal(t, version.Bellatrix, sb.Version())
		assert.Equal(t, true, sb.IsBlinded())
	})
	t.Run("Capella", func(t *testing.T) {
		b := &BeaconBlock{version: version.Capella, body: &BeaconBlockBody{version: version.Capella}}
		sb, err := BuildSignedBeaconBlock(b, sig[:])
		require.NoError(t, err)
		assert.DeepEqual(t, sig, sb.Signature())
		assert.Equal(t, version.Capella, sb.Version())
	})
	t.Run("CapellaBlind", func(t *testing.T) {
		b := &BeaconBlock{version: version.Capella, body: &BeaconBlockBody{version: version.Capella}}
		sb, err := BuildSignedBeaconBlock(b, sig[:])
		require.NoError(t, err)
		assert.DeepEqual(t, sig, sb.Signature())
		assert.Equal(t, version.Capella, sb.Version())
		assert.Equal(t, true, sb.IsBlinded())
	})
	t.Run("Deneb", func(t *testing.T) {
		b := &BeaconBlock{version: version.Deneb, body: &BeaconBlockBody{version: version.Deneb}}
		sb, err := BuildSignedBeaconBlock(b, sig[:])
		require.NoError(t, err)
		assert.DeepEqual(t, sig, sb.Signature())
		assert.Equal(t, version.Deneb, sb.Version())
	})
	t.Run("DenebBlind", func(t *testing.T) {
		b := &BeaconBlock{version: version.Deneb, body: &BeaconBlockBody{version: version.Deneb}}
		sb, err := BuildSignedBeaconBlock(b, sig[:])
		require.NoError(t, err)
		assert.DeepEqual(t, sig, sb.Signature())
		assert.Equal(t, version.Deneb, sb.Version())
		assert.Equal(t, true, sb.IsBlinded())
	})
	t.Run("Gloas", func(t *testing.T) {
		b := &BeaconBlock{version: version.Gloas, body: &BeaconBlockBody{version: version.Gloas}}
		sb, err := BuildSignedBeaconBlock(b, sig[:])
		require.NoError(t, err)
		assert.DeepEqual(t, sig, sb.Signature())
		assert.Equal(t, version.Gloas, sb.Version())
		assert.Equal(t, false, sb.IsBlinded())
	})
}

func TestBuildSignedBeaconBlockFromSilaPayload(t *testing.T) {
	t.Run("nil block check", func(t *testing.T) {
		_, err := BuildSignedBeaconBlockFromSilaPayload(nil, nil)
		require.ErrorIs(t, ErrNilSignedBeaconBlock, err)
	})
	t.Run("not blinded payload", func(t *testing.T) {
		altairBlock := &sila.SignedBeaconBlockAltair{
			Block: &sila.BeaconBlockAltair{
				Body: &sila.BeaconBlockBodyAltair{}}}
		blk, err := NewSignedBeaconBlock(altairBlock)
		require.NoError(t, err)
		_, err = BuildSignedBeaconBlockFromSilaPayload(blk, nil)
		require.Equal(t, true, errors.Is(err, errNonBlindedSignedBeaconBlock))
	})
	t.Run("payload header root and payload root mismatch", func(t *testing.T) {
		blockHash := bytesutil.Bytes32(1)
		payload := &silaenginev1.SilaPayload{
			ParentHash:    make([]byte, fieldparams.RootLength),
			FeeRecipient:  make([]byte, 20),
			StateRoot:     make([]byte, fieldparams.RootLength),
			ReceiptsRoot:  make([]byte, fieldparams.RootLength),
			LogsBloom:     make([]byte, 256),
			PrevRandao:    make([]byte, fieldparams.RootLength),
			BaseFeePerGas: make([]byte, fieldparams.RootLength),
			BlockHash:     blockHash,
			Transactions:  make([][]byte, 0),
		}
		wrapped, err := WrappedSilaPayload(payload)
		require.NoError(t, err)
		header, err := PayloadToHeader(wrapped)
		require.NoError(t, err)
		blindedBlock := &sila.SignedBlindedBeaconBlockBellatrix{
			Block: &sila.BlindedBeaconBlockBellatrix{
				Body: &sila.BlindedBeaconBlockBodyBellatrix{}}}

		// Modify the header.
		header.GasUsed += 1
		blindedBlock.Block.Body.SilaPayloadHeader = header

		blk, err := NewSignedBeaconBlock(blindedBlock)
		require.NoError(t, err)
		_, err = BuildSignedBeaconBlockFromSilaPayload(blk, payload)
		require.ErrorContains(t, "roots do not match", err)
	})
	t.Run("ok", func(t *testing.T) {
		payload := &silaenginev1.SilaPayload{
			ParentHash:    make([]byte, fieldparams.RootLength),
			FeeRecipient:  make([]byte, 20),
			StateRoot:     make([]byte, fieldparams.RootLength),
			ReceiptsRoot:  make([]byte, fieldparams.RootLength),
			LogsBloom:     make([]byte, 256),
			PrevRandao:    make([]byte, fieldparams.RootLength),
			BaseFeePerGas: make([]byte, fieldparams.RootLength),
			BlockHash:     make([]byte, fieldparams.RootLength),
			Transactions:  make([][]byte, 0),
		}
		wrapped, err := WrappedSilaPayload(payload)
		require.NoError(t, err)
		header, err := PayloadToHeader(wrapped)
		require.NoError(t, err)
		blindedBlock := &sila.SignedBlindedBeaconBlockBellatrix{
			Block: &sila.BlindedBeaconBlockBellatrix{
				Body: &sila.BlindedBeaconBlockBodyBellatrix{}}}
		blindedBlock.Block.Body.SilaPayloadHeader = header

		blk, err := NewSignedBeaconBlock(blindedBlock)
		require.NoError(t, err)
		builtBlock, err := BuildSignedBeaconBlockFromSilaPayload(blk, payload)
		require.NoError(t, err)

		got, err := builtBlock.Block().Body().SilaData()
		require.NoError(t, err)
		require.DeepEqual(t, payload, got.Proto())
	})
	t.Run("deneb", func(t *testing.T) {
		payload := &silaenginev1.SilaPayloadDeneb{
			ParentHash:    make([]byte, fieldparams.RootLength),
			FeeRecipient:  make([]byte, 20),
			StateRoot:     make([]byte, fieldparams.RootLength),
			ReceiptsRoot:  make([]byte, fieldparams.RootLength),
			LogsBloom:     make([]byte, 256),
			PrevRandao:    make([]byte, fieldparams.RootLength),
			BaseFeePerGas: make([]byte, fieldparams.RootLength),
			BlockHash:     make([]byte, fieldparams.RootLength),
			Transactions:  make([][]byte, 0),
			ExcessBlobGas: 123,
			BlobGasUsed:   321,
		}
		wrapped, err := WrappedSilaPayloadDeneb(payload)
		require.NoError(t, err)
		header, err := PayloadToHeaderDeneb(wrapped)
		require.NoError(t, err)
		blindedBlock := &sila.SignedBlindedBeaconBlockDeneb{
			Message: &sila.BlindedBeaconBlockDeneb{
				Body: &sila.BlindedBeaconBlockBodyDeneb{}}}
		blindedBlock.Message.Body.SilaPayloadHeader = header

		blk, err := NewSignedBeaconBlock(blindedBlock)
		require.NoError(t, err)
		builtBlock, err := BuildSignedBeaconBlockFromSilaPayload(blk, payload)
		require.NoError(t, err)

		got, err := builtBlock.Block().Body().SilaData()
		require.NoError(t, err)
		require.DeepEqual(t, payload, got.Proto())
		require.DeepEqual(t, uint64(123), payload.ExcessBlobGas)
		require.DeepEqual(t, uint64(321), payload.BlobGasUsed)
	})
	t.Run("gloas execution unsupported", func(t *testing.T) {
		base := &SignedBeaconBlock{
			version: version.Gloas,
			block:   &BeaconBlock{version: version.Gloas, body: &BeaconBlockBody{version: version.Gloas}},
		}
		blinded := &testBlindedSignedBeaconBlock{SignedBeaconBlock: base}
		_, err := BuildSignedBeaconBlockFromSilaPayload(blinded, nil)
		require.ErrorContains(t, "Execution is not supported for gloas", err)
	})
}

type testBlindedSignedBeaconBlock struct {
	*SignedBeaconBlock
}

func (b *testBlindedSignedBeaconBlock) IsBlinded() bool {
	return true
}
