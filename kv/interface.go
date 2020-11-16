package kv

// KV -
type KV interface {
	Get(key string) []byte
	Set(key string, value interface{}) error
	Keys() []string
}
