package output

import (
	"context"

	"github.com/google/uuid"
)

type OrganizationDTO struct {
	ID            uuid.UUID      `json:"id"`
	Type          string         `json:"type"`
	Name          string         `json:"name"`
	CNPJ          *string        `json:"cnpj,omitempty"`
	Email         *string        `json:"email,omitempty"`
	PhoneNumber   *string        `json:"phone_number,omitempty"`
	Address       map[string]any `json:"address,omitempty"` //TODO SOLID liskov ser uma struct real e não generico
	OwnerInfo     map[string]any `json:"owner_info,omitempty"`
	CustomLogoURL *string        `json:"custom_logo_url,omitempty" validate:"omitempty,url"`
}

type OrganizationService interface {
	GetOrganization(ctx context.Context, id string) (*OrganizationDTO, error)
}
