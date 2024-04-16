package main

import (
	"context"
	"fmt"
	"net/http"

	viacep "github.com/allanCordeiro/pos-fc-cloud-run/internal/infra/service/retrievecep/impl"
	weatherApi "github.com/allanCordeiro/pos-fc-cloud-run/internal/infra/service/retrieveweather/impl"
	"github.com/allanCordeiro/pos-fc-cloud-run/internal/usecase/cep"
	"github.com/allanCordeiro/pos-fc-cloud-run/internal/usecase/weather"
)

func main() {
	searchCep := viacep.NewViaCep(http.DefaultClient)
	usecaseCep := cep.NewRetrieveUseCase(searchCep)
	cep := cep.Input{Zipcode: "04266-060"}
	output, err := usecaseCep.Execute(context.TODO(), cep)
	if err != nil {
		panic(err)
	}

	fmt.Println(output)

	searchCity := weatherApi.NewWeatherApi(http.DefaultClient)
	useCaseWeather := weather.NewRetrieveUseCase(searchCity)
	weatherOutput, err := useCaseWeather.Execute(context.TODO(), weather.Input{City: output.City})
	if err != nil {
		panic(err)
	}

	fmt.Println(weatherOutput)

}
