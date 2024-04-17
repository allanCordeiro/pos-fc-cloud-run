package impl

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strings"

	"github.com/allanCordeiro/pos-fc-cloud-run/internal/domain"
)

type ViaCep struct {
	Client *http.Client
}

type Output struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func NewViaCep(client *http.Client) *ViaCep {
	return &ViaCep{Client: client}
}

func (v *ViaCep) Retrieve(ctx context.Context, cep string) (*domain.Cep, error) {
	//KLUDGE:: find a way to put this url out of this
	url := "https://viacep.com.br/ws/zipcode/json/"
	url = strings.Replace(url, "zipcode", cep, 1)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	res, err := v.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		if res.StatusCode == http.StatusBadRequest {
			return nil, domain.ErrInvalidZipCode
		}
		return nil, errors.New("http error status code: " + res.Status)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	output, err := parser(body)
	if err != nil {
		return nil, err
	}
	if output.Cep == "" {
		return nil, domain.ErrZipCodeNotFound
	}

	return &domain.Cep{
		Code:     output.Cep,
		Address:  output.Logradouro,
		District: output.Bairro,
		City:     output.Localidade,
		State:    output.Uf,
	}, nil

}

func parser(body []byte) (Output, error) {
	var data Output
	err := json.Unmarshal(body, &data)
	if err != nil {
		return Output{}, err
	}
	return data, nil
}
