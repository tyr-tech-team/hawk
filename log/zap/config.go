// Package zap provides zap ﳑ
package zap

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// type zapSugaryLogger func(msg string, kv ...interface{})

// func (z *zapSugaryLogger) Log(msg string, kv ...interface{}) {

// }

var zaplogger *zap.Logger

// New function
func New(level zapcore.Level) *zap.Logger {
    if zaplogger != nil {
        return zaplogger
    }
	return setup(level)
}

// NewSugger function
func NewSugger(level zapcore.Level) *zap.SugaredLogger {
    if zaplogger != nil {
        return zaplogger.Sugar()
    }
	return setup(level).Sugar()
}

// Setup -
func setup(level zapcore.Level) *zap.Logger {
	cfg := zap.Config{
		Level:            zap.NewAtomicLevelAt(level),
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
		EncoderConfig:    encoderConfig(),
		Sampling:         nil,
		Encoding:         "json",
	}

	logger, err := cfg.Build()
	if err != nil {
		panic(err)
	}

	logger = logger.WithOptions(
		zap.AddCallerSkip(1),
		zap.AddStacktrace(zapcore.PanicLevel),
	)

	return logger
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
	// 小寫的 error , panic , debug , info 標誌
	base.EncodeLevel = zapcore.LowercaseLevelEncoder

	return base
}



