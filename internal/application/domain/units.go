package domain

import "github.com/google/uuid"

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

type UnitsAvailable struct {
	ShipID   uuid.UUID
	ShipName string
	Seats    Seats
	Cabins   Cabins
}
