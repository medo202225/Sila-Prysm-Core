package kv

import (
	"context"

	"github.com/sila-chain/Sila-Consensus-Core/v7/consensus-types/primitives"
	"github.com/sila-chain/Sila-Consensus-Core/v7/encoding/bytesutil"
	"github.com/sila-chain/Sila-Consensus-Core/v7/monitoring/tracing/trace"
	silapb "github.com/sila-chain/Sila-Consensus-Core/v7/proto/sila/v1alpha1"
	"github.com/golang/snappy"
	"github.com/pkg/errors"
	bolt "go.etcd.io/bbolt"
)

// SaveSilaPayloadEnvelope blinds and saves a signed sila payload envelope keyed by
// beacon block root. The envelope is stored in blinded form: the full sila payload is replaced
// with its block hash. The full payload can later be retrieved from the EL via
// silaEngine_getPayloadBodiesByHash.
// A secondary index from BlockHash → BeaconBlockRoot is maintained so that
// envelopes can be looked up by execution block hash.
func (s *Store) SaveSilaPayloadEnvelope(ctx context.Context, env *silapb.SignedSilaPayloadEnvelope) error {
	_, span := trace.StartSpan(ctx, "BeaconDB.SaveSilaPayloadEnvelope")
	defer span.End()

	if env == nil || env.Message == nil || env.Message.Payload == nil {
		return errors.New("cannot save nil sila payload envelope")
	}

	blockRoot := bytesutil.ToBytes32(env.Message.BeaconBlockRoot)
	blockHash := bytesutil.ToBytes32(env.Message.Payload.BlockHash)
	blinded := blindEnvelope(env)

	enc, err := encodeBlindedEnvelope(blinded)
	if err != nil {
		return err
	}

	return s.db.Update(func(tx *bolt.Tx) error {
		if err := tx.Bucket(silaPayloadEnvelopesBucket).Put(blockRoot[:], enc); err != nil {
			return err
		}
		return tx.Bucket(silaPayloadEnvelopeBlockHashBucket).Put(blockHash[:], blockRoot[:])
	})
}

// SilaPayloadEnvelope retrieves the blinded signed sila payload envelope by beacon block root.
func (s *Store) SilaPayloadEnvelope(ctx context.Context, blockRoot [32]byte) (*silapb.SignedBlindedSilaPayloadEnvelope, error) {
	_, span := trace.StartSpan(ctx, "BeaconDB.SilaPayloadEnvelope")
	defer span.End()

	var enc []byte
	if err := s.db.View(func(tx *bolt.Tx) error {
		bkt := tx.Bucket(silaPayloadEnvelopesBucket)
		enc = bkt.Get(blockRoot[:])
		return nil
	}); err != nil {
		return nil, err
	}
	if enc == nil {
		return nil, errors.Wrap(ErrNotFound, "sila payload envelope not found")
	}
	return decodeBlindedEnvelope(enc)
}

// SilaPayloadEnvelopeByBlockHash retrieves the blinded signed sila payload envelope
// by execution block hash. It uses the secondary BlockHash → BeaconBlockRoot index and then
// fetches the envelope from the primary bucket.
func (s *Store) SilaPayloadEnvelopeByBlockHash(ctx context.Context, blockHash [32]byte) (*silapb.SignedBlindedSilaPayloadEnvelope, error) {
	_, span := trace.StartSpan(ctx, "BeaconDB.SilaPayloadEnvelopeByBlockHash")
	defer span.End()

	var enc []byte
	if err := s.db.View(func(tx *bolt.Tx) error {
		blockRoot := tx.Bucket(silaPayloadEnvelopeBlockHashBucket).Get(blockHash[:])
		if blockRoot == nil {
			return nil
		}
		enc = tx.Bucket(silaPayloadEnvelopesBucket).Get(blockRoot)
		return nil
	}); err != nil {
		return nil, err
	}
	if enc == nil {
		return nil, errors.Wrap(ErrNotFound, "sila payload envelope not found for block hash")
	}
	return decodeBlindedEnvelope(enc)
}

// HasSilaPayloadEnvelope checks whether an sila payload envelope exists for the given beacon block root.
func (s *Store) HasSilaPayloadEnvelope(ctx context.Context, blockRoot [32]byte) bool {
	_, span := trace.StartSpan(ctx, "BeaconDB.HasSilaPayloadEnvelope")
	defer span.End()

	var exists bool
	if err := s.db.View(func(tx *bolt.Tx) error {
		bkt := tx.Bucket(silaPayloadEnvelopesBucket)
		exists = bkt.Get(blockRoot[:]) != nil
		return nil
	}); err != nil {
		return false
	}
	return exists
}

// DeleteSilaPayloadEnvelope removes a signed sila payload envelope by beacon block root
// and cleans up the BlockHash index entry.
func (s *Store) DeleteSilaPayloadEnvelope(ctx context.Context, blockRoot [32]byte) error {
	_, span := trace.StartSpan(ctx, "BeaconDB.DeleteSilaPayloadEnvelope")
	defer span.End()

	return s.db.Update(func(tx *bolt.Tx) error {
		bkt := tx.Bucket(silaPayloadEnvelopesBucket)
		// Read the existing entry to find the BlockHash for index cleanup.
		enc := bkt.Get(blockRoot[:])
		if enc != nil {
			blinded, err := decodeBlindedEnvelope(enc)
			if err == nil && blinded.Message != nil {
				blockHash := bytesutil.ToBytes32(blinded.Message.BlockHash)
				if err := tx.Bucket(silaPayloadEnvelopeBlockHashBucket).Delete(blockHash[:]); err != nil {
					return err
				}
			}
		}
		return bkt.Delete(blockRoot[:])
	})
}

// blindEnvelope converts a full signed envelope to its blinded form by replacing
// the sila payload with its block hash. This avoids computing the expensive
// payload hash tree root on the critical path.
func blindEnvelope(env *silapb.SignedSilaPayloadEnvelope) *silapb.SignedBlindedSilaPayloadEnvelope {
	return &silapb.SignedBlindedSilaPayloadEnvelope{
		Message: &silapb.BlindedSilaPayloadEnvelope{
			BlockHash:             env.Message.Payload.BlockHash,
			ExecutionRequests:     env.Message.ExecutionRequests,
			BuilderIndex:          env.Message.BuilderIndex,
			BeaconBlockRoot:       env.Message.BeaconBlockRoot,
			Slot:                  primitives.Slot(env.Message.Payload.SlotNumber),
			ParentBlockHash:       env.Message.Payload.ParentHash,
			ParentBeaconBlockRoot: env.Message.ParentBeaconBlockRoot,
		},
		Signature: env.Signature,
	}
}

// encodeBlindedEnvelope SSZ-encodes and snappy-compresses a blinded envelope for storage.
func encodeBlindedEnvelope(env *silapb.SignedBlindedSilaPayloadEnvelope) ([]byte, error) {
	sszBytes, err := env.MarshalSSZ()
	if err != nil {
		return nil, errors.Wrap(err, "could not marshal blinded envelope")
	}
	return snappy.Encode(nil, sszBytes), nil
}

// decodeBlindedEnvelope snappy-decompresses and SSZ-decodes a blinded envelope from storage.
func decodeBlindedEnvelope(enc []byte) (*silapb.SignedBlindedSilaPayloadEnvelope, error) {
	dec, err := snappy.Decode(nil, enc)
	if err != nil {
		return nil, errors.Wrap(err, "could not snappy decode envelope")
	}
	blinded := &silapb.SignedBlindedSilaPayloadEnvelope{}
	if err := blinded.UnmarshalSSZ(dec); err != nil {
		return nil, errors.Wrap(err, "could not unmarshal blinded envelope")
	}
	return blinded, nil
}
