package consul

import (
	"context"
	"log"
	"time"

	"github.com/hashicorp/consul/api"
	"github.com/tyr-tech-team/hawk/env"
	"github.com/tyr-tech-team/hawk/status"
)

// -
const (
	TTL = time.Duration(10 * time.Second)
)

type client struct {
	ctx             context.Context
	cancel          context.CancelFunc
	config          Config
	consul          *api.Client
	kv              *api.KV
	sRegistryConfig *api.AgentServiceRegistration
}

func (c *client) Client() *api.Client {
	return c.consul
}

func (c *client) Get(key string) ([]byte, error) {
	pair, _, err := c.kv.Get(key, nil)
	if err != nil {
		return nil, status.NotFound.Err()
	}
	return pair.Value, nil
}

func (c *client) Set(key string, value []byte) error {
	_, err := c.kv.Put(&api.KVPair{
		Key:   key,
		Value: value,
	}, nil)
	if err != nil {
		return status.CreatedFailed.Err()
	}
	return nil
}

// SetRegisterConfig -
func (c *client) SetRegisterConfig(config *ServiceRegisterConfig) {
	c.sRegistryConfig = config.ToAgentServiceRegistration()
}

// Register -
func (c *client) Register() error {
	defer c.healthCheck()
	return c.consul.Agent().ServiceRegister(c.sRegistryConfig)
}

// HealthCheck -
func (c *client) healthCheck() {
	go func(c *client) {
		log.Println("health checking")
		defer c.Deregister()
		for {
			select {
			case <-c.ctx.Done():
				log.Printf("error: %v", c.ctx.Err())
				return
			case <-time.After(c.config.TTL):
				if err := c.healthcheck(); err != nil {
					log.Println(err)
					return
				}
			}
		}
	}(c)
}

func (c *client) healthcheck() error {
	if err := c.consul.Agent().UpdateTTL(c.sRegistryConfig.ID, "", env.Pass); err != nil {
		return status.HealthCheckFailed.Err()
	}
	return nil
}

// Deregister -
func (c *client) Deregister() error {
	return c.consul.Agent().ServiceDeregister(c.sRegistryConfig.ID)
}

func (c *client) Close() {
	c.cancel()
}
