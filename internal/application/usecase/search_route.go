package usecase

import (
	"context"

	"github.com/rafaellima1412/uber-dos-rios/internal/application/ports/input"
	"github.com/rafaellima1412/uber-dos-rios/internal/application/ports/output"
)

type searchRouteUseCase struct {
	routeRepo           output.RouteRepository
	organizationService output.OrganizationService
}

func NewSearchRouteUseCase(routeRepo output.RouteRepository, organizationService output.OrganizationService) *searchRouteUseCase {
	return &searchRouteUseCase{
		routeRepo:           routeRepo,
		organizationService: organizationService,
	}
}

func (uc *searchRouteUseCase) Execute(ctx context.Context, query string, inputSearch input.SearchRouteInput) (*input.SearchRouteOutput, error) {
	limit := inputSearch.PerPage
	offset := (inputSearch.Page - 1) * inputSearch.PerPage

	routes, err := uc.routeRepo.SearchRoutes(ctx, query, limit, offset)
	if err != nil {
		return nil, err
	}

	routeSummaries := make([]input.SearchRouteSummary, len(routes))
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

		routeSummaries[i] = input.SearchRouteSummary{
			ID:                 route.ID.String(),
			Name:               route.Name,
			Active:             route.Active,
			OrganizationOutput: orgOutput,
		}
	}

	output := &input.SearchRouteOutput{
		Routes:      routeSummaries,
		TotalCount:  len(routeSummaries), // This should ideally come from a count query
		TotalPages:  1,                   // This should be calculated based on total count and per page
		CurrentPage: inputSearch.Page,
	}

	return output, nil
}
