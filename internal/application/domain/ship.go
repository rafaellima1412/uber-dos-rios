package domain

import (
	"time"

	"github.com/google/uuid"
)

type Ship struct {
	ID                uuid.UUID
	OrganizationID    *uuid.UUID
	ConfigurationID   *[]uuid.UUID
	Name              string
	IMO               string // International Maritime Organization number (registro)
	TypeShip          ShipType
	StatusShip        ShipStatus
	PassengerCapacity int
	WeightCapacity    float64
	TotalSeats        int
	TotalCabins       int
	AvailableSeats    int
	CreatedAt         time.Time
	UpdatedAt         *time.Time
	ImageUrl          []string
}
