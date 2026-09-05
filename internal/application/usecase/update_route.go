package usecase

import (
	"context"

	"github.com/google/uuid"
	"github.com/riolivre/nautical_logistics/internal/application/ports/input"
	"github.com/riolivre/nautical_logistics/internal/application/ports/output"
)

type updateRouteUseCase struct {
	routeRepo output.RouteRepository
}

func NewUpdateRouteUseCase(routeRepo output.RouteRepository) input.UpdateRouteUseCase {
	return &updateRouteUseCase{
		routeRepo: routeRepo,
	}
}

func (uc *updateRouteUseCase) Execute(ctx context.Context, id string, inputRoute input.UpdateRouteInput) error {
	// Fetch the existing route
	routeID, err := uuid.Parse(id)
	if err != nil {
		return err
	}

	route, err := uc.routeRepo.GetRoute(ctx, routeID)
	if err != nil {
		return err
	}

	// Update the route fields
	route.Name = inputRoute.Name
	route.OrganizationID, err = uuid.Parse(inputRoute.OrganizationID)
	route.UpdatedAt = ptrTimeNow()
	if err != nil {
		return err
	}

	if err := uc.routeRepo.UpdateRoute(ctx, route); err != nil {
		return err
	}

	return nil
}
