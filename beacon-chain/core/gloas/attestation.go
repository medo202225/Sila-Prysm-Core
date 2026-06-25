package gloas

import (
	"fmt"

	"github.com/sila-chain/Sila-Consensus-Core/v7/beacon-chain/state"
	"github.com/sila-chain/Sila-Consensus-Core/v7/consensus-types/primitives"
	"github.com/sila-chain/Sila-Consensus-Core/v7/runtime/version"
	"github.com/pkg/errors"
)

// MatchingPayload returns true if the attestation's committee index matches the expected payload index.
//
// For pre-Gloas forks, this always returns true.
//
// Spec v1.7.0-alpha (pseudocode):
//
//	# [New in Gloas:SIP7732]
//	if is_attestation_same_slot(state, data):
//	    assert data.index == 0
//	    payload_matches = True
//	else:
//	    slot_index = data.slot % SLOTS_PER_HISTORICAL_ROOT
//	    payload_index = state.sila_payload_availability[slot_index]
//	    payload_matches = data.index == payload_index
func MatchingPayload(
	beaconState state.ReadOnlyBeaconState,
	beaconBlockRoot [32]byte,
	slot primitives.Slot,
	committeeIndex uint64,
) (bool, error) {
	if beaconState.Version() < version.Gloas {
		return true, nil
	}

	sameSlot, err := beaconState.IsAttestationSameSlot(beaconBlockRoot, slot)
	if err != nil {
		return false, errors.Wrap(err, "failed to get same slot attestation status")
	}
	if sameSlot {
		if committeeIndex != 0 {
			return false, fmt.Errorf("committee index %d for same slot attestation must be 0", committeeIndex)
		}
		return true, nil
	}

	silaPayloadAvail, err := beaconState.SilaPayloadAvailability(slot)
	if err != nil {
		return false, errors.Wrap(err, "failed to get sila payload availability status")
	}
	return silaPayloadAvail == committeeIndex, nil
}
