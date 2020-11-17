package config

import "github.com/tyr-tech-team/hawk/config/source"

// Config -
type Config interface {
	Decode(value interface{}) error
	Raw() ([]byte, error)
}

// New -
func New(s source.Source) Config {
	return nil
}
