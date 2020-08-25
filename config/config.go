package config

import (
	"hawk/config/source"

	"hawk/config/source/file"
)

// Config is an interface abstraction for dynamic configuration
type Config interface {
}

// Load config sources
func Load(source ...source.Source) error {
	return nil
	//return DefaultConfig.Load(source...)
}

// LoadFile is short hand for creating a file source and loading it
func LoadFile(path string) error {
	return Load(file.NewSource(
		file.WithPath(path),
	))
}

