package main

import (
	"log"
	"net/http"
	"os"

	"github.com/allanCordeiro/pos-fc-cloud-run/temperature-service/internal/infra/handler"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/zipkin"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
)

var logger = log.New(os.Stderr, "servico-cep", log.Ldate|log.Ltime|log.Llongfile)

func main() {
	initTracer()

	http.HandleFunc("POST /temperatura", handler.EntranceHandler)

	log.Fatal(http.ListenAndServe(":8082", nil))
}

func initTracer() {
	exporter, err := zipkin.New("http://zipkin:9411/api/v2/spans", zipkin.WithLogger(logger))
	if err != nil {
		log.Fatalf("failed to instantiate zipkin: %v", err)
	}

	tp := trace.NewTracerProvider(
		trace.WithBatcher(exporter),
		trace.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String("temperatura-service"),
		)),
	)

	otel.SetTracerProvider(tp)

	otel.SetTextMapPropagator(propagation.TraceContext{})
}
