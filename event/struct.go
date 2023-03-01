package event

import "time"

// Event is the struct that represents an event.
type Event struct {
	Topic       string
	ID          string
	JSONPayload string
	CreatedAt   time.Time
}

// Handler -
type Handler interface {
	EvnetRun() error
}
