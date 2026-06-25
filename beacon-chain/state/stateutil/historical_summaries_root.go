package stateutil

import (
	fieldparams "github.com/sila-chain/Sila-Consensus-Core/v7/config/fieldparams"
	"github.com/sila-chain/Sila-Consensus-Core/v7/encoding/ssz"
	ethpb "github.com/sila-chain/Sila-Consensus-Core/v7/proto/sila/v1alpha1"
)

func HistoricalSummariesRoot(summaries []*ethpb.HistoricalSummary) ([32]byte, error) {
	return ssz.SliceRoot(summaries, fieldparams.HistoricalRootsLength)
}
