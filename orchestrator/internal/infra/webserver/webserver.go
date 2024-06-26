package webserver

import (
	"log"
	"net/http"

	"github.com/allanCordeiro/pos-fc-cloud-run/orchestrator/internal/infra/webserver/handler"
	"github.com/allanCordeiro/pos-fc-cloud-run/orchestrator/internal/usecase/cep"
	"github.com/allanCordeiro/pos-fc-cloud-run/orchestrator/internal/usecase/weather"
)

func Serve(port string, cepUseCase cep.RetrieveUseCase, weatherUseCase weather.RetrieveUseCase) {
	weatherHandler := handler.NewWeatherHandler(cepUseCase, weatherUseCase)
	http.HandleFunc("GET /weather/{zipcode}", weatherHandler.GetWeather)

	log.Fatal(http.ListenAndServe(":"+port, nil))
}
