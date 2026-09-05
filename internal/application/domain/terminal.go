package domain

import (
	"time"

	"github.com/google/uuid"
)

// Terminal represents a terminal entity in the system.
type Terminal struct {
	ID        uuid.UUID
	Name      string
	UF        string
	CityID    int
	Activity  bool
	Latitude  float64
	Longitude float64
	CreatedAt time.Time
	UpdatedAt time.Time
	CityName  string
}
