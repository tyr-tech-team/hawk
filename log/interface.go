package log

import (
	"context"
	"strings"

	"github.com/sirupsen/logrus"
	logruslog "github.com/tyr-tech-team/hawk/log/logrus"
	zaplog "github.com/tyr-tech-team/hawk/log/zap"
	"github.com/tyr-tech-team/hawk/status"
	"go.uber.org/zap"
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
func NewZapLogger(level string) *zap.SugaredLogger {
	if strings.ToLower(level) != "prd" {
		return zaplog.NewSuggerLogger(zaplog.DevCore())
	}
	return zaplog.NewSuggerLogger(zaplog.PRDCore())
}

// NewLogrusLogger -
func NewLogrusLogger(level string) *logrus.Logger {
	return logruslog.NewLogrus(level)
}
