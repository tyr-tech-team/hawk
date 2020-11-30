package zap

import (
	"os"

	"github.com/tyr-tech-team/hawk/env"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type zapSugaryLogger func(msg string, kv ...interface{})

var zaplogger *zap.Logger

// NewLogger -
func NewLogger(mode env.Mode) *zap.Logger {
	zaplogger = zap.New(newCore(mode), zap.AddCallerSkip(2), zap.AddStacktrace(zapcore.ErrorLevel))
	return zaplogger
}

// NewSuggerLogger -
func NewSuggerLogger(mode env.Mode) *zap.SugaredLogger {
	if zaplogger != nil {
		return zaplogger.Sugar()
	}

	zaplogger = zap.New(newCore(mode), zap.AddCallerSkip(2), zap.AddStacktrace(zapcore.ErrorLevel))
	return zaplogger.Sugar()
}

// newCore -
func newCore(mode env.Mode) zapcore.Core {
	// 高優先權
	hightPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.PanicLevel
	})
	// 低優先權
	lowPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl < zapcore.PanicLevel
	})
	// 編譯方式
	encoder := zapcore.NewJSONEncoder(encoderConfig(mode))

	// 輸出
	outputDebuging := zapcore.Lock(os.Stdout)
	outputErrors := zapcore.Lock(os.Stderr)

	core := zapcore.NewTee(
		zapcore.NewCore(
			encoder,
			outputErrors,
			hightPriority,
		),
		zapcore.NewCore(
			encoder,
			outputDebuging,
			lowPriority,
		),
	)

	return core
}

func developmentEncoderConfig() zapcore.EncoderConfig {
	return encoderConfig(env.DEV)
}

// EncoderConfig  -
func encoderConfig(mode env.Mode) zapcore.EncoderConfig {

	base := zapcore.EncoderConfig{
		MessageKey:     "ms",
		TimeKey:        "ts",
		LevelKey:       "level",
		NameKey:        "name",
		CallerKey:      "caller",
		FunctionKey:    "fun",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.RFC3339TimeEncoder,
		EncodeDuration: zapcore.MillisDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	switch mode {
	case env.DEV:
		base.EncodeLevel = zapcore.LowercaseColorLevelEncoder
	}

	return base
}

func config(mode env.Mode, ec zapcore.EncoderConfig) zap.Config {
	base := zap.Config{
		Level:         zap.NewAtomicLevelAt(zapcore.ErrorLevel),
		Development:   false,
		EncoderConfig: ec,
		Encoding:      "json",
		Sampling: &zap.SamplingConfig{
			Initial:    100,
			Thereafter: 100,
		},
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
	}

	switch mode {
	case env.DEV:
		base.Encoding = "console"
		base.Development = true
		base.Sampling = nil
		base.OutputPaths = []string{"stderr"}
	}

	return base
}
