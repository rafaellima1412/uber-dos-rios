package usecase

import (
	"context"

	"github.com/rafaellima1412/uber-dos-rios/internal/application/ports/input"
	"github.com/rafaellima1412/uber-dos-rios/internal/application/ports/output"
)

type searchScheduleUseCase struct {
	scheduleRepo output.ScheduleRepository
	routeRepo    output.RouteRepository
	terminalRepo output.TerminalRepository
}

func NewSearchScheduleUseCase(
	scheduleRepo output.ScheduleRepository,
	routeRepo output.RouteRepository,
	terminalRepo output.TerminalRepository,
) input.SearchScheduleUseCase {
	return &searchScheduleUseCase{
		scheduleRepo: scheduleRepo,
		routeRepo:    routeRepo,
		terminalRepo: terminalRepo,
	}
}

// Execute implements input.SearchScheduleUseCase.
func (s *searchScheduleUseCase) Execute(ctx context.Context, query string, inputSearch input.SearchScheduleInput) (*input.SearchScheduleOutput, error) {
	limit := inputSearch.PerPage
	offset := (inputSearch.Page - 1) * inputSearch.PerPage

	schedules, err := s.scheduleRepo.SearchSchedules(ctx, query, limit, offset)
	if err != nil {
		return nil, err
	}

	scheduleSummaries := make([]input.ScheduleSummary, len(schedules))
	for i, schedule := range schedules {
		route, err := s.routeRepo.GetRoute(ctx, schedule.RouteID)
		if err != nil {
			return nil, err
		}

		terminal, err := s.terminalRepo.GetTerminal(ctx, schedule.TerminalID)
		if err != nil {
			return nil, err
		}

		scheduleSummaries[i] = input.ScheduleSummary{
			ID:         schedule.ID.String(),
			RouteID:    route.ID.String(),
			TerminalID: terminal.ID.String(),
			Value:      schedule.Value,
			Active:     schedule.Active,
		}
	}

	output := &input.SearchScheduleOutput{
		Schedules:  scheduleSummaries,
		TotalCount: len(scheduleSummaries), // This should ideally come from a count query
		TotalPages: 1,                      // This should be calculated based on total count and per page
	}

	return output, nil
}
