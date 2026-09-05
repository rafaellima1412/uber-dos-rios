package input

import "context"

type ListScheduleInput struct {
	Page    int
	PerPage int
}

type ScheduleSummary struct {
	ID         string
	RouteID    string
	TerminalID string
	TenimalName string
	RouteName string
	Value      float64
	Active     bool
}

type ListScheduleOutput struct {
	Schedules   []ScheduleSummary
	TotalCount  int
	TotalPages  int
	CurrentPage int
}

type ListScheduleUseCase interface {
	Execute(ctx context.Context, input ListScheduleInput) (*ListScheduleOutput, error)
}
