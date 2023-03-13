package consul

import (
	"os"
	"time"

	"github.com/hashicorp/consul/api"
	"github.com/tyr-tech-team/hawk/config"
	"github.com/tyr-tech-team/hawk/pkg/traefik"
)

// Config -
type Config struct {
	Address string `json:"url"`
	ACL     string `json:"acl"`
	TTL     time.Duration
}

// ToAgentServiceRegistration -
func ToAgentServiceRegistration(s config.ServiceRegister) *api.AgentServiceRegistration {
	// # 增加 POD Hostname 來當作註冊名稱
	if hostname := os.Getenv("HOSTNAME"); hostname != "" {
		s.ID = s.Name + "-" + hostname
	} else {
		s.ID = s.Name + "-" + s.Address
	}
	asr := &api.AgentServiceRegistration{
		ID:      s.ID,
		Name:    s.Name,
		Address: s.Address,
		Port:    s.Port,
		Tags:    append(s.Tags, s.Name),
		Meta:    map[string]string{"time": time.Now().Format(time.RFC3339)},
		Check: &api.AgentServiceCheck{
			CheckID: s.ID,
			TTL:     (TTL + time.Second).String(),
			// 修改 Timeout 一分半
			Timeout: time.Duration(60 * time.Second).String(),
			// 成功幾次才叫成功
			SuccessBeforePassing: 1,
			// 錯誤幾次就失敗
			FailuresBeforeCritical: 3,
			FailuresBeforeWarning:  3,
			// 當失敗十秒後取消註冊
			DeregisterCriticalServiceAfter: time.Duration(10 * time.Second).String(),
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
