package input

import (
	"context"

	"github.com/google/uuid"
)

type DeleteSeatUseCase interface {
	Execute(ctx context.Context, id uuid.UUID) error
}