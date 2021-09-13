package client

import (
	"github.com/tyr-tech-team/hawk/broker"
)

type client struct {
	broker broker.Broker
}

// NewTsmmClient - 新增Tsmm客戶端
// Options選項：
func NewTsmmClient(broker broker.Broker, opts ...Options) Client {
	c := &client{
		broker: broker,
	}

	for _, v := range opts {
		v(c)
	}

	return c
}
