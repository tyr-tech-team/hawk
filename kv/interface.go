// Package kv provides kv ﳑ
package kv

// KV -
type KV interface {
	Get(key string) ([]byte, error)
	Set(key string, value []byte) error
}
