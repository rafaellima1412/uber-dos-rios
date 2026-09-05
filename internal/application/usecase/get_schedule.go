package usecase

import (
	"context"

	"github.com/riolivre/nautical_logistics/internal/application/ports/input"
	"github.com/riolivre/nautical_logistics/internal/application/ports/output"
)

type getScheduleUseCase struct {
	scheduleRepo output.ScheduleRepository
	terminalRepo output.TerminalRepository
	routeRepo    output.RouteRepository
}

func NewGetScheduleUseCase(
	scheduleRepo output.ScheduleRepository,
	terminalRepo output.TerminalRepository,
	routeRepo output.RouteRepository,
) input.GetScheduleUseCase {
	return &getScheduleUseCase{
		scheduleRepo: scheduleRepo,
		terminalRepo: terminalRepo,
		routeRepo:    routeRepo,
	}
}

// Execute implements input.GetScheduleUseCase.
func (g *getScheduleUseCase) Execute(ctx context.Context, id string) (*input.GetScheduleOutput, error) {
	schedule, err := g.scheduleRepo.GetSchedule(ctx, id)
	if err != nil {
		return nil, err
	}

	terminal, err := g.terminalRepo.GetTerminal(ctx, schedule.TerminalID)
	if err != nil {
		return nil, err
	}

	route, err := g.routeRepo.GetRoute(ctx, schedule.RouteID)
	if err != nil {
		return nil, err
	}

	return &input.GetScheduleOutput{
		ID: schedule.ID.String(),
		Route: input.GetRouteOutput{
			ID:     route.ID.String(),
			Name:   route.Name,
			Active: route.Active,
		},
		Terminal: input.GetTerminalOutput{
			ID:        terminal.ID,
			Name:      terminal.Name,
			UF:        terminal.UF,
			CityID:    terminal.CityID,
			Latitude:  terminal.Latitude,
			Longitude: terminal.Longitude,
			Active:    terminal.Activity,
		},
		Active: schedule.Active,
		Value:  schedule.Value,
	}, nil
}
