package input

import (
	"context"

	"github.com/google/uuid"
)

type UpdateShipOutput struct {
	ID       uuid.UUID
	TypeShip string
	Status   string
	Name     string
}

type UpdateShipInput struct {
	TypeShip          string
	Status            string
	Name              string
	IMO               string
	TotalSeats        int
	TotalCabins       int
	PassengerCapacity int
	WeightCapacity    float64
	ConfigurationsID  string
	OrganizationID    string
	ImageUrl          *[]string
}

type UpdateShipUseCase interface {
	Execute(ctx context.Context, input *UpdateShipInput, id uuid.UUID) error
}
