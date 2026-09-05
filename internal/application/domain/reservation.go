package domain

import (
	"time"

	"github.com/google/uuid"
)

type Reservation struct {
	ID     uuid.UUID
	UserID uuid.UUID
	ReservationDate *time.Time
	Status          ReservationStatus
	CreatedAt       time.Time
	UpdatedAt       *time.Time
}

