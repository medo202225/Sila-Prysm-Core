package mainnet

import (
	"testing"

	"github.com/sila-chain/Sila-Prysm-Core/v7/testing/spectest/shared/altair/sanity"
)

func TestMainnet_Altair_Sanity_Slots(t *testing.T) {
	sanity.RunSlotProcessingTests(t, "mainnet")
}
