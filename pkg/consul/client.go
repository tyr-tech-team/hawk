package consul

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/hashicorp/consul/api"
	"github.com/tyr-tech-team/hawk/status"
)

// -
const (
	TTL         = time.Duration(10 * time.Second)
	AvalibeTime = time.Duration(5 * time.Second)
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
		fmt.Println("in healthCheck")
		for {
			select {
			case <-c.ctx.Done():
				fmt.Println("cancel healthcheck")
				return
			// 每五秒鐘更新一次狀態
			default:
				if err := c.updateHealth(); err != nil {
					return
				}
				time.Sleep(AvalibeTime)
			}
		}
	}(c)
}

func (c *client) updateHealth() error {
	fmt.Println(time.Now().Format(time.RFC3339), " - healthcheck")
	if err := c.consul.Agent().PassTTL(c.sRegistryConfig.ID, time.Now().Format(time.RFC3339)); err != nil {
		return status.HealthCheckFailed.Err()
	}
	return nil
}

// Deregister -
func (c *client) Deregister() error {
	fmt.Println("deregister")
	err := c.consul.Agent().ServiceDeregister(c.sRegistryConfig.ID)
	if err != nil {
		log.Fatalf("deregister failed ,%v  ", err)
	}
	return err
}

func (c *client) Close() {
	c.cancel()
}
