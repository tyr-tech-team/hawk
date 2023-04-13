// Package eventbus provides eventbus  
package eventbus

import (
	"github.com/nats-io/nats.go"
	"github.com/tyr-tech-team/arceus/topic/v1"
)

// FuncHandler  
type FuncHandler func(msg *nats.Msg)

// EventBus -
type EventBus interface {
	// Publish(topic topic.Topic, data []byte) error
	// Subscribe(topic topic.Topic, queue string, handler FuncHandler) error
	Subscriber
	Publisher
}

// Publisher interface  
type Publisher interface {
	Publish(topic topic.Topic, data []byte) error
	Close()
}

// Subscriber interface  
type Subscriber interface {
	Subscribe(topic topic.Topic, queue string, handler FuncHandler) error
	Close()
}
