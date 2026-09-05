package usecase

import (
	"context"

	"github.com/rafaellima1412/uber-dos-rios/internal/application/ports/input"
	"github.com/rafaellima1412/uber-dos-rios/internal/application/ports/output"
)

type listRouteUseCase struct {
	routeRepo           output.RouteRepository
	organizationService output.OrganizationService
}

// NewListRouteUseCase creates a new instance of ListRouteUseCase.
func NewListRouteUseCase(routeRepo output.RouteRepository, organizationService output.OrganizationService) input.ListRouteUseCase {
	return &listRouteUseCase{
		routeRepo:           routeRepo,
		organizationService: organizationService,
	}
}

// Execute lists routes based on the provided input parameters.
func (uc *listRouteUseCase) Execute(ctx context.Context, inputListRoute input.ListRouteInput) (*input.ListRouteOutput, error) {
	routes, err := uc.routeRepo.ListRoutes(ctx, inputListRoute.PerPage, (inputListRoute.Page-1)*inputListRoute.PerPage)
	if err != nil {
		return nil, err
	}

	routeSummaries := make([]input.RouteSummary, len(routes))
	for i, route := range routes {
		org, err := uc.organizationService.GetOrganization(ctx, route.OrganizationID.String())
		if err != nil {
			return nil, err
		}

		orgOutput := input.GetRouteOrganizationOutput{
			ID:            org.ID.String(),
			Type:          org.Type,
			Name:          org.Name,
			CNPJ:          org.CNPJ,
			Email:         org.Email,
			PhoneNumber:   org.PhoneNumber,
			Address:       org.Address,
			OwnerInfo:     org.OwnerInfo,
			CustomLogoURL: org.CustomLogoURL,
		}

		routeSummaries[i] = input.RouteSummary{
			ID:                 route.ID.String(),
			Name:               route.Name,
			Active:             route.Active,
			OrganizationOutput: orgOutput,
		}
	}

	output := &input.ListRouteOutput{
		Routes:      routeSummaries,
		TotalCount:  len(routeSummaries), // This should ideally come from a count query
		TotalPages:  1,                   // This should be calculated based on total count and per page
		CurrentPage: inputListRoute.Page,
	}

	return output, nil
}
