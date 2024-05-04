package main

import (
	"net/http"

	viacep "github.com/allanCordeiro/pos-fc-cloud-run/orchestrator/internal/infra/service/retrievecep/impl"
	weatherApi "github.com/allanCordeiro/pos-fc-cloud-run/orchestrator/internal/infra/service/retrieveweather/impl"
	"github.com/allanCordeiro/pos-fc-cloud-run/orchestrator/internal/infra/webserver"
	"github.com/allanCordeiro/pos-fc-cloud-run/orchestrator/internal/usecase/cep"
	"github.com/allanCordeiro/pos-fc-cloud-run/orchestrator/internal/usecase/weather"
)

func main() {
	searchCep := viacep.NewViaCep(http.DefaultClient)
	usecaseCep := cep.NewRetrieveUseCase(searchCep)
	searchCity := weatherApi.NewWeatherApi(http.DefaultClient)
	useCaseWeather := weather.NewRetrieveUseCase(searchCity)

	webserver.Serve("8080", *usecaseCep, *useCaseWeather)

}
