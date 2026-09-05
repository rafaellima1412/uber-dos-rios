package output

import (
	"context"

	"github.com/google/uuid"
	"github.com/rafaellima1412/uber-dos-rios/internal/application/domain"
)

// TripRepository defines the interface for trip data persistence.
type TripRepository interface {
	CreateTrip(ctx context.Context, trip *domain.Trip) error
	GetTrip(ctx context.Context, id uuid.UUID) (*domain.Trip, error)
	UpdateTrip(ctx context.Context, trip *domain.Trip) error
	DeleteTrip(ctx context.Context, id uuid.UUID) error
	ListTrips(ctx context.Context, limit int, offset int) ([]*domain.Trip, error)
	FilterTrips(ctx context.Context, trip *domain.TripFilter, limit int, offset int) ([]*domain.Trip, error)
}
