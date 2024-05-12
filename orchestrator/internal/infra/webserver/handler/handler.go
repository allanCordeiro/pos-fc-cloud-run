package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/allanCordeiro/pos-fc-cloud-run/orchestrator/internal/usecase/cep"
	"github.com/allanCordeiro/pos-fc-cloud-run/orchestrator/internal/usecase/weather"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
)

type WeatherHandler struct {
	ZipCodeUseCase cep.RetrieveUseCase
	WeatherUseCase weather.RetrieveUseCase
}

func NewWeatherHandler(zipcode cep.RetrieveUseCase, weather weather.RetrieveUseCase) *WeatherHandler {
	return &WeatherHandler{
		ZipCodeUseCase: zipcode,
		WeatherUseCase: weather,
	}
}

func (h *WeatherHandler) GetWeather(w http.ResponseWriter, r *http.Request) {
	tracer := otel.Tracer("servico B")
	carrier := propagation.HeaderCarrier(r.Header)

	w.Header().Set("Content-Type", "application/json")
	ctx, cancel := context.WithTimeout(r.Context(), time.Second*10)
	defer cancel()

	ctx = otel.GetTextMapPropagator().Extract(ctx, carrier)
	ctx, span := tracer.Start(ctx, "busca-temperatura")
	span.End()

	zipcodeInput := cep.Input{Zipcode: r.PathValue("zipcode")}

	ctx, span = tracer.Start(ctx, "requisicao para o viacep")
	zipcodeOutput, errCep := h.ZipCodeUseCase.Execute(ctx, zipcodeInput)
	if errCep != nil {
		w.WriteHeader(errCep.Code)
		json.NewEncoder(w).Encode(errCep.Message)
		return
	}
	span.End()

	ctx, span = tracer.Start(ctx, "requisicao para o weatherapi")
	weatherInput := weather.Input{City: zipcodeOutput.City}
	weatherOutput, errWeather := h.WeatherUseCase.Execute(ctx, weatherInput)
	if errWeather != nil {
		w.WriteHeader(errWeather.Code)
		json.NewEncoder(w).Encode(errWeather.Message)
		return
	}
	span.End()

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(&weatherOutput)
}
