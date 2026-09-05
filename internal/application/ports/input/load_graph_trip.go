package input

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type Edge struct {
	To                  int
	Cost                int
	OrganizationID      uuid.UUID
	RouteID             uuid.UUID
	RouteName           string
	DateDepartureTime   time.Time
	DateArrivalTime     time.Time
	ShipID              uuid.UUID
	ShipName            string
	ShipImageURL        string
	TripConfigurationID uuid.UUID
}

type Graph map[int][]*Edge

type LoadGraphTripInput interface {
	Execute(ctx context.Context, DateOrigin, DateDest string) (Graph, error)
}
