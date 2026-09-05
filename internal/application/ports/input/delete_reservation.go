package input

import (
	"context"

	"github.com/google/uuid"
)

type DeleteReservationUseCase interface {
	Execute(ctx context.Context, id uuid.UUID) error
}