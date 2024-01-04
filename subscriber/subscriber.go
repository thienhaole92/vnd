package subscriber

import (
	"context"

	"github.com/thienhaole92/vnd/event"
)

type ConsumeFunc func(ctx context.Context, message event.EventString) error

// Subscriber is the interface to manage ONE topic
type Subscriber interface {
	// Start the subscriber
	Start()
	// Close all underlying connections and resources
	Close() error
	// GroupID that is used
	GroupID() string
	// Topic that it is subscribed
	Topic() string
}
