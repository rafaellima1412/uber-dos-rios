package input

import (
	"time"

	"github.com/google/uuid"
)

type TripSearchResponse struct {
	TotalCost int
	Path      []PathStep
}

type PathStep struct {
	OrganizationID    uuid.UUID
	CityID            int
	Cost              int
	RouteID           uuid.UUID
	DateDepartureTime time.Time
	DateArrivalTime   time.Time
	RouteName         string
}

type DijkstraTripInput interface {
	Execute(graph Graph, origin, destination int) (TripSearchResponse, bool)
}
