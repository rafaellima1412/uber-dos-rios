package input

import "context"

// CreateRouteInput represents the input data required to create a new route.
type CreateRouteInput struct {
	Name           string
	OrganizationID string
}

// CreateRouteOutput represents the output data after creating a new route.
type CreateRouteOutput struct {
	ID   string
	Name string
}

// CreateRouteUseCase defines the interface for creating a new route.
type CreateRouteUseCase interface {
	Execute(ctx context.Context, inputRoute CreateRouteInput) error
}
