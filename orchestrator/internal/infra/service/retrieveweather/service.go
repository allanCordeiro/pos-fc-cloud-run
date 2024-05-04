package retrieveweather

import (
	"context"

	"github.com/allanCordeiro/pos-fc-cloud-run/orchestrator/internal/domain"
)

type Retrieve interface {
	Retrieve(ctx context.Context, city string) (*domain.Temperature, error)
}
