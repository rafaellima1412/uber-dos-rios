package usecase

import (
	"context"

	"github.com/google/uuid"
	"github.com/riolivre/nautical_logistics/internal/application/ports/input"
	"github.com/riolivre/nautical_logistics/internal/application/ports/output"
)

type getRouteUseCase struct {
	routeRepo           output.RouteRepository
	organizationService output.OrganizationService
}

// Ensure getRouteUseCase implements input.GetRouteUseCase interface.
var _ input.GetRouteUseCase = (*getRouteUseCase)(nil)

// NewGetRouteUseCase creates a new instance of GetRouteUseCase.
func NewGetRouteUseCase(routeRepo output.RouteRepository, organizationService output.OrganizationService) input.GetRouteUseCase {
	return &getRouteUseCase{
		routeRepo:           routeRepo,
		organizationService: organizationService,
	}
}

// Execute implements input.GetRouteUseCase.
func (uc *getRouteUseCase) Execute(ctx context.Context, inputID string) (input.GetRouteOutput, error) {
	routeID, err := uuid.Parse(inputID)
	if err != nil {
		return input.GetRouteOutput{}, err
	}

	route, err := uc.routeRepo.GetRoute(ctx, routeID)
	if err != nil {
		return input.GetRouteOutput{}, err
	}

	org, err := uc.organizationService.GetOrganization(ctx, route.OrganizationID.String())
	if err != nil {
		return input.GetRouteOutput{}, err
	}

	return input.GetRouteOutput{
		ID:     route.ID.String(),
		Name:   route.Name,
		Active: route.Active,
		OrganizationOutput: input.GetRouteOrganizationOutput{
			ID:            org.ID.String(),
			Type:          org.Type,
			Name:          org.Name,
			CNPJ:          org.CNPJ,
			Email:         org.Email,
			PhoneNumber:   org.PhoneNumber,
			Address:       org.Address,
			OwnerInfo:     org.OwnerInfo,
			CustomLogoURL: org.CustomLogoURL,
		},
	}, nil
}
