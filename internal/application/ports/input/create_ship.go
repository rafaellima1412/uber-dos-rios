package input

import (
	"context"

	"github.com/google/uuid"
)



type CreateShipInput struct {
	Name              string
	TypeShip          string
	IMO               string
	TotalSeats        int
	TotalCabins       int
	Status            string
	PassengerCapacity int
	WeightCapacity    float64
	CreatedAt         string
	OrganizationID    string
	ConfigurationsID  []uuid.UUID
}

type CreateShipOutput struct {
	ID       uuid.UUID
	TypeShip string
	Status   string
	Name     string
}

type CreateShipUseCase interface {
	Execute(ctx context.Context, input *CreateShipInput) error
}
