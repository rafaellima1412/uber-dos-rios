package output

import (
	"context"

	"github.com/google/uuid"
	"github.com/rafaellima1412/uber-dos-rios/internal/application/domain"
)

// RouteRepository defines the interface for route data persistence.
type RouteRepository interface {
	CreateRoute(ctx context.Context, route *domain.Route) error
	GetRoute(ctx context.Context, id uuid.UUID) (*domain.Route, error)
	UpdateRoute(ctx context.Context, route *domain.Route) error
	DeleteRoute(ctx context.Context, id uuid.UUID) error
	ListRoutes(ctx context.Context, limit int, offset int) ([]*domain.Route, error)
	SearchRoutes(ctx context.Context, query string, limit int, offset int) ([]*domain.Route, error)
}
