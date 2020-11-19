package logger

// Logger -
type Logger interface {
	Log(kv ...interface{}) error
}
