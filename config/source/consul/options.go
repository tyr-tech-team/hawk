package consul

import (
	"context"

	"github.com/tyr-tech-team/hawk/config/source"
)

type address struct{}
type key struct{}
type dialTimeoutKey struct{}
type configType struct{}

// SetAddrs - Addrs is the registry addresses to use
func SetAddrs(addrs string) source.Option {
	return func(o *source.Options) {
		// o.Context = context.Background()
		o.Context = context.WithValue(o.Context, address{}, addrs)
	}

}

// SetKey -
func SetKey(k string) source.Option {
	return func(o *source.Options) {
		o.Context = context.WithValue(o.Context, key{}, k)
	}
}

// SetConfigType -
func SetConfigType(t string) source.Option {
	return func(o *source.Options) {
		o.Context = context.WithValue(o.Context, configType{}, t)
	}
}
