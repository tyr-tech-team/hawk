package consul

import (
	"context"
	"log"

	"github.com/hashicorp/consul/api"
	"github.com/tyr-tech-team/hawk/srv"
)

// Client -
type Client interface {
	Client() *api.Client
	Get(key string) ([]byte, error)
	Set(key string, value []byte) error
	SetRegisterConfig(r srv.ServiceRegisterConfig)
	Register() error
	Deregister() error
	Close()
}

// NewClient -
func NewClient(ctx context.Context, c Config) Client {
	cc := api.Config{
		Address: c.Address,
		Token:   c.ACL,
	}

	conn, err := api.NewClient(&cc)
	if err != nil {
		log.Panic(err)
	}

	if c.TTL == 0 {
		c.TTL = TTL
	}

	nctx, cancel := context.WithCancel(ctx)

	return &client{
		ctx:    nctx,
		cancel: cancel,
		config: c,
		consul: conn,
		kv:     conn.KV(),
	}
}
