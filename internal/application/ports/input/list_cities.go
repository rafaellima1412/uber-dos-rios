package input

import (
	"context"
)


type CityOutput struct {
	ID   int
	Name string
	State  string
}

type ListCitiesUseCase interface {
	Execute(ctx context.Context) ([]CityOutput, error)
}
