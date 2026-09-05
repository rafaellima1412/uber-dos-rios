package domain

import (
	"time"

	"github.com/google/uuid"
)

type GPS struct {
	ID        uuid.UUID
	ShipName  string
	Latitude  float64
	Longitude float64
	Speed     float64
	Course    float64
	ReceivedAt time.Time
}
