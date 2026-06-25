package sync

import (
	"testing"

	"github.com/sila-chain/go-bitfield"
	mock "github.com/sila-chain/Sila-Consensus-Core/v7/beacon-chain/blockchain/testing"
	"github.com/sila-chain/Sila-Consensus-Core/v7/beacon-chain/operations/attestations"
	lruwrpr "github.com/sila-chain/Sila-Consensus-Core/v7/cache/lru"
	fieldparams "github.com/sila-chain/Sila-Consensus-Core/v7/config/fieldparams"
	ethpb "github.com/sila-chain/Sila-Consensus-Core/v7/proto/sila/v1alpha1"
	"github.com/sila-chain/Sila-Consensus-Core/v7/testing/assert"
	"github.com/sila-chain/Sila-Consensus-Core/v7/testing/require"
	"github.com/sila-chain/Sila-Consensus-Core/v7/testing/util"
)

func TestBeaconAggregateProofSubscriber_CanSaveAggregatedAttestation(t *testing.T) {
	r := &Service{
		cfg: &config{
			attPool:             attestations.NewPool(),
			attestationNotifier: (&mock.ChainService{}).OperationNotifier(),
		},
		seenUnAggregatedAttestationCache: lruwrpr.New(10),
	}

	a := &ethpb.SignedAggregateAttestationAndProof{
		Message: &ethpb.AggregateAttestationAndProof{
			Aggregate: util.HydrateAttestation(&ethpb.Attestation{
				AggregationBits: bitfield.Bitlist{0x07},
			}),
			AggregatorIndex: 100,
		},
		Signature: make([]byte, fieldparams.BLSSignatureLength),
	}
	require.NoError(t, r.beaconAggregateProofSubscriber(t.Context(), a))
	assert.DeepSSZEqual(t, []ethpb.Att{a.Message.Aggregate}, r.cfg.attPool.AggregatedAttestations(), "Did not save aggregated attestation")
}

func TestBeaconAggregateProofSubscriber_CanSaveUnaggregatedAttestation(t *testing.T) {
	r := &Service{
		cfg: &config{
			attPool:             attestations.NewPool(),
			attestationNotifier: (&mock.ChainService{}).OperationNotifier(),
		},
		seenUnAggregatedAttestationCache: lruwrpr.New(10),
	}

	a := &ethpb.SignedAggregateAttestationAndProof{
		Message: &ethpb.AggregateAttestationAndProof{
			Aggregate: util.HydrateAttestation(&ethpb.Attestation{
				AggregationBits: bitfield.Bitlist{0x03},
				Signature:       make([]byte, fieldparams.BLSSignatureLength),
			}),
			AggregatorIndex: 100,
		},
	}
	require.NoError(t, r.beaconAggregateProofSubscriber(t.Context(), a))

	atts := r.cfg.attPool.UnaggregatedAttestations()
	assert.DeepEqual(t, []ethpb.Att{a.Message.Aggregate}, atts, "Did not save unaggregated attestation")
}
