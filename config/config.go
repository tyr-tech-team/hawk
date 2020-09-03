package config

import "hawk/config/source"

// Config -
type Config interface {
	source.Source
}

// newConfig -
func NewConfig(source source.Source) Config {
	return source
}

