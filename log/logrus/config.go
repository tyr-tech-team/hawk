package logrus

import (
	"os"
	"path"
	"runtime"
	"strconv"

	"github.com/sirupsen/logrus"
)

// NewLogger -
func NewLogger() *logrus.Logger {
	l := logrus.New()
	l.SetOutput(os.Stdout)
	l.SetReportCaller(true)
	l.SetFormatter(&logrus.JSONFormatter{
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			filename := path.Base(f.File)
			//return fmt.Sprintf("%s()", f.Function), fmt.Sprintf("%s:%d", filename, f.Line)
			return f.Function + "()", filename + ":" + strconv.Itoa(f.Line)
		},
	})

	return l
}
