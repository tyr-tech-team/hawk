package consul

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"time"

	"github.com/hashicorp/consul/api"
	"github.com/tyr-tech-team/hawk/env"
	"github.com/tyr-tech-team/hawk/pkg/traefik"
)

// Config -
type Config struct {
	Address string `json:"url"`
	ACL     string `json:"acl"`
	TTL     time.Duration
}

// ServiceRegisterConfig -
type ServiceRegisterConfig struct {
	ID       string       `json:"id"`
	Name     string       `json:"name"`
	Tags     []string     `json:"tags"`
	Port     int          `json:"port"`
	Address  string       `json:"address"`
	Traefik  bool         `json:"traefik"`
	Protocol env.Protocol `json:"protocol"`
}

// ToAgentServiceRegistration -
func (s *ServiceRegisterConfig) ToAgentServiceRegistration() *api.AgentServiceRegistration {
	s.ID = fmt.Sprintf("%s-%s", s.Name, s.md5())
	asr := &api.AgentServiceRegistration{
		ID:      s.ID,
		Name:    s.Name,
		Address: s.Address,
		Port:    s.Port,
		Tags:    append(s.Tags, s.Name),
		Check: &api.AgentServiceCheck{
			CheckID: s.ID,
			TTL:     (TTL + time.Second).String(),
			Timeout: time.Minute.String(),
			// 成功幾次才叫成功
			SuccessBeforePassing: 1,
			// 錯誤幾次就失敗
			FailuresBeforeCritical:         2,
			DeregisterCriticalServiceAfter: time.Duration(3 * time.Second).String(),
		},
	}

	if s.Traefik {
		asr.Tags = append(asr.Tags, traefik.NewTags(asr.Name, s.Protocol)...)
	}

	return asr
}

func (s *ServiceRegisterConfig) md5() string {
	var h [8]byte
	rand.Read(h[:])
	return hex.EncodeToString(md5.New().Sum(h[:]))
}

// DefaultConsulConfig -
func DefaultConsulConfig() Config {
	return Config{
		Address: "localhost:8500",
	}
}
