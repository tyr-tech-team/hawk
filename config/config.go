package config

import "hawk/config/source"

// Config -
type Config interface {
	source.Source
}

// NewConfig -
func NewConfig(source source.Source) Config {
	return source
}

