package registry

import (
	"context"
	"time"
)

// Option -
type Option func(*Options)

// Options -
type Options struct {
	Addrs   []string
	Timeout time.Duration
	Context context.Context
}

// Addrs is the registry addresses to use
func Addrs(addrs ...string) Option {
	return func(o *Options) {
		o.Addrs = addrs
	}
}

func Timeout(t time.Duration) Option {
	return func(o *Options) {
		o.Timeout = t
	}
}

