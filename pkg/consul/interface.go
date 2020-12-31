package consul

import (
	"github.com/hashicorp/consul/api"
	"github.com/tyr-tech-team/hawk/config"
)

// Client -
type Client interface {
	Client() *api.Client
	Get(key string) ([]byte, error)
	Set(key string, value []byte) error
	SetRegisterConfig(r config.ServiceRegister)
	Register() error
	Deregister() error
	Close()
}
