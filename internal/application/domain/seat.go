package domain

import (
	"time"

	"github.com/google/uuid"
)

type Seat struct {
	ID                uuid.UUID
	IsActive          bool
	LeftColumnsCount  int
	RightColumnsCount int
	Prefix            bool
	CreatedAt         time.Time
	UpdatedAt         *time.Time
}
