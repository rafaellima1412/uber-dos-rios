package input

import (
	"context"

	"github.com/google/uuid"
)

type DeleteCabinUseCase interface {
	Execute(ctx context.Context, id uuid.UUID) error
}