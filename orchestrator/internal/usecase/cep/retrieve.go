package cep

import (
	"context"
	"log"

	"github.com/allanCordeiro/pos-fc-cloud-run/orchestrator/internal/domain"
	"github.com/allanCordeiro/pos-fc-cloud-run/orchestrator/internal/infra/service/retrievecep"
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

type CepErrorsOutput struct {
	Code    int
	Message string
}

func NewRetrieveUseCase(service retrievecep.Service) *RetrieveUseCase {
	return &RetrieveUseCase{Service: service}
}

func (u *RetrieveUseCase) Execute(ctx context.Context, input Input) (Output, *CepErrorsOutput) {
	cep, err := u.Service.Retrieve(ctx, input.Zipcode)
	if err != nil {
		if err == domain.ErrInvalidZipCode {
			return Output{}, &CepErrorsOutput{
				Code:    422,
				Message: err.Error(),
			}
		}
		if err == domain.ErrZipCodeNotFound {
			return Output{}, &CepErrorsOutput{
				Code:    404,
				Message: err.Error(),
			}
		}
		log.Println(err)
		return Output{}, &CepErrorsOutput{
			Code:    500,
			Message: "internal server error. please try again later",
		}
	}

	return Output{
		Address:  cep.Address,
		District: cep.District,
		City:     cep.City,
		State:    cep.State,
	}, nil
}
