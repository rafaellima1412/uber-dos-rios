package input

import (
	"context"

	"github.com/google/uuid"
)

type Seats struct {
	TotalSeats     int
	OccupiedUnits  []string
	AvailableSeats int
}
type Cabins struct {
	TotalCabins     int
	OccupiedUnits   []string
	AvailableCabins int
}

type ListAvailableOutput struct {
	ShipID   uuid.UUID
	ShipName string
	Seats    Seats
	Cabins   Cabins
}

type ListAvailableUseCase interface {
	Execute(ctx context.Context, id uuid.UUID, dateOrigin, dateDest string) (*ListAvailableOutput, error)
}
