package stateutil

import (
	fieldparams "github.com/sila-chain/Sila-Prysm-Core/v7/config/fieldparams"
	"github.com/sila-chain/Sila-Prysm-Core/v7/encoding/ssz"
	ethpb "github.com/sila-chain/Sila-Prysm-Core/v7/proto/prysm/v1alpha1"
)

func PendingPartialWithdrawalsRoot(slice []*ethpb.PendingPartialWithdrawal) ([32]byte, error) {
	return ssz.SliceRoot(slice, fieldparams.PendingPartialWithdrawalsLimit)
}
