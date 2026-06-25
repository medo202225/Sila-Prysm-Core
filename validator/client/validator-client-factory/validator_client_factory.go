package validator_client_factory

import (
	"github.com/sila-chain/Sila-Prysm-Core/v7/config/features"
	beaconApi "github.com/sila-chain/Sila-Prysm-Core/v7/validator/client/beacon-api"
	grpcApi "github.com/sila-chain/Sila-Prysm-Core/v7/validator/client/grpc-api"
	"github.com/sila-chain/Sila-Prysm-Core/v7/validator/client/iface"
	validatorHelpers "github.com/sila-chain/Sila-Prysm-Core/v7/validator/helpers"
)

func NewValidatorClient(
	validatorConn validatorHelpers.NodeConnection,
	opt ...beaconApi.ValidatorClientOpt,
) iface.ValidatorClient {
	if features.Get().EnableBeaconRESTApi {
		return beaconApi.NewBeaconApiValidatorClient(validatorConn.GetRestConnectionProvider(), opt...)
	}
	return grpcApi.NewGrpcValidatorClient(validatorConn)
}
