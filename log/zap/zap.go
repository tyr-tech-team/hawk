package zap

import (
	"github.com/tyr-tech-team/hawk/status"
	"go.uber.org/zap"
)

type zaplog struct {
	zslog *zap.SugaredLogger
}

//  Errorf -
func (z *zaplog) Errorf(format string, args ...interface{}) {
	z.zslog.Errorf(format, args)
}

func (z *zaplog) With(kv ...interface{}) *zaplog {
	return &zaplog{zslog: z.zslog.With(kv)}
}

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
