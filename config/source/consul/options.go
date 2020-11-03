package consul

import (
	"context"
	"hawk/config/source"

	"github.com/hashicorp/consul/api"
)

type addressKey struct{}
type key struct{}
type dialTimeoutKey struct{}
type configType struct{}
type client struct{}

// SetClient - Client is the
func SetClient(c *api.Client) source.Option {
	return func(o *source.Options) {
		o.Context = context.WithValue(o.Context, client{}, c)
	}
}

// SetAddrs - Addrs is the registry addresses to use
func SetAddrs(addrs ...string) source.Option {
	return func(o *source.Options) {
		// o.Context = context.Background()
		o.Context = context.WithValue(o.Context, addressKey{}, addrs)
	}

}

// SetName -
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
