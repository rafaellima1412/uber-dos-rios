package input

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type SelectedTrips struct {
	ID                  uuid.UUID
	TripConfigurationID uuid.UUID
	ShipID              uuid.UUID
	RouteID             uuid.UUID
	DepartureAt         time.Time
	ArrivalAt           time.Time
	OccupiedUnits       []string
}
type TripToReservationInput struct {
	ReservationID uuid.UUID
	SelectedTrips []SelectedTrips
}

type AddTripsToReservationUseCase interface {
	Execute(
		ctx context.Context,
		input *TripToReservationInput,
	) error
}
