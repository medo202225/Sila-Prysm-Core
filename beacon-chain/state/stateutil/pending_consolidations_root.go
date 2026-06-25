package stateutil

import (
	fieldparams "github.com/sila-chain/Sila-Consensus-Core/v7/config/fieldparams"
	"github.com/sila-chain/Sila-Consensus-Core/v7/encoding/ssz"
	ethpb "github.com/sila-chain/Sila-Consensus-Core/v7/proto/sila/v1alpha1"
)

func PendingConsolidationsRoot(slice []*ethpb.PendingConsolidation) ([32]byte, error) {
	return ssz.SliceRoot(slice, fieldparams.PendingConsolidationsLimit)
}
