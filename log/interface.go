package log

import (
	"context"

	"github.com/tyr-tech-team/hawk/status"
)

// Logger -
type Logger interface {
	Debugf(format string, args ...interface{})
	Infof(format string, args ...interface{})
	Warnf(format string, args ...interface{})
	Errorf(format string, args ...interface{})
	Paincf(format string, args ...interface{})
	With(kv ...interface{}) Logger
	WithStatus(s status.Status) Logger
	WithError(err error) Logger
	WithContext(ctx context.Context) Logger
}

// NewZapLogger --
func NewZapLogger() {

}

// NewLogrusLogger -
func NewLogrusLogger() {

}
