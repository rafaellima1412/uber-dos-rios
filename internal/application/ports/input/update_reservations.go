// UpdateReservationUseCase defines the interface for updating an existing Reservation
package input

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type UpdateReservationOutput struct {
	ID              string
	UserID          string
	TripConfigID    []string
	OccupiedSeats   []string
	ReservationDate string
	Status          string
}

type UpdateReservationInput struct {
	UserID          string
	TripConfigID    []string
	OccupiedSeats   []string
	ReservationDate *time.Time
	Status          string
}

type UpdateReservationUseCase interface {
	Execute(ctx context.Context, input *UpdateReservationInput, id uuid.UUID) error
}
