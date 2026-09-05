package input

import (
	"time"

	"github.com/google/uuid"
)

type TripSearchInput struct {
	DateDepartureTime    time.Time
	DateArrivalTime      time.Time
	CityOriginID         int
	CityDestID           int
	MaxStops             int
	MaxCost              int
	ShipID               string
	PassangerCountSeats  int
	PassangerCountCabins int
	MaxResults           int
}

type Path struct {
	OrganizationID      uuid.UUID
	CityID              int
	Cost                int
	RouteID             uuid.UUID
	RouteName           string
	DateDepartureTime   time.Time
	DateArrivalTime     time.Time
	ShipID              uuid.UUID
	ShipName            string
	ShipURL             string
	TripConfigurationID uuid.UUID
}

type TripSearchOutput struct {
	TotalCost int
	Path      []Path
}

type DFSTripsUseCase interface {
	Execute(
		graph Graph,
		input TripSearchInput,
	) ([]TripSearchOutput, error)
}
