package env

// ContextKey -
type ContextKey int

//
const (
	TraceID ContextKey = iota
	Logger
)

func (c ContextKey) String() string {
	switch c {
	case TraceID:
		return "traceID"
	case Logger:
		return "logger"
	default:
		return ""
	}
}
