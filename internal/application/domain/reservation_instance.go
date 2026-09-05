package domain

import "github.com/google/uuid"

type TripToReservation struct {
	ReservationID  uuid.UUID
	TripInstanceID uuid.UUID
	OccupiedUnits  []string
}