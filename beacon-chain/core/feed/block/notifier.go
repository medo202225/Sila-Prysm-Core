package block

import "github.com/sila-chain/Sila-Prysm-Core/v7/async/event"

// Notifier interface defines the methods of the service that provides block updates to consumers.
type Notifier interface {
	BlockFeed() *event.Feed
}
