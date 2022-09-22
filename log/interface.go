// Package log provides log ﳑ
package log

import (
	"strings"

	"github.com/sirupsen/logrus"
	logruslog "github.com/tyr-tech-team/hawk/log/logrus"
	zaplog "github.com/tyr-tech-team/hawk/log/zap"
	"github.com/tyr-tech-team/hawk/status"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
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
	// WithContext(ctx context.Context) Logger
}

// NewZapLogger -
func NewZapLogger(level string) *zap.Logger {
	var zapLevel zapcore.Level

	switch strings.ToLower(level) {
	case "info":
		zapLevel = zapcore.InfoLevel
	case "error":
		zapLevel = zapcore.ErrorLevel
	case "warning":
		zapLevel = zapcore.WarnLevel
	case "fatal":
		zapLevel = zapcore.FatalLevel
	case "panic":
		zapLevel = zapcore.PanicLevel
	case "debug":
		zapLevel = zapcore.DebugLevel
	default:
		zapLevel = zapcore.DebugLevel
	}

	return zaplog.New(zapLevel)
}

// NewLogrusLogger -
func NewLogrusLogger(level string) *logrus.Logger {
	return logruslog.NewLogrus(level)
}
