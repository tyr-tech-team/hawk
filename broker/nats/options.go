package nats

import (
	"context"

	"github.com/tyr-tech-team/hawk/broker"
)

type url struct{}

type user struct{}

type password struct{}

// SetAddrs - Addrs is the registry addresses to use
func SetURL(in string) broker.Option {
	return func(o *broker.Options) {
		o.Context = context.WithValue(o.Context, url{}, in)
	}
}

// SetUser -
func SetUser(in string) broker.Option {
	return func(o *broker.Options) {
		o.Context = context.WithValue(o.Context, user{}, in)
	}
}

// SetPassword -
func SetPassword(in string) broker.Option {
	return func(o *broker.Options) {
		o.Context = context.WithValue(o.Context, password{}, in)
	}
}
