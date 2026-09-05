// GetTripUseCase defines the interface for updating an existing trip.
package input

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type GetTripConfigOutput struct {
	ID             uuid.UUID
	ShipID         uuid.UUID
	RouteID        uuid.UUID
	ShipName       string
	RouteName      string
	Recurrence     string
	ExpirationDate time.Time
	DepartureTime  string
	StartDate      string
	DurationDays   int
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type GetTripConfigUseCase interface {
	Execute(ctx context.Context, id string) (*GetTripConfigOutput, error)
}
