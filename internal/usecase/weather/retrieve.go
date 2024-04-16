package weather

import (
	"context"

	"github.com/allanCordeiro/pos-fc-cloud-run/internal/infra/service/retrieveweather"
)

type RetrieveUseCase struct {
	Service retrieveweather.Retrieve
}

type Input struct {
	City string
}

type Output struct {
	Celsius    float64 `json:"temp_C"`
	Fahrenheit float64 `json:"temp_F"`
	Kelvin     float64 `json:"temp_K"`
}

func NewRetrieveUseCase(service retrieveweather.Retrieve) *RetrieveUseCase {
	return &RetrieveUseCase{Service: service}
}

func (r *RetrieveUseCase) Execute(ctx context.Context, input Input) (*Output, error) {
	temperatures, err := r.Service.Retrieve(ctx, input.City)
	if err != nil {
		return nil, err
	}

	return &Output{
		Celsius:    temperatures.Celsius,
		Fahrenheit: temperatures.Fahrenheit,
		Kelvin:     temperatures.Kelvin,
	}, nil
}
