package tracing

import (
	"context"
	"github.com/rs/zerolog/log"
	"github.com/theykk/golactic/pkg/config"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/trace/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/trace"
	"os"
	"runtime"
	"time"
)

func InitTracer(conf *config.Configuration) trace.Tracer {
	exp, err := jaeger.NewRawExporter(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(conf.Tracer.URl)))
	if err != nil {
		log.Fatal().Err(err)
	}
	hostname, _ := os.Hostname()
	tp := sdktrace.NewTracerProvider(sdktrace.WithBatcher(exp,
		sdktrace.WithBatchTimeout(time.Second*3),
		sdktrace.WithMaxExportBatchSize(conf.Tracer.BatchSize)),

		sdktrace.WithSampler(sdktrace.AlwaysSample()),
		// tüm Span'lere eklenicek attribute'ler
		sdktrace.WithResource(resource.NewWithAttributes(
			attribute.String("service.name", conf.Service.Name),
			attribute.String("service.version", conf.Service.Version),
			attribute.String("instance.id", hostname),
			attribute.String("runtime.os", runtime.GOOS),
			attribute.String("runtime.arch", runtime.GOARCH),
			attribute.String("runtime.go", runtime.Version()),
		)),
	)
	otel.SetTracerProvider(tp)

	defer func() {
		if err := tp.ForceFlush(context.Background()); err != nil {
			log.Fatal().Msg("failed to flush tracer")
		}
	}()
	return otel.GetTracerProvider().Tracer(conf.Service.Project)
}
