package mainnet

import (
	"testing"

	"github.com/sila-chain/Sila-Consensus-Core/v7/testing/spectest/shared/deneb/merkle_proof"
)

func TestMainnet_Deneb_MerkleProof(t *testing.T) {
	merkle_proof.RunMerkleProofTests(t, "mainnet")
}
