// Package nats provides nats  
package nats

import (
	"github.com/nats-io/nats.go"
)

type client struct {
	natCli *nats.Conn
}

// JetStreamClient interface  
type JetStreamClient interface {
}

// New function  
func New() (JetStreamClient, error) {
	// create the nats client
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		return nil, err
	}

	return &client{
		natCli: nc,
	}, nil
}
