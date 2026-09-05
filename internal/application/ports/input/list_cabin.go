package input

import (
	"context"

	"github.com/google/uuid"
)

type ListCabinInput struct {
	Page    int
	PerPage int
}

type CabinOutput struct {
	ID          uuid.UUID
	Name        string
	BedType     string
	Capacity    int
	Description *string
}

type ListCabinOutput struct {
	Cabins      []CabinOutput
	TotalCount  int
	TotalPages  int
	CurrentPage int
}

type ListShipCabinUseCase interface {
	Execute(ctx context.Context, inputDTo ListCabinInput) (ListCabinOutput, error)
}
