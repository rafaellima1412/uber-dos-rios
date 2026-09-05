package usecase

import (
	"context"

	"github.com/google/uuid"
	"github.com/rafaellima1412/uber-dos-rios/internal/application/ports/input"
	"github.com/rafaellima1412/uber-dos-rios/internal/application/ports/output"
)

type deleteRouteUseCase struct {
	routeRepo output.RouteRepository
}

func NewDeleteRouteUseCase(routeRepo output.RouteRepository) input.DeleteRouteUseCase {
	return &deleteRouteUseCase{
		routeRepo: routeRepo,
	}
}

// Execute deletes a route by its ID.
func (uc *deleteRouteUseCase) Execute(ctx context.Context, id string) error {
	routeID, err := uuid.Parse(id)
	if err != nil {
		return err
	}
	return uc.routeRepo.DeleteRoute(ctx, routeID)
}
