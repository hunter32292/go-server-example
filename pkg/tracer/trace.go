package trace

import (
	"io"
	"log"

	opentracing "github.com/opentracing/opentracing-go"

	jaeger "github.com/uber/jaeger-client-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
	jaegerlog "github.com/uber/jaeger-client-go/log"
	"github.com/uber/jaeger-lib/metrics"
)

var (
	traceConfig *jaegercfg.Configuration
	closer      io.Closer
	// globalTrace - Trace object to pass around application
	globalTrace opentracing.Tracer
)

// NewTraceConfig - Setup a trace config to use with the global tracer
func NewTraceConfig(serviceName string) {
	traceConfig = &jaegercfg.Configuration{
		ServiceName: serviceName,
		Sampler: &jaegercfg.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		Reporter: &jaegercfg.ReporterConfig{
			LogSpans: true,
		},
	}
}

// CreateGlobalTracer - Create a global trace object that can be used across services
func CreateGlobalTracer() {
	jLogger := jaegerlog.StdLogger
	jMetricsFactory := metrics.NullFactory
	globalTrace, c, err := traceConfig.NewTracer(
		jaegercfg.Logger(jLogger),
		jaegercfg.Metrics(jMetricsFactory),
	)
	if err != nil {
		log.Fatal(err)
	}
	closer = c
	opentracing.SetGlobalTracer(globalTrace)
}

// GetGlobalTracer - Return the global trace object to use for tracing
func GetGlobalTracer() opentracing.Tracer {
	return opentracing.GlobalTracer()
}

// CloseTracer - Close the tracer on cleanup and exit
func CloseTracer() error {
	return closer.Close()
}
