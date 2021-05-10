package broker

import (
	"context"
)

// Options -
type Options struct {
	Context context.Context
}

// Option -
type Option func(o *Options)

// NewOptions -
func NewOptions(opts ...Option) Options {
	options := Options{
		Context: context.Background(),
	}

	for _, o := range opts {
		o(&options)
	}

	return options
}

// SubscribeOption -
type SubscribeOption func(*SubscribeOptions)

// SubscribeOptions -
type SubscribeOptions struct {
	Queue   string
	Context context.Context
}

// Queue sets the name of the queue to share messages on
func Queue(name string) SubscribeOption {
	return func(o *SubscribeOptions) {
		o.Queue = name
	}
}
