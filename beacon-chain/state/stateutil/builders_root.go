package stateutil

import (
	fieldparams "github.com/sila-chain/Sila-Consensus-Core/v7/config/fieldparams"
	"github.com/sila-chain/Sila-Consensus-Core/v7/encoding/ssz"
	ethpb "github.com/sila-chain/Sila-Consensus-Core/v7/proto/sila/v1alpha1"
)

// BuildersRoot computes the SSZ root of a slice of Builder.
func BuildersRoot(slice []*ethpb.Builder) ([32]byte, error) {
	return ssz.SliceRoot(slice, uint64(fieldparams.BuilderRegistryLimit))
}
