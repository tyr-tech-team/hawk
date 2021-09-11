package server

import (
	"github.com/tyr-tech-team/hawk/broker"
)

const (
	DEFAULT_QUEUE_NAME = "tsmm"
)

type server struct {
	broker     broker.Broker
	subscriber []broker.Subscriber
	queueName  string
}

func NewTsmmServer(broker broker.Broker, opts ...Options) Server {
	s := &server{
		queueName: DEFAULT_QUEUE_NAME,
		broker:    broker,
	}

	for _, v := range opts {
		v(s)
	}

	return s
}

func (s *server) SetQueueName(queueName string) {
	if queueName == "" {
		s.queueName = DEFAULT_QUEUE_NAME
		return
	}
	s.queueName = queueName
}
