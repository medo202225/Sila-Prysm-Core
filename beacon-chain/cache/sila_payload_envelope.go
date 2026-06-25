package cache

import (
	"sync"

	consensusblocks "github.com/sila-chain/Sila-Consensus-Core/v7/consensus-types/blocks"
	silapb "github.com/sila-chain/Sila-Consensus-Core/v7/proto/sila/v1alpha1"
)

// SilaPayloadContents holds the producer's envelope with precomputed
// data column sidecars; raw blobs/proofs are derived from the columns at read
// time so the publish hot path skips the KZG cell extension.
type SilaPayloadContents struct {
	Envelope    *silapb.SilaPayloadEnvelope
	DataColumns []consensusblocks.RODataColumn
}

// SilaPayloadEnvelopeCache holds the most recent SilaPayloadContents
// produced by the proposer. Single-entry; Set replaces.
type SilaPayloadEnvelopeCache struct {
	mu       sync.RWMutex
	contents *SilaPayloadContents
}

func NewSilaPayloadEnvelopeCache() *SilaPayloadEnvelopeCache {
	return &SilaPayloadEnvelopeCache{}
}

// Set replaces the cached contents. No-op on nil receiver/contents/envelope so
// readers can treat Envelope and Envelope.Payload as non-nil on a hit.
func (c *SilaPayloadEnvelopeCache) Set(contents *SilaPayloadContents) {
	if c == nil || contents == nil || contents.Envelope == nil || contents.Envelope.Payload == nil {
		return
	}
	c.mu.Lock()
	defer c.mu.Unlock()
	c.contents = contents
}

// Contents returns a snapshot of the cached bundle. The struct is freshly
// allocated; inner slices alias the cache (safe — Set re-assigns whole).
func (c *SilaPayloadEnvelopeCache) Contents() (*SilaPayloadContents, bool) {
	if c == nil {
		return nil, false
	}
	c.mu.RLock()
	defer c.mu.RUnlock()
	if c.contents == nil {
		return nil, false
	}
	snapshot := *c.contents
	return &snapshot, true
}
