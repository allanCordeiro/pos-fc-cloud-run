package webserver

import (
	"log"
	"net/http"

	"github.com/allanCordeiro/pos-fc-cloud-run/internal/infra/webserver/handler"
	"github.com/allanCordeiro/pos-fc-cloud-run/internal/usecase/cep"
	"github.com/allanCordeiro/pos-fc-cloud-run/internal/usecase/weather"
)

func Serve(port string, cepUseCase cep.RetrieveUseCase, weatherUseCase weather.RetrieveUseCase) {
	weatherHandler := handler.NewWeatherHandler(cepUseCase, weatherUseCase)
	http.HandleFunc("GET /weather/{zipcode}", weatherHandler.GetWeather)

	log.Fatal(http.ListenAndServe(":"+port, nil))
}
