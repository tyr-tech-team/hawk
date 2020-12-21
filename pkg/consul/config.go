package consul

import (
	"fmt"
	"time"

	"github.com/hashicorp/consul/api"
	"github.com/tyr-tech-team/hawk/pkg/traefik"
	"github.com/tyr-tech-team/hawk/srv"
)

// Config -
type Config struct {
	Address string `json:"url"`
	ACL     string `json:"acl"`
	TTL     time.Duration
}

// ToAgentServiceRegistration -
func ToAgentServiceRegistration(s srv.ServiceRegisterConfig) *api.AgentServiceRegistration {
	s.ID = fmt.Sprintf("%s-%v", s.Name, time.Now().UnixNano())
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

// DefaultConsulConfig -
func DefaultConsulConfig() Config {
	return Config{
		Address: "localhost:8500",
	}
}
