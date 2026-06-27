//go:build minimal

package sila

import "github.com/sila-chain/go-bitfield"

func NewPayloadAttestationAggregationBits() bitfield.Bitvector16 {
	return bitfield.NewBitvector16()
}
