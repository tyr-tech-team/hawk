package logrus

import (
	"os"
	"path"
	"runtime"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
)

var logger *logrus.Logger

//NewLogrus -
func NewLogrus(level string) *logrus.Logger {
	logger = logrus.New()
	logger.SetOutput(os.Stdout)
	logger.SetReportCaller(true)
	logger.SetFormatter(&logrus.JSONFormatter{
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			filename := path.Base(f.File)
			return f.Function + "()", filename + ":" + strconv.Itoa(f.Line)
		},
	})
	switch strings.ToLower(level) {
	case "trace":
		logger.SetLevel(logrus.TraceLevel)
	case "info":
		logger.SetLevel(logrus.InfoLevel)
	case "error":
		logger.SetLevel(logrus.ErrorLevel)
	case "warning":
		logger.SetLevel(logrus.WarnLevel)
	case "fatal":
		logger.SetLevel(logrus.FatalLevel)
	case "panic":
		logger.SetLevel(logrus.PanicLevel)
	default:
		logger.SetLevel(logrus.DebugLevel)
	}

	return logger
}
