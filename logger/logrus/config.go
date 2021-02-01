package logrus

import (
	"os"
	"path"
	"runtime"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/tyr-tech-team/hawk/config"
)

//NewLogrus -
func NewLogrus(c config.Log) *logrus.Logger {
	l := logrus.New()
	l.SetOutput(os.Stdout)
	l.SetReportCaller(true)
	l.SetFormatter(&logrus.JSONFormatter{
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			filename := path.Base(f.File)
			return f.Function + "()", filename + ":" + strconv.Itoa(f.Line)
		},
	})

	setLevel(l, c.Level)

	return l
}

// SetLevel -
func setLevel(l *logrus.Logger, level string) {
	switch strings.ToLower(level) {
	case "trace":
		l.SetLevel(logrus.TraceLevel)
	case "info":
		l.SetLevel(logrus.InfoLevel)
	case "error":
		l.SetLevel(logrus.ErrorLevel)
	case "warn":
		l.SetLevel(logrus.WarnLevel)
	default:
		l.SetLevel(logrus.DebugLevel)
	}
}
