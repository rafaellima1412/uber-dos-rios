package output

import (
	"context"

	"github.com/google/uuid"
	"github.com/riolivre/nautical_logistics/internal/application/domain"
)

// ReservationRepository
type ReservationRepository interface {
	CreateReservation(ctx context.Context, reservation *domain.Reservation) (string, error)
	UpdateReservation(ctx context.Context, reservation *domain.Reservation) error
	GetReservation(ctx context.Context, id uuid.UUID) (*domain.Reservation, error)
	DeleteReservation(ctx context.Context, id uuid.UUID) error
	ListReservations(ctx context.Context, limit int, offset int) ([]*domain.Reservation, error)
	CreateTripReservation(ctx context.Context, tripInstance *domain.Trip, reservation *domain.TripToReservation) error
}
