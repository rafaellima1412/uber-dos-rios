// ListTripUseCase defines the interface for updating an existing trip.
package input

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/riolivre/nautical_logistics/internal/application/domain"
)

type ListTripConfigInput struct {
	Page    int
	PerPage int
}
type ListTripConfigOutput struct {
	Trips       []TripConfigSummary
	TotalCount  int
	TotalPages  int
	CurrentPage int
}
type TripConfigSummary struct {
	ID             uuid.UUID
	ShipID         uuid.UUID
	RouteID        uuid.UUID
	ShipName       string
	RouteName      string
	Recurrence     domain.TripRecurrence
	ExpirationDate time.Time
	DepartureTime  string
	ArrivalTime    string
	DurationDays   int
	StartDate      time.Time
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
type ListTripConfigUseCase interface {
	Execute(ctx context.Context, input ListTripConfigInput) (*ListTripConfigOutput, error)
}
