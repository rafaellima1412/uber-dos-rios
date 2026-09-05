package usecase

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/rafaellima1412/uber-dos-rios/internal/application/domain"
	"github.com/rafaellima1412/uber-dos-rios/internal/application/ports/input"
	"github.com/rafaellima1412/uber-dos-rios/internal/application/ports/output"
	"github.com/rafaellima1412/uber-dos-rios/internal/logger"
	"go.uber.org/zap"
)

// createRouteUseCase implements the CreateRouteUseCase interface.
type createRouteUseCase struct {
	routeRepo output.RouteRepository
}

// Ensure createRouteUseCase implements input.CreateRouteUseCase.
var _ input.CreateRouteUseCase = (*createRouteUseCase)(nil)

// NewCreateRouteUseCase creates a new instance of CreateRouteUseCase.
func NewCreateRouteUseCase(routeRepo output.RouteRepository) input.CreateRouteUseCase {
	return &createRouteUseCase{
		routeRepo: routeRepo,
	}
}

// Execute implements input.CreateRouteUseCase.
func (u *createRouteUseCase) Execute(ctx context.Context, inputRoute input.CreateRouteInput) error {
	routeID, err := uuid.NewUUID()
	if err != nil {
		logger.Error("failed to generate UUID for route", zap.Error(err))
		return err
	}

	organizationID, err := uuid.Parse(inputRoute.OrganizationID)
	if err != nil {
		logger.Error("failed to parse organization ID", zap.Error(err))
		return err
	}
	route := &domain.Route{
		ID:             routeID,
		Name:           inputRoute.Name,
		Active:         true,
		OrganizationID: organizationID,
		CreatedAt:      ptrTimeNow(),
		UpdatedAt:      ptrTimeNow(),
	}

	if err := u.routeRepo.CreateRoute(ctx, route); err != nil {
		logger.Error("failed to create route", zap.Error(err))
		return err
	}

	return nil
}

// ptrTimeNow returns a pointer to the current time.
func ptrTimeNow() *time.Time {
	t := time.Now()
	return &t
}
