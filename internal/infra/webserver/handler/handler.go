package handler

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/allanCordeiro/pos-fc-cloud-run/internal/usecase/cep"
	"github.com/allanCordeiro/pos-fc-cloud-run/internal/usecase/weather"
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
	w.Header().Set("Content-Type", "application/json")
	ctx, cancel := context.WithTimeout(r.Context(), time.Second*10)
	defer cancel()

	zipcodeInput := cep.Input{Zipcode: r.PathValue("zipcode")}

	zipcodeOutput, errCep := h.ZipCodeUseCase.Execute(ctx, zipcodeInput)
	if errCep != nil {
		w.WriteHeader(errCep.Code)
		json.NewEncoder(w).Encode(errCep.Message)
		return
	}

	weatherInput := weather.Input{City: zipcodeOutput.City}
	weatherOutput, err := h.WeatherUseCase.Execute(ctx, weatherInput)
	if err != nil {
		log.Println(err)
		//TODO send back error to response
	}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(&weatherOutput)
}
