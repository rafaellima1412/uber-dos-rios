package input

import (
	"context"
)

type SearchScheduleInput struct {
	Page    int
	PerPage int
}

type SearchScheduleSummary struct {
	ID         string
	RouteID    string
	TerminalID string
	Value      float64
	Active     bool
}

type SearchScheduleOutput struct {
	Schedules   []ScheduleSummary
	TotalCount  int
	TotalPages  int
	CurrentPage int
}

type SearchScheduleUseCase interface {
	Execute(ctx context.Context, query string, inputSearch SearchScheduleInput) (*SearchScheduleOutput, error)
}
