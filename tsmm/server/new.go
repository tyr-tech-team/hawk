package server

import (
	"github.com/tyr-tech-team/hawk/broker"
)

const (
	// DEFAULT_QUEUE_NAME -
	DEFAULT_QUEUE_NAME = "tsmm"
)

type server struct {
	broker     broker.Broker
	subscriber []broker.Subscriber
	queueName  string
}

// NewTsmmServer - TsmmServer端
// Options:
//     設定Queue group name
//     server.SetQueueName(queueName string)
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

// SetQueueName - 設定Queue group name
func (s *server) SetQueueName(queueName string) {
	if queueName == "" {
		s.queueName = DEFAULT_QUEUE_NAME
		return
	}
	s.queueName = queueName
}
