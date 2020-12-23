package traefik

import (
	"fmt"

	"github.com/tyr-tech-team/hawk/config"
)

// NewTags -
func NewTags(name string, protocol config.Protocol) []string {
	tags := []string{
		"traefik.enable=true",
		fmt.Sprintf("traefik.http.routers.%s.rule=Host(`%s.traefik`)", name, name),
		fmt.Sprintf("traefik.http.routers.%s.service=%s-service", name, name),
		fmt.Sprintf("traefik.http.routers.%s.middlewares=latency-check,do-retry", name),
		fmt.Sprintf("traefik.http.services.%s-service.loadbalancer.passhostheader=true", name),
		fmt.Sprintf("traefik.http.services.%s-service.loadbalancer.server.scheme=%s", name, traefikScheme(protocol)),
	}
	// 中間件標籤
	tags = append(tags, httpMiddlewareTags()...)
	return tags
}

func httpMiddlewareTags() []string {
	return []string{
		// 錯誤閘道
		"traefik.http.middlewares.latency-check.circuitbreaker.expression=NetworkErrorRatio() > 0.50",
		// 重傳 3次
		"traefik.http.middlewares.do-retry.retry.attempts=3",
	}
}

func traefikScheme(p config.Protocol) string {
	scheme := "http"
	switch p {
	case config.GRPC:
		scheme = "h2c"
	}
	return scheme
}
