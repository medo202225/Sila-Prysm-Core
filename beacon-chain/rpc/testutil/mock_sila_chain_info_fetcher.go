package testutil

import (
	"math/big"
)

// MockSilaChainInfoFetcher is a fake implementation of the powchain.ChainInfoFetcher
type MockSilaChainInfoFetcher struct {
	CurrEndpoint string
	CurrError    error
}

func (*MockSilaChainInfoFetcher) GenesisSilaChainInfo() (uint64, *big.Int) {
	return uint64(0), &big.Int{}
}

func (*MockSilaChainInfoFetcher) SilaClientConnected() bool {
	return true
}

func (m *MockSilaChainInfoFetcher) SilaClientEndpoint() string {
	return m.CurrEndpoint
}

func (m *MockSilaChainInfoFetcher) SilaClientConnectionErr() error {
	return m.CurrError
}
