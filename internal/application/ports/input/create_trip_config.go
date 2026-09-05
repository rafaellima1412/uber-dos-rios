// CreateTripInputPort defines the input port for creating a trip
package input

import (
	"context"
	"time"

	"github.com/google/uuid"
)

// CreateTripInput represents the input data required to create a new trip.
type CreateTripConfigInput struct {
	ShipID         string
	RouteID        string
	Recurrence     string
	ExpirationDate time.Time // TIMESTAMPTZ
	DepartureTime  string    // "HH:MM:SS"
	ArrivalTime    string    // "HH:MM:SS"
	DurationDays   int
	StartDate      time.Time // TIMESTAMPTZ
}	

// CreateTripOutput represents the output data after creating a new trip.
type CreateTripConfigOutput struct {
	ID uuid.UUID
	ShipID         uuid.UUID
	RouteID        uuid.UUID
	Recurrence     string
	ExpirationDate time.Time
	DepartureTime  time.Time
	DurationDays   int
	StartDate      time.Time
}

// CreateTripUseCase defines the interface for creating a new trip.
type CreateTripConfigUseCase interface {
	Execute(ctx context.Context, inputTrip *CreateTripConfigInput) error
}