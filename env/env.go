package env

// ContextKey -
type ContextKey int

//
const (
	RequestID ContextKey = iota
	Logger
)

func (c ContextKey) String() string {
	switch c {
	case RequestID:
		return "requestid"
	case Logger:
		return "logger"
	default:
		return ""
	}
}
