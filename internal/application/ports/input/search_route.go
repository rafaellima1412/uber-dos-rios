package input

import "context"

type SearchRouteInput struct {
	Page    int
	PerPage int
}

type SearchRouteSummary struct {
	ID                 string
	Name               string
	Active             bool
	OrganizationOutput GetRouteOrganizationOutput
}

type SearchRouteOutput struct {
	Routes      []SearchRouteSummary
	TotalCount  int
	TotalPages  int
	CurrentPage int
}

type SearchRouteUseCase interface {
	Execute(ctx context.Context, query string, inputSearch SearchRouteInput) (*SearchRouteOutput, error)
}
