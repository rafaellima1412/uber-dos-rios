package output

import (
	"context"

	"github.com/google/uuid"
)


type UsersDTO struct{
	FullName       string     `json:"full_name" validate:"required"`
	Email          string     `json:"email" validate:"required,email"`
	CPF            string     `json:"cpf" validate:"required,len=11"`
	Password       string     `json:"password" validate:"required,min=8"`
	OrganizationID *uuid.UUID `json:"organization_id,omitempty" validate:"omitempty,uuid|null"`
	Role           string     `json:"role" validate:"required,oneof=admin member viewer"`
	ContactPhone   string     `json:"contact_phone,omitempty"`
	PhotoURL       string     `json:"photo_url,omitempty"`
}

type UserService interface {
	GetUsers(ctx context.Context, id string) (*UsersDTO, error)
}	