package consul

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"time"

	"github.com/hashicorp/consul/api"
)

// Protocol -
type Protocol string

const (
	// GRPC -
	GRPC Protocol = "grpc"
	// HTTP -
	HTTP Protocol = "http"
)

// Config -
type Config struct {
	Address string `json:"url"`
	ACL     string `json:"acl"`
	TTL     time.Duration
}

// ServiceRegisterConfig -
type ServiceRegisterConfig struct {
	ID       string   `json:"id"`
	Name     string   `json:"name"`
	Tags     []string `json:"tags"`
	Port     int      `json:"port"`
	Address  string   `json:"address"`
	Traefik  bool     `json:"traefik"`
	Protocol Protocol `json:"protocol"`
}

// ToAgentServiceRegistration -
func (s *ServiceRegisterConfig) ToAgentServiceRegistration() *api.AgentServiceRegistration {

	asr := &api.AgentServiceRegistration{
		ID:      fmt.Sprintf("%s-%s", s.Name, s.md5()),
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
			FailuresBeforeCritical:         3,
			DeregisterCriticalServiceAfter: TTL.String(),
		},
	}

	if s.Traefik {
		asr.Tags = append(asr.Tags, createTraefikTags(asr.Name, s.Protocol)...)
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

func createTraefikTags(name string, protocal Protocol) []string {
	s := []string{
		"traefik.enable=true",
		fmt.Sprintf("traefik.http.routers.%s.service=%s-service", name, name),
		"traefik.http.middlewares.latency-check.circuitbreaker.expression=NetworkErrorRatio() > 0.50",
		fmt.Sprintf("traefik.http.routers.%s.middlewares=latency-check", name),
		func(protocol Protocol) string {
			scheme := "http"
			switch protocol {
			case GRPC:
				scheme = "h2c"
			}
			return fmt.Sprintf("traefik.http.services.test-service.loadbalancer.server.scheme=%s", scheme)
		}(protocal),
		"traefik.http.services.test-service.loadbalancer.passhostheader=true",
	}
	return s
}
