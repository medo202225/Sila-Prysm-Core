package stateutil

import (
	fieldparams "github.com/sila-chain/Sila-Prysm-Core/v7/config/fieldparams"
	"github.com/sila-chain/Sila-Prysm-Core/v7/encoding/ssz"
	ethpb "github.com/sila-chain/Sila-Prysm-Core/v7/proto/prysm/v1alpha1"
)

func PendingDepositsRoot(slice []*ethpb.PendingDeposit) ([32]byte, error) {
	return ssz.SliceRoot(slice, fieldparams.PendingDepositsLimit)
}
