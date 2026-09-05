package input

import (
	"context"
	"time"

	"github.com/google/uuid"
)
type SearchTripInput struct {
	Page    int
	PerPage int
}

type TripFilterInput struct {
	RouteID              *string
	ShipID               *string
	DepartureAfter       *time.Time
	DepartureBefore      *time.Time
	TripConfigurationsID *string
}

type TripFilterOutput struct {
	Trips       []TripSummary
	TotalCount  int
	TotalPages  int
	CurrentPage int
}

type TripSummary struct {
	ID                  uuid.UUID
	ShipID              uuid.UUID
	RouteID             uuid.UUID
	TripConfigurationID uuid.UUID
	DepartureAt         time.Time
	ArrivalAt           time.Time
	CreatedAt           time.Time
	UpdatedAt           *time.Time
}

type TripSearchUseCase interface {
	Execute(ctx context.Context, input *TripFilterInput, inputSearch SearchTripInput) (TripFilterOutput, error)
}
