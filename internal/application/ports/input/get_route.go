package input

import (
	"context"
)

// GetRouteOrganizationOutput represents the organization details in the GetRoute output.
type GetRouteOrganizationOutput struct {
	ID            string
	Type          string
	Name          string
	CNPJ          *string
	Email         *string
	PhoneNumber   *string
	Address       map[string]any
	OwnerInfo     map[string]any
	CustomLogoURL *string
}

// GetRouteOutput represents the output data for the GetRoute use case.
type GetRouteOutput struct {
	ID                 string
	Name               string
	Active             bool
	OrganizationOutput GetRouteOrganizationOutput
}

// GetRouteUseCase defines the interface for the GetRoute use case.
type GetRouteUseCase interface {
	Execute(ctx context.Context, inputID string) (GetRouteOutput, error)
}
