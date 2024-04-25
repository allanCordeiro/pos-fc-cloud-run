package impl

import (
	"context"
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func TestViaCep_Retrieve(t *testing.T) {
	t.Run("given a valid city when retrieve then should return weather temperatures", func(t *testing.T) {
		expectedCity := "Salvador"
		expectedCelsius := 28.0
		expectedFahreinheit := 82.4
		expectedKelvin := 301.0
		// Configuração do mock HTTP
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		// Mock da resposta HTTP
		mockResponse := `{
			"location": {
				"name": "Salvador",
				"region": "Bahia",
				"country": "Brazil",
				"lat": -12.98,
				"lon": -38.52,
				"tz_id": "America/Bahia",
				"localtime_epoch": 1713392700,
				"localtime": "2024-04-17 19:25"
			},
			"current": {
				"temp_c": 28.0,
				"temp_f": 82.4,
				"condition": {}
			}
		}`
		mockURL := "http://api.weatherapi.com/v1/current.json?key=3a86487cb0804004a3b10835241004&q=salvador"
		httpmock.RegisterResponder(http.MethodGet, mockURL,
			httpmock.NewStringResponder(http.StatusOK, mockResponse))

		// Configuração do cliente HTTP para usar o mock
		httpClient := &http.Client{Transport: httpmock.DefaultTransport}

		weatherApi := NewWeatherApi(httpClient)

		// Teste do método Retrieve
		result, err := weatherApi.Retrieve(context.Background(), "salvador")
		assert.NoError(t, err)
		assert.Equal(t, expectedCity, result.City)
		assert.Equal(t, expectedCelsius, result.Celsius)
		assert.Equal(t, expectedFahreinheit, result.Fahrenheit)
		assert.Equal(t, expectedKelvin, result.Kelvin)
	})
}
