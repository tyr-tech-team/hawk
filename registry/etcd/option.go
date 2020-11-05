package etcd

import (
	"context"

	"hawk/registry"
)

type addressKey struct{}

// WithAddress sets the etcd address
func WithAddress(a ...string) registry.Option {
	return func(o *registry.Options) {
		if o.Context == nil {
			o.Context = context.Background()
		}
		o.Context = context.WithValue(o.Context, addressKey{}, a)
	}
}

type authKey struct{}

type authCreds struct {
	Username string
	Password string
}

// Auth allows you to specify username/password
func SetAuth(username, password string) registry.Option {
	return func(o *registry.Options) {
		if o.Context == nil {
			o.Context = context.Background()
		}
		o.Context = context.WithValue(o.Context, authKey{}, &authCreds{Username: username, Password: password})
	}
}

// SetAddrs -
func SetAddrs(addrs ...string) registry.Option {
	return func(o *registry.Options) {
		o.Addrs = addrs
	}
}

