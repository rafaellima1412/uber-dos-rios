// GetReservationUseCase defines 
package input

import "context"

type GetReservationOutput struct {
	ID              string
	UserID          string
	TripID          []string
	OccupiedSeats   []string
	ReservationDate string
	Status          string
}

type GetReservationUseCase interface {
	Execute(ctx context.Context, id string) (*GetReservationOutput, error)
}



