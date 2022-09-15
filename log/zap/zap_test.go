package zap

import (
	"testing"

	"go.uber.org/zap/zapcore"
)

func Test_NewDevZap(t *testing.T) {
	z := New(zapcore.WarnLevel)
	z.Sugar().Error("this is a test for Error")
	z.Sugar().Info("this is a test for Info")
	z.Sugar().Warn("this is a test for Warn")
}

//func Test_NewPRDZap(t *testing.T) {
//z := NewLogger(PRDCore())
//z.Sugar().Error("this is a test for Error")
//z.Sugar().Info("this is a test for Info")
//z.Sugar().Warn("this is a test for Warn")
//}
