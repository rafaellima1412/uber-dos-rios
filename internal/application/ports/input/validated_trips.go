package input

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type ValidatedTripsUseCase interface {
	Execute(ctx context.Context, departure, arrival string, occupiedSeats []string, id uuid.UUID) (*time.Time,
		*time.Time, error)
}
