package log

// Logger -
type Logger interface {
	Log(kv ...interface{}) error
}
