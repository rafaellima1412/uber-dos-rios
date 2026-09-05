package input

import "context"

type GetScheduleOutput struct {
	ID       string
	Route    GetRouteOutput
	Terminal GetTerminalOutput
	Value    float64
	Active   bool
}

// GetScheduleUseCase defines the interface for retrieving a schedule by its ID.
type GetScheduleUseCase interface {
	Execute(ctx context.Context, id string) (*GetScheduleOutput, error)
}
