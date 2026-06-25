package testing

import (
	"github.com/sila-chain/Sila-Consensus-Core/v7/time/slots"
)

var _ slots.Ticker = (*MockTicker)(nil)
