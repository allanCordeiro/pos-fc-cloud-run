package weather

import (
	"context"
	"log"
	"strings"

	"github.com/allanCordeiro/pos-fc-cloud-run/orchestrator/internal/domain"
	"github.com/allanCordeiro/pos-fc-cloud-run/orchestrator/internal/infra/service/retrieveweather"
)

type RetrieveUseCase struct {
	Service retrieveweather.Retrieve
}

type Input struct {
	City string
}

type Output struct {
	City       string  `json:"city"`
	Celsius    float64 `json:"temp_C"`
	Fahrenheit float64 `json:"temp_F"`
	Kelvin     float64 `json:"temp_K"`
}

type WeatherErrorsOutput struct {
	Code    int
	Message string
}

func NewRetrieveUseCase(service retrieveweather.Retrieve) *RetrieveUseCase {
	return &RetrieveUseCase{Service: service}
}

func (r *RetrieveUseCase) Execute(ctx context.Context, input Input) (*Output, *WeatherErrorsOutput) {
	temperatures, err := r.Service.Retrieve(ctx, input.City)
	if err != nil {
		if strings.Contains(err.Error(), "internal") {
			log.Println(err)
			return nil, &WeatherErrorsOutput{
				Code:    500,
				Message: "internal server error. please try again later",
			}
		}
		return nil, &WeatherErrorsOutput{
			Code:    404,
			Message: domain.ErrZipCodeNotFound.Error(),
		}

	}

	return &Output{
		City:       temperatures.City,
		Celsius:    temperatures.Celsius,
		Fahrenheit: temperatures.Fahrenheit,
		Kelvin:     temperatures.Kelvin,
	}, nil
}
