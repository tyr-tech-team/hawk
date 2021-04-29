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
