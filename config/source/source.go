package source

import (
	"time"
)

// Source is the source from which config is loaded
type Source interface {
	Read() (*ChangeSet, error)
}

// ChangeSet represents a set of changes from a source
type ChangeSet struct {
	Data      []byte
	Checksum  string
	Format    string
	Source    string
	Timestamp time.Time
}
