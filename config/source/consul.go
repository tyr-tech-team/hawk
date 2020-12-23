package source

import (
	"time"

	"github.com/tyr-tech-team/hawk/config"
	"github.com/tyr-tech-team/hawk/pkg/consul"
)

type consulClient struct {
	client consul.Client
	key    string
}

// NewConsul -
func NewConsul(client consul.Client, key string) config.Source {
	return consulClient{
		client: client,
		key:    key,
	}
}

func (c consulClient) Read() (*config.ChangeSet, error) {
	b, err := c.client.Get(c.key)
	if err != nil {
		return nil, err
	}
	return &config.ChangeSet{
		Data:      b,
		Checksum:  Sum(b),
		Timestamp: time.Now(),
	}, nil
}
