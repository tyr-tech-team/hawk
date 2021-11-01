package zap

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

//TODO: lookup go-mirco logger
type zapSugaryLogger func(msg string, kv ...interface{})

func (z *zapSugaryLogger) Log(msg string, kv ...interface{}) {

}

var zaplogger *zap.Logger

// NewLogger -
func NewLogger(core zapcore.Core, zapLevel zapcore.Level) *zap.Logger {
	// AddCallerSkip(2) - 顯示調用位置(行數)
	// AddStacktrace - 跳過哪些level的log不秀出來(Warn 只會秀出 Warn & Error log)
	zaplogger = zap.New(core, zap.AddCallerSkip(2), zap.AddStacktrace(zapLevel))
	return zaplogger
}

// NewSuggerLogger -
func NewSuggerLogger(core zapcore.Core, zapLevel zapcore.Level) *zap.SugaredLogger {
	if zaplogger != nil {
		return zaplogger.Sugar()
	}

	zaplogger = zap.New(core, zap.AddCallerSkip(2), zap.AddStacktrace(zapLevel))
	return zaplogger.Sugar()
}

// NewCore -
func NewCore(zapLevel zapcore.Level) zapcore.Core {
	// 編譯方式
	encoder := zapcore.NewConsoleEncoder(encoderConfig())

	// 輸出樣式
	output := zapcore.Lock(os.Stderr)

	core := zapcore.NewTee(
		zapcore.NewCore(
			encoder,
			output,
			zapLevel,
		),
	)

	return core
}

// EncoderConfig  -
func encoderConfig() zapcore.EncoderConfig {
	base := zapcore.EncoderConfig{
		MessageKey:     "msg",
		TimeKey:        "ts",
		LevelKey:       "level",
		NameKey:        "name",
		CallerKey:      "caller",
		FunctionKey:    "fun",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeTime:     zapcore.RFC3339TimeEncoder,
		EncodeDuration: zapcore.MillisDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	// 顏色大小寫區隔
	base.EncodeLevel = zapcore.LowercaseColorLevelEncoder

	return base
}
