package accounts

import (
	"github.com/sila-chain/Sila-Consensus-Core/v7/validator/keymanager"
)

var (
	errKeymanagerNotSupported = "keymanager kind not supported: %s"
	// ErrCouldNotInitializeKeymanager informs about failed keymanager initialization
	ErrCouldNotInitializeKeymanager = "could not initialize keymanager"
)

// DeleteConfig specifies parameters for the accounts delete command.
type DeleteConfig struct {
	Keymanager       keymanager.IKeymanager
	DeletePublicKeys [][]byte
}
