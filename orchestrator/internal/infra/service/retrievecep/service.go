package retrievecep

import (
	"context"

	"github.com/allanCordeiro/pos-fc-cloud-run/internal/domain"
)

type Service interface {
	Retrieve(ctx context.Context, cep string) (*domain.Cep, error)
}
