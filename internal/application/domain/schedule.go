package domain

import (
	"time"

	"github.com/google/uuid"
)

type Schedule struct {
	ID         uuid.UUID
	RouteID    uuid.UUID
	TerminalID uuid.UUID
	Value      float64
	Active     bool
	StopOrder  int
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
