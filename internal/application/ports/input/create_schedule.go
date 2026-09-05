package input

import (
	"context"
	"time"
)

type CreateScheduleInput struct {
	RouteID    string
	TerminalID string
	Value      float64
	Active     bool
	StopOrder  int
	CreatedAt  time.Time
}

// CreateScheduleUseCase defines the interface for creating a new schedule.
type CreateScheduleUseCase interface {
	Execute(ctx context.Context, input CreateScheduleInput) error
}
