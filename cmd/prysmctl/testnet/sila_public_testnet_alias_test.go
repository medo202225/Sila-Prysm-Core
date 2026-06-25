package testnet

import (
	"testing"

	"github.com/sila-chain/Sila-Prysm-Core/v7/config/params"
)

func TestSetGlobalParamsSilaPublicTestnetAlias(t *testing.T) {
	previousChainConfigFile := generateGenesisStateFlags.ChainConfigFile
	previousConfig := params.BeaconConfig().Copy()

	t.Cleanup(func() {
		generateGenesisStateFlags.ChainConfigFile = previousChainConfigFile
		if err := params.SetActive(previousConfig); err != nil {
			t.Fatalf("could not restore previous config: %v", err)
		}
	})

	generateGenesisStateFlags.ChainConfigFile = params.SilaPublicTestnetName
	if err := setGlobalParams(); err != nil {
		t.Fatalf("setGlobalParams returned error: %v", err)
	}

	cfg := params.BeaconConfig()
	if cfg.ConfigName != params.SilaPublicTestnetName {
		t.Fatalf("unexpected active config name: got %q want %q", cfg.ConfigName, params.SilaPublicTestnetName)
	}
	if cfg.DepositChainID != 20263001 {
		t.Fatalf("unexpected deposit chain id: got %d want %d", cfg.DepositChainID, uint64(20263001))
	}
	if cfg.DepositNetworkID != 20263001 {
		t.Fatalf("unexpected deposit network id: got %d want %d", cfg.DepositNetworkID, uint64(20263001))
	}
}
