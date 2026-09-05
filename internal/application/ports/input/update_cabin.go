package input

import (
	"context"

	"github.com/google/uuid"
)

type UpdateCabinInput struct {
	ID          uuid.UUID
	Name        string
	BedType     string
	Capacity    int
	Description *string
}

type UpdateCabinOutput struct {
	ID          uuid.UUID
	Name        string
	BedType     string
	Capacity    int
	Description *string
}

type UpdateCabinUseCase interface {
	Execute(ctx context.Context, input *UpdateCabinInput, id uuid.UUID) error
}
