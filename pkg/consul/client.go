package consul

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/hashicorp/consul/api"
	"github.com/tyr-tech-team/hawk/config"
	"github.com/tyr-tech-team/hawk/status"
)

// -
const (
	TTL           = time.Duration(30 * time.Second)
	AvailableTime = time.Duration(5 * time.Second)
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
func (c *client) SetRegisterConfig(cfg config.ServiceRegister) {
	c.sRegistryConfig = ToAgentServiceRegistration(cfg)
}

// Register -
func (c *client) Register() error {
	defer c.healthCheck()
	err := c.consul.Agent().ServiceRegister(c.sRegistryConfig)
	if err != nil {
		return status.ConnectFailed.WithDetail(err.Error()).Err()
	}

	return nil
}

// HealthCheck -
func (c *client) healthCheck() {
	go func(c *client) {
		defer c.Deregister()
		for {
			select {
			case <-c.ctx.Done():
				fmt.Println("cancel healthcheck")
				return
			// 每五秒鐘更新一次狀態
			default:
				if err := c.updateHealth(); err != nil {
					panic(err)
				}
				time.Sleep(AvailableTime)
			}
		}
	}(c)
}

func (c *client) updateHealth() error {
	if err := c.consul.Agent().PassTTL(c.sRegistryConfig.ID, ""); err != nil {
		return status.HealthCheckFailed.WithDetail(err.Error()).Err()
	}
	return nil
}

// Deregister -
func (c *client) Deregister() error {
	err := c.consul.Agent().ServiceDeregister(c.sRegistryConfig.ID)
	if err != nil {
		log.Fatalf("deregister failed ,%v  ", err)
	}
	return nil
}

func (c *client) Close() {
	c.cancel()
}
