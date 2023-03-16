package event

import (
	"encoding/json"
	"time"
)

// Msg struct  
type Msg struct {
	Topic       string
	ID          string
	JSONPayload []byte
	CreatedAt   time.Time
	//
	needAck bool
}

// Decode -
func (e Msg) Decode(value interface{}) error {
	return json.Unmarshal(e.JSONPayload, value)
}

// Handler  
type Handler func(msg *Msg)
