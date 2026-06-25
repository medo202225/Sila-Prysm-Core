package cache

import (
	"bytes"
	"testing"

	"github.com/sila-chain/Sila-Consensus-Core/v7/consensus-types/primitives"
	silapb "github.com/sila-chain/Sila-Consensus-Core/v7/proto/sila/v1alpha1"
	"github.com/sila-chain/Sila-Consensus-Core/v7/testing/require"
)

func TestHighestSilaPayloadBidCache_GetSetIfHigher(t *testing.T) {
	c := NewHighestSilaPayloadBidCache()
	bid := testSignedSilaPayloadBid(10, [32]byte{0x01}, [32]byte{0x02}, 100)

	inserted := c.SetIfHigher(bid)
	require.Equal(t, true, inserted)

	got, ok := c.Get(10, [32]byte{0x01}, [32]byte{0x02})
	require.Equal(t, true, ok)
	require.DeepEqual(t, bid, got)
}

func TestHighestSilaPayloadBidCache_SetIfHigher_ReplacesOnlyOnHigherValue(t *testing.T) {
	c := NewHighestSilaPayloadBidCache()
	low := testSignedSilaPayloadBid(10, [32]byte{0x01}, [32]byte{0x02}, 100)
	same := testSignedSilaPayloadBid(10, [32]byte{0x01}, [32]byte{0x02}, 100)
	high := testSignedSilaPayloadBid(10, [32]byte{0x01}, [32]byte{0x02}, 101)

	require.Equal(t, true, c.SetIfHigher(low))
	require.Equal(t, false, c.SetIfHigher(same))

	got, ok := c.Get(10, [32]byte{0x01}, [32]byte{0x02})
	require.Equal(t, true, ok)
	require.DeepEqual(t, low, got)

	require.Equal(t, true, c.SetIfHigher(high))
	got, ok = c.Get(10, [32]byte{0x01}, [32]byte{0x02})
	require.Equal(t, true, ok)
	require.DeepEqual(t, high, got)
}

func TestHighestSilaPayloadBidCache_SetIfHigher_KeepsDistinctTuples(t *testing.T) {
	c := NewHighestSilaPayloadBidCache()
	first := testSignedSilaPayloadBid(10, [32]byte{0x01}, [32]byte{0x02}, 100)
	second := testSignedSilaPayloadBid(10, [32]byte{0x03}, [32]byte{0x02}, 50)
	third := testSignedSilaPayloadBid(10, [32]byte{0x01}, [32]byte{0x04}, 75)

	require.Equal(t, true, c.SetIfHigher(first))
	require.Equal(t, true, c.SetIfHigher(second))
	require.Equal(t, true, c.SetIfHigher(third))

	got, ok := c.Get(10, [32]byte{0x01}, [32]byte{0x02})
	require.Equal(t, true, ok)
	require.DeepEqual(t, first, got)

	got, ok = c.Get(10, [32]byte{0x03}, [32]byte{0x02})
	require.Equal(t, true, ok)
	require.DeepEqual(t, second, got)

	got, ok = c.Get(10, [32]byte{0x01}, [32]byte{0x04})
	require.Equal(t, true, ok)
	require.DeepEqual(t, third, got)
}

func TestHighestSilaPayloadBidCache_PruneBefore(t *testing.T) {
	c := NewHighestSilaPayloadBidCache()
	oldBid := testSignedSilaPayloadBid(9, [32]byte{0x01}, [32]byte{0x02}, 100)
	currentBid := testSignedSilaPayloadBid(10, [32]byte{0x03}, [32]byte{0x04}, 101)

	require.Equal(t, true, c.SetIfHigher(oldBid))
	require.Equal(t, true, c.SetIfHigher(currentBid))

	c.PruneBefore(10)

	_, ok := c.Get(9, [32]byte{0x01}, [32]byte{0x02})
	require.Equal(t, false, ok)

	got, ok := c.Get(10, [32]byte{0x03}, [32]byte{0x04})
	require.Equal(t, true, ok)
	require.DeepEqual(t, currentBid, got)
}

func testSignedSilaPayloadBid(
	slot primitives.Slot,
	parentHash [32]byte,
	parentRoot [32]byte,
	value uint64,
) *silapb.SignedSilaPayloadBid {
	return &silapb.SignedSilaPayloadBid{
		Message: &silapb.SilaPayloadBid{
			Slot:                  slot,
			ParentBlockHash:       bytes.Clone(parentHash[:]),
			ParentBlockRoot:       bytes.Clone(parentRoot[:]),
			BlockHash:             bytes.Repeat([]byte{0x03}, 32),
			PrevRandao:            bytes.Repeat([]byte{0x04}, 32),
			FeeRecipient:          bytes.Repeat([]byte{0x05}, 20),
			GasLimit:              30_000_000,
			BuilderIndex:          1,
			Value:                 primitives.Gwei(value),
			ExecutionPayment:      10,
			SilaRequestsRoot: make([]byte, 32),
		},
		Signature: bytes.Repeat([]byte{0x06}, 96),
	}
}
