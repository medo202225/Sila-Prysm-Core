package params

import "github.com/sila-chain/Sila-Consensus-Core/v7/consensus-types/primitives"

const BasisPoints = primitives.BP(10000)

// SlotBP returns the duration of a slot expressed in milliseconds, represented as basis points of a slot.
func SlotBP() primitives.BP {
	return primitives.BP(BeaconConfig().SlotDurationMillis())
}
