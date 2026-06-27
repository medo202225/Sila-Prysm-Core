package attestation_test

import (
	"testing"

	fieldparams "github.com/sila-chain/Sila-Consensus-Core/v7/config/fieldparams"
	"github.com/sila-chain/Sila-Consensus-Core/v7/config/params"
	"github.com/sila-chain/Sila-Consensus-Core/v7/consensus-types/primitives"
	sila "github.com/sila-chain/Sila-Consensus-Core/v7/proto/sila/v1alpha1"
	"github.com/sila-chain/Sila-Consensus-Core/v7/proto/sila/v1alpha1/attestation"
	"github.com/sila-chain/Sila-Consensus-Core/v7/testing/assert"
	"github.com/sila-chain/Sila-Consensus-Core/v7/testing/require"
	"github.com/sila-chain/go-bitfield"
)

func TestAttestingIndices(t *testing.T) {
	type args struct {
		att        sila.Att
		committees [][]primitives.ValidatorIndex
	}
	tests := []struct {
		name string
		args args
		want []uint64
		err  string
	}{
		{
			name: "Full committee attested",
			args: args{
				att:        &sila.Attestation{AggregationBits: bitfield.Bitlist{0b1111}},
				committees: [][]primitives.ValidatorIndex{{0, 1, 2}},
			},
			want: []uint64{0, 1, 2},
		},
		{
			name: "Partial committee attested",
			args: args{
				att:        &sila.Attestation{AggregationBits: bitfield.Bitlist{0b1101}},
				committees: [][]primitives.ValidatorIndex{{0, 1, 2}},
			},
			want: []uint64{0, 2},
		},
		{
			name: "Invalid bit length",
			args: args{
				att:        &sila.Attestation{AggregationBits: bitfield.Bitlist{0b11111}},
				committees: [][]primitives.ValidatorIndex{{0, 1, 2}},
			},
			err: "bitfield length 4 is not equal to committee length 3",
		},
		{
			name: "Electra - Full committee attested",
			args: args{
				att:        &sila.AttestationElectra{AggregationBits: bitfield.Bitlist{0b11111}},
				committees: [][]primitives.ValidatorIndex{{0, 1}, {2, 3}},
			},
			want: []uint64{0, 1, 2, 3},
		},
		{
			name: "Electra - Partial committee attested",
			args: args{
				att:        &sila.AttestationElectra{AggregationBits: bitfield.Bitlist{0b10110}},
				committees: [][]primitives.ValidatorIndex{{0, 1}, {2, 3}},
			},
			want: []uint64{1, 2},
		},
		{
			name: "Electra - Invalid bit length",
			args: args{
				att:        &sila.AttestationElectra{AggregationBits: bitfield.Bitlist{0b111111}},
				committees: [][]primitives.ValidatorIndex{{0, 1}, {2, 3}},
			},
			err: "bitfield length 5 is not equal to committee length 4",
		},
		{
			name: "Electra - No duplicates",
			args: args{
				att:        &sila.AttestationElectra{AggregationBits: bitfield.Bitlist{0b11111}},
				committees: [][]primitives.ValidatorIndex{{0, 1}, {0, 1}},
			},
			want: []uint64{0, 1},
		},
		{
			name: "Electra - No attester in committee",
			args: args{
				att:        &sila.AttestationElectra{AggregationBits: bitfield.Bitlist{0b11100}},
				committees: [][]primitives.ValidatorIndex{{0, 1}, {0, 1}},
			},
			err: "no attesting indices found for committee index 0",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := attestation.AttestingIndices(tt.args.att, tt.args.committees...)
			if tt.err == "" {
				require.NoError(t, err)
				assert.DeepEqual(t, tt.want, got)
			} else {
				require.ErrorContains(t, tt.err, err)
			}
		})
	}
}

func TestIsValidAttestationIndices(t *testing.T) {
	tests := []struct {
		name      string
		att       sila.IndexedAtt
		wantedErr string
	}{
		{
			name: "Indices should not be nil",
			att: &sila.IndexedAttestation{
				Data: &sila.AttestationData{
					Target: &sila.Checkpoint{},
					Source: &sila.Checkpoint{},
				},
				Signature: make([]byte, fieldparams.BLSSignatureLength),
			},
			wantedErr: "expected non-empty attesting indices",
		},
		{
			name: "Indices should be non-empty",
			att: &sila.IndexedAttestation{
				AttestingIndices: []uint64{},
				Data: &sila.AttestationData{
					Target: &sila.Checkpoint{},
					Source: &sila.Checkpoint{},
				},
				Signature: make([]byte, fieldparams.BLSSignatureLength),
			},
			wantedErr: "expected non-empty",
		},
		{
			name: "Greater than max validators per committee",
			att: &sila.IndexedAttestation{
				AttestingIndices: make([]uint64, params.BeaconConfig().MaxValidatorsPerCommittee+1),
				Data: &sila.AttestationData{
					Target: &sila.Checkpoint{},
					Source: &sila.Checkpoint{},
				},
				Signature: make([]byte, fieldparams.BLSSignatureLength),
			},
			wantedErr: "indices count exceeds",
		},
		{
			name: "Needs to be sorted",
			att: &sila.IndexedAttestation{
				AttestingIndices: []uint64{3, 2, 1},
				Data: &sila.AttestationData{
					Target: &sila.Checkpoint{},
					Source: &sila.Checkpoint{},
				},
				Signature: make([]byte, fieldparams.BLSSignatureLength),
			},
			wantedErr: "not uniquely sorted",
		},
		{
			name: "Valid indices",
			att: &sila.IndexedAttestation{
				AttestingIndices: []uint64{1, 2, 3},
				Data: &sila.AttestationData{
					Target: &sila.Checkpoint{},
					Source: &sila.Checkpoint{},
				},
				Signature: make([]byte, fieldparams.BLSSignatureLength),
			},
		},
		{
			name: "Valid indices with length of 2",
			att: &sila.IndexedAttestation{
				AttestingIndices: []uint64{1, 2},
				Data: &sila.AttestationData{
					Target: &sila.Checkpoint{},
					Source: &sila.Checkpoint{},
				},
				Signature: make([]byte, fieldparams.BLSSignatureLength),
			},
		},
		{
			name: "Valid indices with length of 1",
			att: &sila.IndexedAttestation{
				AttestingIndices: []uint64{1},
				Data: &sila.AttestationData{
					Target: &sila.Checkpoint{},
					Source: &sila.Checkpoint{},
				},
				Signature: make([]byte, fieldparams.BLSSignatureLength),
			},
		},
		{
			name: "Electra - Greater than max validators per slot",
			att: &sila.IndexedAttestationElectra{
				AttestingIndices: make([]uint64, params.BeaconConfig().MaxValidatorsPerCommittee*params.BeaconConfig().MaxCommitteesPerSlot+1),
				Data: &sila.AttestationData{
					Target: &sila.Checkpoint{},
					Source: &sila.Checkpoint{},
				},
				Signature: make([]byte, fieldparams.BLSSignatureLength),
			},
			wantedErr: "indices count exceeds",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := attestation.IsValidAttestationIndices(t.Context(), tt.att, params.BeaconConfig().MaxValidatorsPerCommittee, params.BeaconConfig().MaxCommitteesPerSlot)
			if tt.wantedErr != "" {
				assert.ErrorContains(t, tt.wantedErr, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func BenchmarkAttestingIndices_PartialCommittee(b *testing.B) {
	bf := bitfield.Bitlist{0b11111111, 0b11111111, 0b10000111, 0b11111111, 0b100}
	committee := []primitives.ValidatorIndex{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33}

	for b.Loop() {
		_, err := attestation.AttestingIndices(&sila.Attestation{AggregationBits: bf}, committee)
		require.NoError(b, err)
	}
}

func BenchmarkIsValidAttestationIndices(b *testing.B) {
	indices := make([]uint64, params.BeaconConfig().MaxValidatorsPerCommittee)
	for i := range indices {
		indices[i] = uint64(i)
	}
	att := &sila.IndexedAttestation{
		AttestingIndices: indices,
		Data: &sila.AttestationData{
			Target: &sila.Checkpoint{},
			Source: &sila.Checkpoint{},
		},
		Signature: make([]byte, fieldparams.BLSSignatureLength),
	}

	for b.Loop() {
		if err := attestation.IsValidAttestationIndices(b.Context(), att, params.BeaconConfig().MaxValidatorsPerCommittee, params.BeaconConfig().MaxCommitteesPerSlot); err != nil {
			require.NoError(b, err)
		}
	}
}

func TestAttDataIsEqual(t *testing.T) {
	type test struct {
		name     string
		attData1 *sila.AttestationData
		attData2 *sila.AttestationData
		equal    bool
	}
	tests := []test{
		{
			name: "same",
			attData1: &sila.AttestationData{
				Slot:            5,
				CommitteeIndex:  2,
				BeaconBlockRoot: []byte("great block"),
				Source: &sila.Checkpoint{
					Epoch: 4,
					Root:  []byte("good source"),
				},
				Target: &sila.Checkpoint{
					Epoch: 10,
					Root:  []byte("good target"),
				},
			},
			attData2: &sila.AttestationData{
				Slot:            5,
				CommitteeIndex:  2,
				BeaconBlockRoot: []byte("great block"),
				Source: &sila.Checkpoint{
					Epoch: 4,
					Root:  []byte("good source"),
				},
				Target: &sila.Checkpoint{
					Epoch: 10,
					Root:  []byte("good target"),
				},
			},
			equal: true,
		},
		{
			name: "diff slot",
			attData1: &sila.AttestationData{
				Slot:            5,
				CommitteeIndex:  2,
				BeaconBlockRoot: []byte("great block"),
				Source: &sila.Checkpoint{
					Epoch: 4,
					Root:  []byte("good source"),
				},
				Target: &sila.Checkpoint{
					Epoch: 10,
					Root:  []byte("good target"),
				},
			},
			attData2: &sila.AttestationData{
				Slot:            4,
				CommitteeIndex:  2,
				BeaconBlockRoot: []byte("great block"),
				Source: &sila.Checkpoint{
					Epoch: 4,
					Root:  []byte("good source"),
				},
				Target: &sila.Checkpoint{
					Epoch: 10,
					Root:  []byte("good target"),
				},
			},
		},
		{
			name: "diff block",
			attData1: &sila.AttestationData{
				Slot:            5,
				CommitteeIndex:  2,
				BeaconBlockRoot: []byte("good block"),
				Source: &sila.Checkpoint{
					Epoch: 4,
					Root:  []byte("good source"),
				},
				Target: &sila.Checkpoint{
					Epoch: 10,
					Root:  []byte("good target"),
				},
			},
			attData2: &sila.AttestationData{
				Slot:            5,
				CommitteeIndex:  2,
				BeaconBlockRoot: []byte("great block"),
				Source: &sila.Checkpoint{
					Epoch: 4,
					Root:  []byte("good source"),
				},
				Target: &sila.Checkpoint{
					Epoch: 10,
					Root:  []byte("good target"),
				},
			},
		},
		{
			name: "diff source root",
			attData1: &sila.AttestationData{
				Slot:            5,
				CommitteeIndex:  2,
				BeaconBlockRoot: []byte("great block"),
				Source: &sila.Checkpoint{
					Epoch: 4,
					Root:  []byte("good source"),
				},
				Target: &sila.Checkpoint{
					Epoch: 10,
					Root:  []byte("good target"),
				},
			},
			attData2: &sila.AttestationData{
				Slot:            5,
				CommitteeIndex:  2,
				BeaconBlockRoot: []byte("great block"),
				Source: &sila.Checkpoint{
					Epoch: 4,
					Root:  []byte("bad source"),
				},
				Target: &sila.Checkpoint{
					Epoch: 10,
					Root:  []byte("good target"),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.equal, attestation.AttDataIsEqual(tt.attData1, tt.attData2))
		})
	}
}

func TestCheckPtIsEqual(t *testing.T) {
	type test struct {
		name     string
		checkPt1 *sila.Checkpoint
		checkPt2 *sila.Checkpoint
		equal    bool
	}
	tests := []test{
		{
			name: "same",
			checkPt1: &sila.Checkpoint{
				Epoch: 4,
				Root:  []byte("good source"),
			},
			checkPt2: &sila.Checkpoint{
				Epoch: 4,
				Root:  []byte("good source"),
			},
			equal: true,
		},
		{
			name: "diff epoch",
			checkPt1: &sila.Checkpoint{
				Epoch: 4,
				Root:  []byte("good source"),
			},
			checkPt2: &sila.Checkpoint{
				Epoch: 5,
				Root:  []byte("good source"),
			},
			equal: false,
		},
		{
			name: "diff root",
			checkPt1: &sila.Checkpoint{
				Epoch: 4,
				Root:  []byte("good source"),
			},
			checkPt2: &sila.Checkpoint{
				Epoch: 4,
				Root:  []byte("bad source"),
			},
			equal: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.equal, attestation.CheckPointIsEqual(tt.checkPt1, tt.checkPt2))
		})
	}
}

func BenchmarkAttDataIsEqual(b *testing.B) {
	attData1 := &sila.AttestationData{
		Slot:            5,
		CommitteeIndex:  2,
		BeaconBlockRoot: []byte("great block"),
		Source: &sila.Checkpoint{
			Epoch: 4,
			Root:  []byte("good source"),
		},
		Target: &sila.Checkpoint{
			Epoch: 10,
			Root:  []byte("good target"),
		},
	}
	attData2 := &sila.AttestationData{
		Slot:            5,
		CommitteeIndex:  2,
		BeaconBlockRoot: []byte("great block"),
		Source: &sila.Checkpoint{
			Epoch: 4,
			Root:  []byte("good source"),
		},
		Target: &sila.Checkpoint{
			Epoch: 10,
			Root:  []byte("good target"),
		},
	}

	b.Run("fast", func(b *testing.B) {
		b.ReportAllocs()
		for b.Loop() {
			assert.Equal(b, true, attestation.AttDataIsEqual(attData1, attData2))
		}
	})

	b.Run("proto.Equal", func(b *testing.B) {
		b.ReportAllocs()
		for b.Loop() {
			assert.Equal(b, true, attestation.AttDataIsEqual(attData1, attData2))
		}
	})
}
