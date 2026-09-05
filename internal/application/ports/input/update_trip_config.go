// UpdateTripUseCase defines the interface for updating an existing trip.
package input

import (
	"context"
	"time"
)	

type UpdateTripConfigInput struct {
	TripID         string
	ShipID         string
	RouteID        string
	Recurrence     string
	StartDate    time.Time
	ExpirationDate time.Time	
	DepartureTime  time.Time
	DurationDays   int
}	

type UpdateTripConfigUseCase interface {
	Execute(ctx context.Context, id string, inputDTO *UpdateTripConfigInput) error
}