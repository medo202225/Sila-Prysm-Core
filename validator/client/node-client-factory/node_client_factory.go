package node_client_factory

import (
	"github.com/sila-chain/Sila-Prysm-Core/v7/config/features"
	beaconApi "github.com/sila-chain/Sila-Prysm-Core/v7/validator/client/beacon-api"
	grpcApi "github.com/sila-chain/Sila-Prysm-Core/v7/validator/client/grpc-api"
	"github.com/sila-chain/Sila-Prysm-Core/v7/validator/client/iface"
	validatorHelpers "github.com/sila-chain/Sila-Prysm-Core/v7/validator/helpers"
)

func NewNodeClient(validatorConn validatorHelpers.NodeConnection) iface.NodeClient {
	grpcClient := grpcApi.NewNodeClient(validatorConn)
	if features.Get().EnableBeaconRESTApi {
		return beaconApi.NewNodeClientWithFallback(validatorConn.GetRestHandler(), grpcClient)
	}
	return grpcClient
}
