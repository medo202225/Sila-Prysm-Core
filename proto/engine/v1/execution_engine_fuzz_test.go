package enginev1_test

import (
	"fmt"
	"testing"

	enginev1 "github.com/sila-chain/Sila-Consensus-Core/v7/proto/engine/v1"
	"github.com/sila-chain/Sila-Consensus-Core/v7/testing/require"
	fuzz "github.com/google/gofuzz"
)

func TestCopySilaPayload_Fuzz(t *testing.T) {
	fuzzCopies(t, &enginev1.SilaPayloadDeneb{})
	fuzzCopies(t, &enginev1.SilaPayloadCapella{})
	fuzzCopies(t, &enginev1.SilaPayload{})
}

func TestCopySilaPayloadHeader_Fuzz(t *testing.T) {
	fuzzCopies(t, &enginev1.SilaPayloadHeaderDeneb{})
	fuzzCopies(t, &enginev1.SilaPayloadHeaderCapella{})
	fuzzCopies(t, &enginev1.SilaPayloadHeader{})
}

func fuzzCopies[T any, C enginev1.Copier[T]](t *testing.T, obj C) {
	fuzzer := fuzz.NewWithSeed(0)
	amount := 1000
	t.Run(fmt.Sprintf("%T", obj), func(t *testing.T) {
		for range amount {
			fuzzer.Fuzz(obj) // Populate thing with random values
			got := obj.Copy()
			require.DeepEqual(t, obj, got)
			// check shallow copy working
			fuzzer.Fuzz(got)
			require.DeepNotEqual(t, obj, got)
			// TODO: think of deeper not equal fuzzing
		}
	})
}
