package input

import (
	"context"

	"github.com/google/uuid"
)

type ListShipInput struct {
	Page    int
	PerPage int
}

type ListShipOutput struct {
	Ships       []ShipSummary
	TotalCount  int
	TotalPages  int
	CurrentPage int
}

type ShipSummary struct {
	ID                uuid.UUID
	ConfigurationID   *[]uuid.UUID
	OrganizationID    *uuid.UUID
	Name              string
	IMO               string
	Status            string
	TypeShip          string
	PassengerCapacity int
	WeightCapacity    float64
	TotalSeats        int
	ImageUrl          []string
}

type ListShipUseCase interface {
	Execute(ctx context.Context, input ListShipInput, typeship string) (ListShipOutput, error)
}
