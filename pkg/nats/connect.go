// Package nats provides nats  
package nats

import (
	"time"

	"github.com/nats-io/nats.go"
	"github.com/tyr-tech-team/arceus/topic/v1"
	"github.com/tyr-tech-team/hawk/config"
	"github.com/tyr-tech-team/hawk/eventbus"
	"github.com/tyr-tech-team/hawk/status"
	"go.uber.org/zap"
)

type jetStreamClient struct {
	nc      *nats.Conn
	jc      nats.JetStreamContext
	stream  string
	sublist map[topic.Topic]*nats.Subscription
	log     *zap.Logger
}

// NewJetStream function  
func NewJetStream(config config.NatsJetstream, log *zap.Logger) (eventbus.EventBus, error) {
	if config.Stream == "" {
		return nil, status.ConnectFailed.WithDetail("empyt stream").Err()
	}
	// create the nats client
	nc, err := nats.Connect(
		config.Hosts,
		nats.ClosedHandler(func(nc *nats.Conn) {}),
		nats.RetryOnFailedConnect(true),
		nats.ReconnectWait(5*time.Second),
	)

	if err != nil {
		log.Sugar().With("error", err.Error()).Error("create nats failed")
		return nil, status.ConnectFailed.WithDetail(err.Error()).Err()
	}

	jc, err := nc.JetStream()
	if err != nil {
		log.Sugar().With("error", err.Error()).Error("create nats jetstream failed")
		return nil, status.ConnectFailed.WithDetail(err.Error()).Err()
	}

	return &jetStreamClient{
		nc:      nc,
		jc:      jc,
		stream:  config.Stream,
		sublist: make(map[topic.Topic]*nats.Subscription),
		log:     log,
	}, nil
}

func (cli *jetStreamClient) Close() {
	for _, v := range cli.sublist {
		v.Drain()
	}

	cli.nc.Close()
}
