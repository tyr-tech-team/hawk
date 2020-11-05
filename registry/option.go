package registry

import (
	"context"
	"crypto/tls"
	"time"
)

// Option -
type Option func(*Options)

// Options -
type Options struct {
	Addrs     []string
	Timeout   time.Duration
	Context   context.Context
	Secure    bool
	TLSConfig *tls.Config
}

// Addrs is the registry addresses to use
func Addrs(addrs ...string) Option {
	return func(o *Options) {
		o.Addrs = addrs
	}
}

// Timeout -
func Timeout(t time.Duration) Option {
	return func(o *Options) {
		o.Timeout = t
	}
}

// Secure -
func Secure(b bool) Option {
	return func(o *Options) {
		o.Secure = b
	}
}

// TLSConfig - Specify TLS Config
func TLSConfig(t *tls.Config) Option {
	return func(o *Options) {
		o.TLSConfig = t
	}
}
