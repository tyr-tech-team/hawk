package interceptor

import (
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"
)

// Option -
type Option func(*config)

type config struct {
	TracerProvider trace.TracerProvider
	Propagators    propagation.TextMapPropagator
	Tracer         trace.Tracer
}
