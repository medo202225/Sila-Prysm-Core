//go:build !minimal

package sila

import "github.com/sila-chain/go-bitfield"

func NewPayloadAttestationAggregationBits() bitfield.Bitvector512 {
	return bitfield.NewBitvector512()
}
