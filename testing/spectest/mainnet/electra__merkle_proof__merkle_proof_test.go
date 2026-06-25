package mainnet

import (
	"testing"

	"github.com/sila-chain/Sila-Consensus-Core/v7/testing/spectest/shared/electra/merkle_proof"
)

func TestMainnet_Electra_MerkleProof(t *testing.T) {
	merkle_proof.RunMerkleProofTests(t, "mainnet")
}
