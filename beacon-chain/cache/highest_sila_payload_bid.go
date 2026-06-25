package cache

import (
	"sync"

	"github.com/sila-chain/Sila-Consensus-Core/v7/consensus-types/primitives"
	silapb "github.com/sila-chain/Sila-Consensus-Core/v7/proto/sila/v1alpha1"
)

type silaPayloadBidKey struct {
	slot       primitives.Slot
	parentHash [32]byte
	parentRoot [32]byte
}

// HighestSilaPayloadBidCache stores the highest bid for each
// (slot, parent_block_hash, parent_block_root) tuple.
type HighestSilaPayloadBidCache struct {
	bids map[silaPayloadBidKey]*silapb.SignedSilaPayloadBid
	lock sync.RWMutex
}

// NewHighestSilaPayloadBidCache initializes a highest-bid cache.
func NewHighestSilaPayloadBidCache() *HighestSilaPayloadBidCache {
	return &HighestSilaPayloadBidCache{
		bids: make(map[silaPayloadBidKey]*silapb.SignedSilaPayloadBid),
	}
}

// Get returns the highest cached bid for the given tuple.
func (c *HighestSilaPayloadBidCache) Get(
	slot primitives.Slot,
	parentHash [32]byte,
	parentRoot [32]byte,
) (*silapb.SignedSilaPayloadBid, bool) {
	c.lock.RLock()
	defer c.lock.RUnlock()

	bid, ok := c.bids[silaPayloadBidKey{
		slot:       slot,
		parentHash: parentHash,
		parentRoot: parentRoot,
	}]
	return bid, ok
}

// SetIfHigher inserts the bid if absent, or replaces the cached bid only if
// the incoming value is strictly greater.
func (c *HighestSilaPayloadBidCache) SetIfHigher(bid *silapb.SignedSilaPayloadBid) bool {
	c.lock.Lock()
	defer c.lock.Unlock()

	key := silaPayloadBidKey{
		slot:       bid.Message.Slot,
		parentHash: [32]byte(bid.Message.ParentBlockHash),
		parentRoot: [32]byte(bid.Message.ParentBlockRoot),
	}
	cached, ok := c.bids[key]
	if !ok || bid.Message.Value > cached.Message.Value {
		c.bids[key] = bid
		return true
	}
	return false
}

// PruneBefore removes all cached bids for slots before the provided slot.
func (c *HighestSilaPayloadBidCache) PruneBefore(slot primitives.Slot) {
	c.lock.Lock()
	defer c.lock.Unlock()

	for key := range c.bids {
		if key.slot < slot {
			delete(c.bids, key)
		}
	}
}
