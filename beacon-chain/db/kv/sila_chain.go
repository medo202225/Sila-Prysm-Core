package kv

import (
	"context"
	"errors"

	"github.com/sila-chain/Sila-Consensus-Core/v7/monitoring/tracing"
	"github.com/sila-chain/Sila-Consensus-Core/v7/monitoring/tracing/trace"
	v2 "github.com/sila-chain/Sila-Consensus-Core/v7/proto/sila/v1alpha1"
	bolt "go.etcd.io/bbolt"
	"google.golang.org/protobuf/proto"
)

// SaveSilaChainData saves the Sila chain data.
func (s *Store) SaveSilaChainData(ctx context.Context, data *v2.SilaChainData) error {
	_, span := trace.StartSpan(ctx, "BeaconDB.SaveSilaChainData")
	defer span.End()

	if data == nil {
		err := errors.New("cannot save nil silaData")
		tracing.AnnotateError(span, err)
		return err
	}

	err := s.db.Update(func(tx *bolt.Tx) error {
		bkt := tx.Bucket(powchainBucket)
		enc, err := proto.Marshal(data)
		if err != nil {
			return err
		}
		return bkt.Put(powchainDataKey, enc)
	})
	tracing.AnnotateError(span, err)
	return err
}

// SilaChainData retrieves the Sila chain data.
func (s *Store) SilaChainData(ctx context.Context) (*v2.SilaChainData, error) {
	_, span := trace.StartSpan(ctx, "BeaconDB.SilaChainData")
	defer span.End()

	var data *v2.SilaChainData
	err := s.db.View(func(tx *bolt.Tx) error {
		bkt := tx.Bucket(powchainBucket)
		enc := bkt.Get(powchainDataKey)
		if len(enc) == 0 {
			return nil
		}
		data = &v2.SilaChainData{}
		return proto.Unmarshal(enc, data)
	})
	return data, err
}
