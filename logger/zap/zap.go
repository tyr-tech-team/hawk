package zap

import (
	"context"

	"github.com/tyr-tech-team/hawk/status"
	"go.uber.org/zap"
)

type zaplog struct {
	zslog *zap.SugaredLogger
}

// Debug -
func (z *zaplog) Debug(args ...interface{}) {
	z.zslog.Debug(args)
}

// Debugf -
func (z *zaplog) Debugf(format string, args ...interface{}) {
	z.zslog.Debugf(format, args)
}

func (z *zaplog) Info(args ...interface{}) {
	z.zslog.Info(args)
}

// Info -
func (z *zaplog) Infof(format string, args ...interface{}) {
	z.zslog.Infof(format, args)
}

// Warn -
func (z *zaplog) Warn(args ...interface{}) {
	z.zslog.Warn(args)
}

// Warnf -
func (z *zaplog) Warnf(format string, args ...interface{}) {
	z.zslog.Warnf(format, args)
}

// Error -
func (z *zaplog) Error(args ...interface{}) {
	z.zslog.Error(args)
}

//  Errorf -
func (z *zaplog) Errorf(format string, args ...interface{}) {
	z.zslog.Errorf(format, args)
}

// Panic -
func (z *zaplog) Panic(args ...interface{}) {
	z.zslog.Panic(args)
}

// Panicf -
func (z *zaplog) Paincf(format string, args ...interface{}) {
	z.zslog.Panicf(format, args)
}

// With -
func (z *zaplog) With(kv ...interface{}) *zaplog {
	return &zaplog{zslog: z.zslog.With(kv)}
}

// WithStatus -
func (z *zaplog) WithStatus(s status.Status) *zaplog {
	return &zaplog{
		zslog: z.zslog.With(
			"statusCode",
			s.Code(),
			"message",
			s.Message(),
		),
	}
}

// WithError -
func (z *zaplog) WithError(err error) *zaplog {
	return &zaplog{
		zslog: z.zslog.With("error", err.Error()),
	}
}

func (z *zaplog) WithContext(ctx context.Context) *zaplog {

	return z
}
