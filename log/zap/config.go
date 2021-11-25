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
func NewLogger(core zapcore.Core) *zap.Logger {
	zaplogger = zap.New(core, zap.AddCallerSkip(2), zap.AddStacktrace(zapcore.PanicLevel))
	return zaplogger
}

// NewSuggerLogger -
func NewSuggerLogger(core zapcore.Core) *zap.SugaredLogger {
	if zaplogger != nil {
		return zaplogger.Sugar()
	}

	zaplogger = zap.New(core, zap.AddCallerSkip(2), zap.AddStacktrace(zapcore.ErrorLevel))
	return zaplogger.Sugar()
}

// DevCore -
func DevCore() zapcore.Core {
	// 高優先權
	hightPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.PanicLevel
	})
	// 低優先權
	lowPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl < zapcore.PanicLevel
	})
	// 編譯方式
	encoder := zapcore.NewConsoleEncoder(encoderConfig("dev"))

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

// PRDCore -
func PRDCore() zapcore.Core {
	// 高優先權
	hightPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.PanicLevel
	})
	// 低優先權
	lowPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl < zapcore.PanicLevel
	})
	// 編譯方式
	encoder := zapcore.NewJSONEncoder(encoderConfig("prd"))

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

// EncoderConfig  -
func encoderConfig(mode string) zapcore.EncoderConfig {

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

	switch mode {
	case "dev":
		base.EncodeLevel = zapcore.LowercaseColorLevelEncoder
	default:
		base.EncodeLevel = zapcore.LowercaseLevelEncoder
	}

	return base
}

func config(mode string, ec zapcore.EncoderConfig) zap.Config {
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
	case "dev":
		base.Encoding = "console"
		base.Development = true
		base.Sampling = nil
		base.OutputPaths = []string{"stderr"}
	}

	return base
}
