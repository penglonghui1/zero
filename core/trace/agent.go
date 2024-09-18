package trace

import (
	"fmt"
	"runtime"
	"sync"

	"github.com/pengcainiao/zero/core/env"
	"github.com/pengcainiao/zero/core/lang"
	"github.com/pengcainiao/zero/core/logx"
	"github.com/pengcainiao/zero/core/sysx"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.10.0"
)

const (
	kindJaeger = "jaeger"
	kindZipkin = "zipkin"
)

var (
	agents = make(map[string]lang.PlaceholderType)
	lock   sync.Mutex
)

// StartAgent starts a opentelemetry agent.
func StartAgent(c Config) {
	lock.Lock()
	defer lock.Unlock()

	_, ok := agents[c.Endpoint]
	if ok {
		return
	}

	// if error happens, let later calls run.
	if err := startAgent(c); err != nil {
		return
	}

	agents[c.Endpoint] = lang.Placeholder
}

func createExporter(c Config) (sdktrace.SpanExporter, error) {
	// Just support jaeger and zipkin now, more for later
	switch c.Batcher {
	case kindJaeger:
		return jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(c.Endpoint)))
	//case kindZipkin:
	//	return zipkin.New(c.Endpoint)
	default:
		return nil, fmt.Errorf("unknown exporter: %s", c.Batcher)
	}
}

func startAgent(c Config) error {
	opts := []sdktrace.TracerProviderOption{
		// Set the sampling rate based on the parent span to 100%
		sdktrace.WithSampler(sdktrace.ParentBased(sdktrace.TraceIDRatioBased(c.Sampler))),
		// Record information about this application in an Resource.
		sdktrace.WithResource(
			resource.NewWithAttributes(semconv.SchemaURL, semconv.ServiceNameKey.String(c.Name),
				attribute.String("server.environment", env.ReleaseMode),
				attribute.String("go.version", runtime.Version()),
				attribute.String("app.version", sysx.AppVersion),
				attribute.String("git.hash", sysx.GitCommitLog),
				attribute.String("system.name", sysx.SubSystem),
			)),
	}

	if len(c.Endpoint) > 0 {
		exp, err := createExporter(c)
		if err != nil {
			logx.Error(err)
			return err
		}

		// Always be sure to batch in production.
		opts = append(opts, sdktrace.WithBatcher(exp))
	}

	tp := sdktrace.NewTracerProvider(opts...)
	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(
		propagation.TraceContext{}, propagation.Baggage{}))
	otel.SetErrorHandler(otel.ErrorHandlerFunc(func(err error) {
		logx.Errorf("[otel] error: %v", err)
	}))

	return nil
}
