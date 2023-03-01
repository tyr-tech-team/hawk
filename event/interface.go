// Package event provides event  
package event

// BrokerCast -
type BrokerCast interface {
	Publish(topic string, message interface{}) error
	Subscribe(topic string, handler Handler) error
}
