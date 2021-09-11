package client

import (
	"github.com/tyr-tech-team/hawk/broker"
)

type client struct {
	broker broker.Broker
}

func NewTsmmClient(broker broker.Broker, opts ...Options) Client {
	c := &client{
		broker: broker,
	}

	for _, v := range opts {
		v(c)
	}

	return c
}
