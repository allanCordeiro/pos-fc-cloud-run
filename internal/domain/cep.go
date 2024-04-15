package domain

import "github.com/allanCordeiro/pos-fc-cloud-run/pkg"

type Cep struct {
	Code string
}

func NewCep(cep string) *Cep {
	return &Cep{Code: cep}
}

func (c *Cep) Validate() error {
	cep := pkg.NewCep(c.Code)
	if !cep.IsCepCodeValid() {
		return ErrInvalidZipCode
	}
	return nil
}
