package tracing

import (
	"github.com/opentracing/opentracing-go"
	"github.com/pkg/errors"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
	"io"
	"net/http"
	"os"
)

func Init(service string, env map[string]string) (opentracing.Tracer, io.Closer, error) {
	for key, value := range env {
		os.Setenv(key, value)
	}
	cfg, err := config.FromEnv()
	if err != nil {
		return nil, nil, errors.Wrap(err, "failed to read config from env vars")
	}
	tracer, closer, err := cfg.NewTracer(config.Logger(jaeger.StdLogger))
	if err != nil {
		return nil, nil, errors.Wrap(err, "cannot init Jaeger")
	}
	return tracer, closer, nil
}

// func StartSpanFromRequest(tracer opentracing.Tracer, r *http.Request) opentracing.Span {
// 	spanCtx, _ := Extract(tracer, r)
// 	return tracer.StartSpan("ping-receive", ext.RPCServerOption(spanCtx))
// }

// Inject injects the outbound HTTP request with the given span's context to ensure
// correct propagation of span context throughout the trace.
func Inject(span opentracing.Span, request *http.Request) error {
	return span.Tracer().Inject(
		span.Context(),
		opentracing.HTTPHeaders,
		opentracing.HTTPHeadersCarrier(request.Header))
}

// Extract extracts the inbound HTTP request to obtain the parent span's context to ensure
// correct propagation of span context throughout the trace.
func Extract(tracer opentracing.Tracer, r *http.Request) (opentracing.SpanContext, error) {
	return tracer.Extract(
		opentracing.HTTPHeaders,
		opentracing.HTTPHeadersCarrier(r.Header))
}
