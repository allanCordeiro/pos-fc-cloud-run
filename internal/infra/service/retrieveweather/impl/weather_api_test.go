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
		expectedCelsius := 28.0
		expectedFahreinheit := 82.4
		expectedKelvin := 301.0
		// Configuração do mock HTTP
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		// Mock da resposta HTTP
		mockResponse := `{
			"location": {
				"name": "Mexico City",
				"region": "The Federal District",
				"country": "Mexico",
				"lat": 19.43,
				"lon": -99.13,
				"tz_id": "America/Mexico_City",
				"localtime_epoch": 1713305382,
				"localtime": "2024-04-16 16:09"
			},
			"current": {
				"last_updated_epoch": 1713304800,
				"last_updated": "2024-04-16 16:00",
				"temp_c": 28.0,
				"temp_f": 82.4,
				"is_day": 1,
				"condition": {
					"text": "Partly cloudy",
					"icon": "//cdn.weatherapi.com/weather/64x64/day/116.png",
					"code": 1003
				},
				"wind_mph": 16.1,
				"wind_kph": 25.9,
				"wind_degree": 160,
				"wind_dir": "SSE",
				"pressure_mb": 1025.0,
				"pressure_in": 30.27,
				"precip_mm": 0.0,
				"precip_in": 0.0,
				"humidity": 16,
				"cloud": 50,
				"feelslike_c": 26.0,
				"feelslike_f": 78.8,
				"vis_km": 8.0,
				"vis_miles": 4.0,
				"uv": 8.0,
				"gust_mph": 20.6,
				"gust_kph": 33.1
			}
		}`
		mockURL := "https://api.weatherapi.com/v1/current.json?key=3a86487cb0804004a3b10835241004&q=mexico city"
		httpmock.RegisterResponder(http.MethodGet, mockURL,
			httpmock.NewStringResponder(http.StatusOK, mockResponse))

		// Configuração do cliente HTTP para usar o mock
		httpClient := &http.Client{Transport: httpmock.DefaultTransport}

		weatherApi := NewWeatherApi(httpClient)

		// Teste do método Retrieve
		result, err := weatherApi.Retrieve(context.Background(), "mexico city")
		assert.NoError(t, err)
		assert.Equal(t, expectedCelsius, result.Celsius)
		assert.Equal(t, expectedFahreinheit, result.Fahrenheit)
		assert.Equal(t, expectedKelvin, result.Kelvin)
	})
}
