package broker

import (
	"context"
	"sync"
)

// Router -
type Router interface {
	// 設置
	Set(topic string, h Handler)
}

// router -
type router struct {
	// 序列群組
	queueGroup string
	// context
	ctx context.Context
	// 讀寫鎖
	mu sync.RWMutex
	// 代理
	broker Broker
}

// NewRouter -
func NewRouter(ctx context.Context, broker Broker, queueGroup string) Router {
	return &router{
		queueGroup: queueGroup,
		ctx:        ctx,
		mu:         sync.RWMutex{},
		broker:     broker,
	}
}

// Set -
func (r *router) Set(topic string, h Handler) {
	go r.on(topic, h)
}

// on -
func (r *router) on(topic string, h Handler) {
	r.broker.Sub(topic, h, r.queueGroup)

	<-r.ctx.Done()
}
