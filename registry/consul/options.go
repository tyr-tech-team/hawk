package consul

import (
	"context"
	"hawk/config/source"
)

type addressKey struct{}
type name struct{}
type dialTimeoutKey struct{}
type configType struct{}

// SetAddrs - Addrs is the registry addresses to use
func SetAddrs(addrs ...string) source.Option {
	return func(o *source.Options) {
		o.Context = context.WithValue(o.Context, addressKey{}, addrs)
	}
}

// SetName -
func SetName(n string) source.Option {
	return func(o *source.Options) {
		o.Context = context.WithValue(o.Context, name{}, n)
	}
}

// SetConfigType -
func SetConfigType(t string) source.Option {
	return func(o *source.Options) {
		o.Context = context.WithValue(o.Context, configType{}, t)
	}
}
