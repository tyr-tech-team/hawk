package interceptor

//type config struct {
//Propagators    propagation.TextMapPropagator
//TracerProvider trace.TracerProvider
//}

//// Option applies an option value for a config.
//type Option interface {
//Apply(*config)
//}

//// Inject -
//func Inject(ctx context.Context, metadata *metadata.MD, opts ...Option) {
//c := newConfig(opts)
//c.Propagators.Inject(ctx, &metadataSupplier{
//metadata: metadata,
//})
//}

//func newConfig(opts []Option) *config {
//c := &config{
//Propagators:    otel.GetTextMapPropagator(),
//TracerProvider: otel.GetTracerProvider(),
//}
//for _, o := range opts {
//o.Apply(c)
//}
//return c
//}
