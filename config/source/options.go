package source

import (
	"context"

	"github.com/tyr-tech-team/hawk/config/encoder"

	"github.com/tyr-tech-team/hawk/config/encoder/json"
)

// Options -
type Options struct {
	// Encoder
	Encoder encoder.Encoder

	// for alternative data
	Context context.Context
}

// Option -
type Option func(o *Options)

// NewOptions -
func NewOptions(opts ...Option) Options {
	options := Options{
		Encoder: json.NewEncoder(),
		Context: context.Background(),
	}

	for _, o := range opts {
		o(&options)
	}

	return options
}

// WithEncoder - sets the source encoder
func WithEncoder(e encoder.Encoder) Option {
	return func(o *Options) {
		o.Encoder = e
	}
}
