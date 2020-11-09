package config

import "github.com/tyr-tech-team/hawk/config/source"

// Config -
type Config interface {
	source.Source
}

// NewConfig -
func NewConfig(source source.Source) Config {
	return source
}
