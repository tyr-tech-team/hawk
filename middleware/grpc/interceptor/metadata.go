// Package interceptor provides interceptor ﳑ
package interceptor

import "google.golang.org/grpc/metadata"

type metadataSupplier struct {
	metadata *metadata.MD
}

// Get -
func (s *metadataSupplier) Get(key string) string {
	values := s.metadata.Get(key)
	if len(values) == 0 {
		return ""
	}
	return values[0]
}

// Set -
func (s *metadataSupplier) Set(key string, value string) {
	s.metadata.Set(key, value)
}

// Keys -
func (s *metadataSupplier) Keys() []string {
	out := make([]string, 0, len(*s.metadata))
	for key := range *s.metadata {
		out = append(out, key)
	}
	return out
}
