package input

import "context"

// ListRouteInput is the input data transfer object for listing routes.
type ListRouteInput struct {
	Page    int
	PerPage int
}

// RouteSummary is a summary representation of a route.
type RouteSummary struct {
	ID                 string
	Name               string
	Active             bool
	OrganizationOutput GetRouteOrganizationOutput
}

// ListRouteOutput is the output data transfer object for listing routes.
type ListRouteOutput struct {
	Routes      []RouteSummary
	TotalCount  int
	TotalPages  int
	CurrentPage int
}

// ListRouteUseCase defines the interface for the use case of listing routes.
type ListRouteUseCase interface {
	Execute(ctx context.Context, inputListRoute ListRouteInput) (*ListRouteOutput, error)
}
