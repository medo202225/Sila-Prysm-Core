package beacon_chain_client_factory

import (
	"github.com/sila-chain/Sila-Consensus-Core/v7/config/features"
	beaconApi "github.com/sila-chain/Sila-Consensus-Core/v7/validator/client/beacon-api"
	grpcApi "github.com/sila-chain/Sila-Consensus-Core/v7/validator/client/grpc-api"
	"github.com/sila-chain/Sila-Consensus-Core/v7/validator/client/iface"
	nodeClientFactory "github.com/sila-chain/Sila-Consensus-Core/v7/validator/client/node-client-factory"
	validatorHelpers "github.com/sila-chain/Sila-Consensus-Core/v7/validator/helpers"
)

func NewChainClient(validatorConn validatorHelpers.NodeConnection) iface.ChainClient {
	grpcClient := grpcApi.NewGrpcChainClient(validatorConn)
	if features.Get().EnableBeaconRESTApi {
		return beaconApi.NewBeaconApiChainClientWithFallback(validatorConn.GetRestHandler(), grpcClient)
	}
	return grpcClient
}

func NewSilaChainClient(validatorConn validatorHelpers.NodeConnection) iface.SilaChainClient {
	if features.Get().EnableBeaconRESTApi {
		return beaconApi.NewSilaChainClient(validatorConn.GetRestHandler(), nodeClientFactory.NewNodeClient(validatorConn))
	}
	return grpcApi.NewGrpcSilaChainClient(validatorConn)
}
