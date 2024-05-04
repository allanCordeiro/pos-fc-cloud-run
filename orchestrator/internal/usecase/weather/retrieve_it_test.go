package weather

import (
	"context"
	"net/http"
	"testing"

	"github.com/allanCordeiro/pos-fc-cloud-run/orchestrator/internal/domain"
	"github.com/allanCordeiro/pos-fc-cloud-run/orchestrator/internal/infra/service/retrieveweather/impl"
	"github.com/stretchr/testify/assert"
)

func TestIntegratedWeatherUseCase(t *testing.T) {
	scenarios := []struct {
		name        string
		city        string
		expectedErr *WeatherErrorsOutput
	}{
		{
			name:        "given a valid city code when calls weather usecase should return temperature data",
			city:        "São Paulo",
			expectedErr: nil,
		},
		{
			name:        "given an unknow city code when calls weather usecase should return an error code",
			city:        "Sulvador",
			expectedErr: &WeatherErrorsOutput{Code: 404, Message: domain.ErrZipCodeNotFound.Error()},
		},
	}

	for _, test := range scenarios {
		t.Run(test.name, func(t *testing.T) {
			service := impl.NewWeatherApi(http.DefaultClient)
			usecase := NewRetrieveUseCase(service)
			output, err := usecase.Execute(context.TODO(), Input{City: test.city})
			if err != nil {
				assert.Equal(t, test.expectedErr.Code, err.Code)
				assert.Equal(t, test.expectedErr.Message, err.Message)
				assert.Nil(t, output)
			}
			if err == nil {
				assert.Equal(t, test.expectedErr, err)
				assert.NotNil(t, output)
			}
		})
	}
}
