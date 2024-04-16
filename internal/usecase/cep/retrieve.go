package cep

import (
	"context"

	"github.com/allanCordeiro/pos-fc-cloud-run/internal/infra/service/retrievecep"
)

type RetrieveUseCase struct {
	Service retrievecep.Service
}

type Input struct {
	Zipcode string
}

type Output struct {
	Address  string
	District string
	City     string
	State    string
}

func NewRetrieveUseCase(service retrievecep.Service) *RetrieveUseCase {
	return &RetrieveUseCase{Service: service}
}

func (u *RetrieveUseCase) Execute(ctx context.Context, input Input) (Output, error) {
	cep, err := u.Service.Retrieve(ctx, input.Zipcode)
	if err != nil {
		return Output{}, err
	}

	return Output{
		Address:  cep.Address,
		District: cep.District,
		City:     cep.City,
		State:    cep.State,
	}, nil
}
