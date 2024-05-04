package impl_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/allanCordeiro/pos-fc-cloud-run/orchestrator/internal/domain"
	"github.com/allanCordeiro/pos-fc-cloud-run/orchestrator/internal/infra/service/retrievecep/impl"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func TestViaCep_Retrieve(t *testing.T) {
	t.Run("given a valid cep when retrieve then should return zipcode response", func(t *testing.T) {
		// Configuração do mock HTTP
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		// Mock da resposta HTTP
		mockResponse := `{
			"cep": "12345-678",
			"logradouro": "Rua Teste",
			"complemento": "",
			"bairro": "Bairro Teste",
			"localidade": "Cidade Teste",
			"uf": "TS",
			"ibge": "0000000",
			"gia": "0000",
			"ddd": "00",
			"siafi": "0000"
		}`
		mockURL := "http://viacep.com.br/ws/12345678/json/"
		httpmock.RegisterResponder(http.MethodGet, mockURL,
			httpmock.NewStringResponder(http.StatusOK, mockResponse))

		// Configuração do cliente HTTP para usar o mock
		httpClient := &http.Client{Transport: httpmock.DefaultTransport}

		// Criação da instância ViaCep com o cliente HTTP mockado
		viaCep := impl.NewViaCep(httpClient)

		// Teste do método Retrieve
		result, err := viaCep.Retrieve(context.Background(), "12345678")
		assert.NoError(t, err)
		assert.Equal(t, &domain.Cep{
			Code:     "12345-678",
			Address:  "Rua Teste",
			District: "Bairro Teste",
			City:     "Cidade Teste",
			State:    "TS",
		}, result)
	})

	t.Run("given an unknown zipcode when retrieve should return error cannot find zipcode", func(t *testing.T) {
		// Configuração do mock HTTP
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		// Mock da resposta HTTP
		mockResponse := `{
			"erro": true
		}`
		mockURL := "http://viacep.com.br/ws/01112100/json/"
		httpmock.RegisterResponder(http.MethodGet, mockURL,
			httpmock.NewStringResponder(http.StatusOK, mockResponse))

		// Configuração do cliente HTTP para usar o mock
		httpClient := &http.Client{Transport: httpmock.DefaultTransport}

		// Criação da instância ViaCep com o cliente HTTP mockado
		viaCep := impl.NewViaCep(httpClient)

		// Teste do método Retrieve
		_, err := viaCep.Retrieve(context.Background(), "01112100")
		assert.Error(t, err)
		assert.Equal(t, domain.ErrZipCodeNotFound, err)
	})

	t.Run("given an invalid zipcode when retrieve should return status 400", func(t *testing.T) {
		// Configuração do mock HTTP
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		// Mock da resposta HTTP
		mockURL := "http://viacep.com.br/ws/011121000/json/"
		httpmock.RegisterResponder(http.MethodGet, mockURL,
			httpmock.NewStringResponder(http.StatusBadRequest, ""))

		// Configuração do cliente HTTP para usar o mock
		httpClient := &http.Client{Transport: httpmock.DefaultTransport}

		// Criação da instância ViaCep com o cliente HTTP mockado
		viaCep := impl.NewViaCep(httpClient)

		// Teste do método Retrieve
		_, err := viaCep.Retrieve(context.Background(), "011121000")
		assert.Error(t, err)
		assert.Equal(t, domain.ErrInvalidZipCode, err)
	})

}
