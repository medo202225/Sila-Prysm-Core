package testing

import (
	grpcutil "github.com/sila-chain/Sila-Consensus-Core/v7/api/grpc"
	"github.com/sila-chain/Sila-Consensus-Core/v7/validator/helpers"
)

// MockNodeConnection creates a minimal NodeConnection for testing.
func MockNodeConnection() helpers.NodeConnection {
	conn, _ := helpers.NewNodeConnection(
		helpers.WithGRPCProvider(&grpcutil.MockGrpcProvider{
			MockHosts: []string{"mock:4000"},
		}),
	)
	return conn
}
