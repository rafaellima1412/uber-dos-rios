package usecase

import (
	"context"

	"github.com/rafaellima1412/uber-dos-rios/internal/application/ports/input"
	"github.com/rafaellima1412/uber-dos-rios/internal/application/ports/output"
)

type listScheduleUseCase struct {
	scheduleRepository output.ScheduleRepository
	routeRepo          output.RouteRepository
	terminalRepo       output.TerminalRepository
}

func NewListScheduleUseCase(
	scheduleRepo output.ScheduleRepository,
	routeRepo output.RouteRepository,
	terminalRepo output.TerminalRepository,
) input.ListScheduleUseCase {
	return &listScheduleUseCase{
		scheduleRepository: scheduleRepo,
		routeRepo:          routeRepo,
		terminalRepo:       terminalRepo,
	}
}

// Execute implements input.ListScheduleUseCase.
func (l *listScheduleUseCase) Execute(ctx context.Context, inputList input.ListScheduleInput) (*input.ListScheduleOutput, error) {
	schedules, err := l.scheduleRepository.ListSchedules(ctx, inputList.PerPage, (inputList.Page-1)*inputList.PerPage)
	if err != nil {
		return nil, err
	}

	scheduleSummaries := make([]input.ScheduleSummary, len(schedules))
	for i, schedule := range schedules {
		route, err := l.routeRepo.GetRoute(ctx, schedule.RouteID)
		if err != nil {
			return nil, err
		}

		terminal, err := l.terminalRepo.GetTerminal(ctx, schedule.TerminalID)
		if err != nil {
			return nil, err
		}

		scheduleSummaries[i] = input.ScheduleSummary{
			ID:         schedule.ID.String(),
			RouteID:    route.ID.String(),
			TerminalID: terminal.ID.String(),
			TenimalName: terminal.Name,
			RouteName: route.Name,
			Value:      schedule.Value,
			Active:     schedule.Active,
		}
	}

	output := &input.ListScheduleOutput{
		Schedules:  scheduleSummaries,
		TotalCount: len(scheduleSummaries), // This should ideally come from a count query
		TotalPages: 1,                      // This should be calculated based on total count and per page
	}

	return output, nil
}
