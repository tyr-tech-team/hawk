package register

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"

	"github.com/hashicorp/consul/api"
	"github.com/tyr-tech-team/hawk/status"
)

type register struct {
	consulClient *api.Client
	name         string
	host         string
	port         string
	id           string
}

// New -
func New(opts ...OptionHandler) Register {
	r := new(register)
	for _, v := range opts {
		v(r)
	}
	return r
}

func (r *register) Register() error {
	if r.consulClient != nil {
		return r.consulRegister()
	}
	return status.NoError.Err()
}

func (r *register) Deregister() error {
	if r.consulClient != nil {
		return r.consulDeregister()
	}
	return status.NoError.Err()
}

func (r *register) consulRegister() error {
	service := &api.AgentServiceRegistration{
		Name:    r.name,
		Address: r.host,
		// TODO: 設置配置檔
	}
	if err := r.consulClient.Agent().ServiceRegister(service); err != nil {
		return status.ConnectFailed.Err()
	}
	return status.NoError.Err()
}

func (r *register) consulDeregister() error {
	if err := r.consulClient.Agent().ServiceDeregister(r.id); err != nil {
		return status.ConnectFailed.Err()
	}

	return status.NoError.Err()
}

func (r *register) setName() string {
	var h [16]byte
	rand.Read(h[:])
	r.id = fmt.Sprintf("%s-%s", r.name, hex.EncodeToString(h[:]))
	return r.id
}
