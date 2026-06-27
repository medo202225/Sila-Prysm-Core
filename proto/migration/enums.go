package migration

import (
	"github.com/pkg/errors"
	sila "github.com/sila-chain/Sila-Consensus-Core/v7/proto/sila/v1alpha1"
	silapb "github.com/sila-chain/Sila-Consensus-Core/v7/proto/silaapi/v1"
)

func V1Alpha1ConnectionStateToV1(connState sila.ConnectionState) silapb.ConnectionState {
	alphaString := connState.String()
	v1Value := silapb.ConnectionState_value[alphaString]
	return silapb.ConnectionState(v1Value)
}

func V1Alpha1PeerDirectionToV1(peerDirection sila.PeerDirection) (silapb.PeerDirection, error) {
	alphaString := peerDirection.String()
	if alphaString == sila.PeerDirection_UNKNOWN.String() {
		return 0, errors.New("peer direction unknown")
	}
	v1Value := silapb.PeerDirection_value[alphaString]
	return silapb.PeerDirection(v1Value), nil
}
