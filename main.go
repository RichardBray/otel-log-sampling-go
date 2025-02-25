package main

import (
	"context"
	"fmt"
	"time"

	"go.opentelemetry.io/contrib/bridges/otelslog"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/exporters/otlp/otlplog/otlploghttp"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/sdk/log"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.26.0"
)

func main() {
	ctx := context.Background()

	res, err := newResource()
	if err != nil {
		panic("failed to create resource: " + err.Error())
	}

	// Logs

	logExporter, err := otlploghttp.New(ctx, otlploghttp.WithInsecure())
	if err != nil {
		panic("failed to create exporter: " + err.Error())
	}

	logProvider := log.NewLoggerProvider(
		log.WithResource(res),
		log.WithProcessor(log.NewBatchProcessor(logExporter)),
	)

	defer logProvider.Shutdown(ctx)

	logger := otelslog.NewLogger("log-sampler-demo", otelslog.WithLoggerProvider(logProvider))

	// Traces

	traceExporter, err := otlptracehttp.New(ctx, otlptracehttp.WithInsecure())
	if err != nil {
		panic("failed to create trace exporter: " + err.Error())
	}

	traceProvider := trace.NewTracerProvider(
		trace.WithBatcher(traceExporter),
		trace.WithResource(res),
		trace.WithSampler(trace.AlwaysSample()),
	)

	defer traceProvider.Shutdown(ctx)

	tracer := traceProvider.Tracer("demo-tracer")

	// Counter

	counter := 0

	for {
		counter++
		ctx, span := tracer.Start(ctx, "increment-counter")

		span.SetAttributes(attribute.Int("counter", counter))

		// logger.InfoContext(ctx, "Counter incremented", "value", counter, "timestamp", time.Now())

		if counter%5 == 0 {
			err := fmt.Errorf("simulated error on count %d", counter)
			span.SetStatus(codes.Error, err.Error())
			span.RecordError(err)
			logger.ErrorContext(ctx, "Counter increment failed",
				"error", err,
				"value", counter,
				"timestamp", time.Now())
		} else {
			logger.InfoContext(ctx, "Counter incremented",
				"value", counter,
				"timestamp", time.Now())
		}

		time.Sleep(1 * time.Second)

		span.End()
	}
}

func newResource() (*resource.Resource, error) {
	return resource.Merge(resource.Default(),
		resource.NewWithAttributes(semconv.SchemaURL,
			semconv.ServiceName("log-sampling-go"),
			semconv.ServiceVersion("0.1.0"),
		))
}
