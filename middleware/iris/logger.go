package iris

import (
	"time"

	"github.com/kataras/iris/v12"
	"github.com/sirupsen/logrus"
	"github.com/tyr-tech-team/hawk/trace"
)

type logger struct {
	TraceID      string      `json:"traceid"`
	Method       string      `json:"method"`
	Path         string      `json:"path"`
	Raw          string      `json:"raw"`
	ClientIP     string      `json:"clientip"`
	Size         int64       `json:"size"`
	Start        time.Time   `json:"start"`
	End          time.Time   `json:"end"`
	Latency      string      `json:"latency"`
	HTTPCode     int         `json:"httpcode"`
	ErrorMessage interface{} `json:"error"`
}

// LoggerWithLogrus  -
func LoggerWithLogrus(log *logrus.Entry) iris.Handler {
	return func(ctx iris.Context) {
		l := logger{}

		t := time.Now()

		l.Start = t

		ctx.Next()

		tid := trace.GetTraceID(ctx.Request().Context())
		l.TraceID = tid

		// after request
		l.Method = ctx.Method()
		l.Path = ctx.Path()
		l.ClientIP = ctx.RemoteAddr()
		l.End = time.Now()
		l.Latency = time.Since(t).String()
		l.HTTPCode = ctx.GetStatusCode()
		l.Raw = ctx.Request().URL.RawFragment
		l.Size = ctx.GetContentLength()
		l.ErrorMessage = ctx.Values().Get("error").(error)

		log = log.WithFields(logrus.Fields{
			"traceid":  l.TraceID,
			"start":    l.Start.Format(time.RFC3339),
			"end":      l.End.Format(time.RFC3339),
			"latency":  l.End.Sub(l.Start).Microseconds(),
			"clientip": l.ClientIP,
			"httpcode": l.HTTPCode,
			"error":    l.ErrorMessage,
			"method":   l.Method,
			"path":     l.Path,
			"size":     l.Size,
			"raw":      l.Raw,
		})

		if l.ErrorMessage != "" {
			log.Errorln("failed request")
		} else {
			log.Infoln("success request")
		}
	}
}
