package handler

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/allanCordeiro/pos-fc-cloud-run/orchestrator/pkg"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
)

type Input struct {
	Cep string `json:"cep"`
}

type Output struct {
	City       string  `json:"city"`
	Celsius    float64 `json:"temp_C"`
	Fahrenheit float64 `json:"temp_F"`
	Kelvin     float64 `json:"temp_K"`
}

type ErrorOutput struct {
	Code    int
	Message string
}

func EntranceHandler(w http.ResponseWriter, r *http.Request) {
	tracer := otel.Tracer("servico A")
	ctx := r.Context()
	ctx, span := tracer.Start(ctx, "valida-cep")
	span.End()

	var input Input
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		err = json.NewEncoder(w).Encode(&ErrorOutput{
			Code:    http.StatusUnprocessableEntity,
			Message: "invalid zipcode",
		})
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		return
	}
	cep := pkg.NewCep(input.Cep)

	if !cep.IsCepCodeValid() {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(&ErrorOutput{
			Code:    http.StatusBadRequest,
			Message: "invalid zipcode",
		})
		return
	}
	//KLUDGE:: find a way to put this url out of this
	url := "http://orchestrator-api:8080/weather/" + cep.GetCode()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	ctx, serviceBSpan := tracer.Start(ctx, "calling service B")
	otel.GetTextMapPropagator().Inject(ctx, propagation.HeaderCarrier(req.Header))

	req.Header.Set("accept", "application/json")
	req = req.WithContext(ctx)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer serviceBSpan.End()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if res.StatusCode != http.StatusOK {
		w.WriteHeader(res.StatusCode)
		var errorOutput ErrorOutput
		err = json.Unmarshal(body, &errorOutput)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
	var output Output
	err = json.Unmarshal(body, &output)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(&output)
}
