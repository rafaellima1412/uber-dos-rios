package input

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type CreateReservationInput struct {
	UserID          uuid.UUID
	ReservationDate *time.Time
	// Status          domain.ReservationStatus
}

type CreateReservationUseCase interface {
	Execute(ctx context.Context, input *CreateReservationInput) (string, error)
}
