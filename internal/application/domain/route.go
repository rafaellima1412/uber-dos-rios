package domain

import (
	"time"

	"github.com/google/uuid"
)

// Route represents a nautical route in the system.
type Route struct {
	ID                 uuid.UUID
	Name               string
	Active             bool
	OrganizationOutput GetRouteOrganizationOutput
	OrganizationID     uuid.UUID
	CreatedAt          *time.Time
	UpdatedAt          *time.Time
}

// GetRouteOrganizationOutput represents the organization details in the GetRoute output.
type GetRouteOrganizationOutput struct {
	ID   uuid.UUID
	Type string
	Name string
	CNPJ *string
	Email *string
	PhoneNumber *string
}
