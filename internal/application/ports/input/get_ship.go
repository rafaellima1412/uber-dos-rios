package input

import (
	"context"

	"github.com/google/uuid"
)

type GetShipOutput struct {
	ID                uuid.UUID
	ConfigurationID   *[]uuid.UUID
	OrganizationID    *uuid.UUID
	IMO               string
	Name              string
	Status            string
	TypeShip          string
	TotalSeats        int
	PassengerCapacity int
	WeightCapacity    float64
	ImageUrl          []string
}

type GetShipUseCase interface {
	Execute(ctx context.Context, id uuid.UUID, typeShip string) (*GetShipOutput, error)
}
