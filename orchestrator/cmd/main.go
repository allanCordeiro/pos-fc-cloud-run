package main

import (
	"log"
	"net/http"
	"os"

	viacep "github.com/allanCordeiro/pos-fc-cloud-run/orchestrator/internal/infra/service/retrievecep/impl"
	weatherApi "github.com/allanCordeiro/pos-fc-cloud-run/orchestrator/internal/infra/service/retrieveweather/impl"
	"github.com/allanCordeiro/pos-fc-cloud-run/orchestrator/internal/infra/webserver"
	"github.com/allanCordeiro/pos-fc-cloud-run/orchestrator/internal/usecase/cep"
	"github.com/allanCordeiro/pos-fc-cloud-run/orchestrator/internal/usecase/weather"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/zipkin"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
)

var logger = log.New(os.Stderr, "orquestrador", log.Ldate|log.Ltime|log.Llongfile)

func main() {
	initTracer()
	searchCep := viacep.NewViaCep(http.DefaultClient)
	usecaseCep := cep.NewRetrieveUseCase(searchCep)
	searchCity := weatherApi.NewWeatherApi(http.DefaultClient)
	useCaseWeather := weather.NewRetrieveUseCase(searchCity)

	webserver.Serve("8080", *usecaseCep, *useCaseWeather)

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
