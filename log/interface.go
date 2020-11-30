package log

import (
	"context"

	"github.com/tyr-tech-team/hawk/status"
)

// Logger -
type Logger interface {
	Debugf(format string, args ...interface{})
	Infof(format string, args ...interface{})
	Warningf(format string, args ...interface{})
	Errorf(format string, args ...interface{})
	Paincf(format string, args ...interface{})
	With(field ...Field) Logger
	WithStatus(s status.Status) Logger
	WithContext(ctx context.Context) Logger
}

// Field -
type Field struct {
	Key   string
	Value interface{}
}
