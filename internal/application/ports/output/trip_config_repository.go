package output

import (
	"context"

	"github.com/google/uuid"
	"github.com/riolivre/nautical_logistics/internal/application/domain"
	"github.com/riolivre/nautical_logistics/internal/application/ports/input"
)

// TripRepository defines the interface for trip data persistence.
type TripConfigRepository interface {
	CreateTripConfig(ctx context.Context, trip *domain.TripConfig) error
	UpdateTripConfig(ctx context.Context, trip *domain.TripConfig) error
	DeleteTripConfig(ctx context.Context, id uuid.UUID) error
	ListTripsConfig(ctx context.Context, limit int, offset int) ([]*domain.TripConfig, error)
	GetTripConfig(ctx context.Context, id uuid.UUID) (*domain.TripConfig, error)
	LoadGraph(ctx context.Context, DateOrigin, DateDest string) (input.Graph, error)
}
