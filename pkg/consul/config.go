package consul

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"time"

	"github.com/hashicorp/consul/api"
	"github.com/tyr-tech-team/hawk/env"
)

// Config -
type Config struct {
	Address string `json:"url"`
	ACL     string `json:"acl"`
	TTL     time.Duration
}

// ServiceRegisterConfig -
type ServiceRegisterConfig struct {
	ID      string   `json:"id"`
	Name    string   `json:"name"`
	Tags    []string `json:"tags"`
	Port    int      `json:"port"`
	Address string   `json:"address"`
}

// ToAgentServiceRegistration -
func (s *ServiceRegisterConfig) ToAgentServiceRegistration() *api.AgentServiceRegistration {
	s.ID = fmt.Sprintf("%s-%s", s.Name, s.md5())
	return &api.AgentServiceRegistration{
		ID:      s.ID,
		Name:    s.Name,
		Address: s.Address,
		Port:    s.Port,
		Tags:    append(s.Tags, s.Name, s.ID),
		Check: &api.AgentServiceCheck{
			CheckID:                        s.ID,
			TTL:                            TTL.String(),
			Timeout:                        time.Minute.String(),
			SuccessBeforePassing:           3,
			FailuresBeforeCritical:         3,
			DeregisterCriticalServiceAfter: env.Fail,
		},
	}
}

func (s *ServiceRegisterConfig) md5() string {
	var h [16]byte
	rand.Read(h[:])
	return hex.EncodeToString(md5.New().Sum(h[:]))
}
