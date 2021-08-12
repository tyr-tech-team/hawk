package trace

import (
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
)

// Config -
type Config struct {
	// trace exporter
	Exporter tracesdk.SpanExporter

	// service name
	Service string

	// environment
	Environment string
}

// New -
func New(c Config) *tracesdk.TracerProvider {
	tp := tracesdk.NewTracerProvider(
		// Set the sampling rate based on the parent span to 100%
		tracesdk.WithSampler(tracesdk.ParentBased(tracesdk.TraceIDRatioBased(1.0))),

		// set exporter
		tracesdk.WithBatcher(c.Exporter),

		tracesdk.WithSampler(tracesdk.AlwaysSample()),

		tracesdk.WithResource(
			resource.NewWithAttributes(
				semconv.SchemaURL,
				semconv.ServiceNameKey.String(c.Service),
				attribute.String("environment", c.Environment),
			),
		),
	)

	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))

	return tp
}

// TracerProvider -
//func TracerProvider(c Config) {
//tp := tracesdk.NewTracerProvider(
//// Always be sure to batch in production.
//tracesdk.WithBatcher(c.Exporter),

//// Record information about this application in an Resource.
//tracesdk.WithResource(resource.NewWithAttributes(
//semconv.ServiceNameKey.String(c.Service),
//attribute.String(environment, c.Environment),
//)),
//)

//otel.SetTracerProvider(tp)
//otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))
//}
