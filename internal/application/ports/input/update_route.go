package input

import "context"

type UpdateRouteInput struct {
	ID             string
	Name           string
	OrganizationID string
}

type UpdateRouteOutput struct {
	ID   string
	Name string
}

type UpdateRouteUseCase interface {
	Execute(ctx context.Context, id string, inputRoute UpdateRouteInput) error
}
