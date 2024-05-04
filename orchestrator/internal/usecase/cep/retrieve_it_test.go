package cep

import (
	"context"
	"net/http"
	"testing"

	"github.com/allanCordeiro/pos-fc-cloud-run/orchestrator/internal/domain"
	"github.com/allanCordeiro/pos-fc-cloud-run/orchestrator/internal/infra/service/retrievecep/impl"
	"github.com/stretchr/testify/assert"
)

func TestIntegratedCepUseCase(t *testing.T) {
	scenarios := []struct {
		name         string
		zipcode      string
		expectedCity string
		expectedErr  *CepErrorsOutput
	}{
		{
			name:         "given a valid zip code when calls cep usecase should return city",
			zipcode:      "01211100",
			expectedCity: "São Paulo",
			expectedErr:  nil,
		},
		{
			name:         "given an invalid zip code when calls cep usecase should return error invalid zip code",
			zipcode:      "012111000000",
			expectedCity: "",
			expectedErr:  &CepErrorsOutput{Code: 422, Message: domain.ErrInvalidZipCode.Error()},
		},
		{
			name:         "given an unknow zip code when calls cep usecase should return error zip code not found",
			zipcode:      "55555050",
			expectedCity: "",
			expectedErr:  &CepErrorsOutput{Code: 404, Message: domain.ErrZipCodeNotFound.Error()},
		},
	}

	for _, test := range scenarios {
		t.Run(test.name, func(t *testing.T) {
			service := impl.NewViaCep(http.DefaultClient)
			usecase := NewRetrieveUseCase(service)

			output, err := usecase.Execute(context.TODO(), Input{Zipcode: test.zipcode})
			if err != nil {
				assert.Equal(t, test.expectedErr.Code, err.Code)
				assert.Equal(t, test.expectedErr.Message, err.Message)
			}
			if err == nil {
				assert.Equal(t, test.expectedErr, err)
				assert.Equal(t, test.expectedCity, output.City)
			}
		})
	}
}
