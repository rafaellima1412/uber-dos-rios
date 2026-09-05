package domain

import (
	"time"

	"github.com/google/uuid"
)

type Trip struct {
	ID                  uuid.UUID
	ShipID              uuid.UUID
	RouteID             uuid.UUID
	TripConfigurationID uuid.UUID
	DepartureAt         time.Time
	ArrivalAt           time.Time
	CreatedAt           time.Time
	UpdatedAt           *time.Time
}
type TripFilter struct {
	ID                  *string
	TripConfigurationID *string
	RouteID             *string
	ShipID              *string
	DepartureAfter      *time.Time
	DepartureBefore     *time.Time
}
