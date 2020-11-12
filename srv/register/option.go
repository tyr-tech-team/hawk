package register

import "github.com/hashicorp/consul/api"

// OptionHandler -
type OptionHandler func(*register)

// WithHost -
func WithHost(host string) OptionHandler {
	return func(r *register) {
		r.host = host
	}
}

// WithPort -
func WithPort(port string) OptionHandler {
	return func(r *register) {
		r.port = port
	}
}

// WithConsul -
func WithConsul(consul *api.Client) OptionHandler {
	return func(r *register) {
		r.consulClient = consul
	}
}

// WithName -
func WithName(name string) OptionHandler {
	return func(r *register) {
		r.name = name
	}
}
