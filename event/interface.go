// Package event provides event  
package event

import "context"

// Command -
type Command interface {
	Publish(ctx context.Context, subject string, message interface{}) error
	Subscribe(topic string, handler func(*Msg)) error
}
