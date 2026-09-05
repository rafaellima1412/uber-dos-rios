package output

import (
	"context"

	"github.com/google/uuid"
	"github.com/riolivre/nautical_logistics/internal/application/domain"
)

// ShipRepository defines the interface for ship data persistence.
type ShipRepository interface {
	CreateShip(ctx context.Context, ship *domain.Ship) error
	GetShip(ctx context.Context, id uuid.UUID, typeship string) (*domain.Ship, error)
	UpdateShip(ctx context.Context, ship *domain.Ship) error
	DeleteShip(ctx context.Context, id uuid.UUID) error
	ListShips(ctx context.Context, limit, offset int, typeship string) ([]*domain.Ship, error)
	SearchShips(ctx context.Context, query string, limit, offset int, typeship string) ([]*domain.Ship, error)
}
