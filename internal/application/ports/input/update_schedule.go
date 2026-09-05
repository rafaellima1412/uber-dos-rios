package input

import "context"

type UpdateScheduleInput struct {
	ID         string
	RouteID    string
	TerminalID string
	Value      float64
	Active     bool
}

// UpdateScheduleUseCase defines the interface for updating an existing schedule.
type UpdateScheduleUseCase interface {
	Execute(ctx context.Context, input UpdateScheduleInput) error
}
