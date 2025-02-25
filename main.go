package main

import (
	"context"
	"time"

	"go.opentelemetry.io/contrib/bridges/otelslog"
	"go.opentelemetry.io/otel/exporters/otlp/otlplog/otlploghttp"
	"go.opentelemetry.io/otel/sdk/log"
)

func main() {
	ctx := context.Background()
	logExporter, err := otlploghttp.New(ctx, otlploghttp.WithInsecure())
	if err != nil {
		panic("failed to create exporter: " + err.Error())
	}

	logProvider := log.NewLoggerProvider(
		log.WithProcessor(log.NewBatchProcessor(logExporter)),
	)

	defer logProvider.Shutdown(ctx)

	logger := otelslog.NewLogger("log-sampler-demo", otelslog.WithLoggerProvider(logProvider))

	counter := 0

	for {
		counter++
		logger.Info("Counter incremented", "value", counter, "timestamp", time.Now())
		time.Sleep(1 * time.Second)
	}
}
