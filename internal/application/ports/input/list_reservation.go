package input

import (
	"context"
	"time"

	"github.com/rafaellima1412/uber-dos-rios/internal/application/domain"
)

type ListReservationInput struct {
	Page    int
	PerPage int
}

type ReservationTripConfigurationSummary struct {
	ID                  string
	TripConfigurationID string
	OccupiedSeats       []string
}
type ReservationSummary struct {
	ID              string
	UserID          string
	UserFullName    string
	RouteID         string
	RouteName       string
	ShipID          string
	ShipName        string
	ReservationName string
	ReservationDate *time.Time
	Status          domain.ReservationStatus
	CreatedAt       time.Time
	UpdatedAt       *time.Time

	TripConfigurations []ReservationTripConfigurationSummary
}

type ListReservationOutput struct {
	Reservations []ReservationSummary
	TotalCount   int
	TotalPages   int
	CurrentPage  int
}

type ListReservationUseCase interface {
	Execute(ctx context.Context, input ListReservationInput) (*ListReservationOutput, error)
}
