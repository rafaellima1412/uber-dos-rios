package usecase

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/rafaellima1412/uber-dos-rios/internal/application/ports/input"
	"github.com/rafaellima1412/uber-dos-rios/internal/application/ports/output"
)

type updateScheduleUseCase struct {
	scheduleRepo output.ScheduleRepository
	terminalRepo output.TerminalRepository
	routeRepo    output.RouteRepository
}

func NewUpdateScheduleUseCase(
	scheduleRepo output.ScheduleRepository,
	terminalRepo output.TerminalRepository,
	routeRepo output.RouteRepository,
) input.UpdateScheduleUseCase {
	return &updateScheduleUseCase{
		scheduleRepo: scheduleRepo,
		terminalRepo: terminalRepo,
		routeRepo:    routeRepo,
	}
}

// Execute implements input.UpdateScheduleUseCase.
func (u *updateScheduleUseCase) Execute(ctx context.Context, inputSchedule input.UpdateScheduleInput) error {
	terminalUUID, err := uuid.Parse(inputSchedule.TerminalID)
	if err != nil {
		return err
	}
	terminal, err := u.terminalRepo.GetTerminal(ctx, terminalUUID)
	if err != nil {
		return err
	}

	routeUUID, err := uuid.Parse(inputSchedule.RouteID)
	if err != nil {
		return err
	}
	route, err := u.routeRepo.GetRoute(ctx, routeUUID)
	if err != nil {
		return err
	}

	scheduleUUID, err := uuid.Parse(inputSchedule.ID)
	if err != nil {
		return err
	}

	schedule, err := u.scheduleRepo.GetSchedule(ctx, scheduleUUID.String())
	if err != nil {
		return err
	}

	schedule.TerminalID = terminal.ID
	schedule.RouteID = route.ID
	schedule.Active = inputSchedule.Active
	schedule.Value = inputSchedule.Value
	schedule.UpdatedAt = time.Now()

	return u.scheduleRepo.UpdateSchedule(ctx, schedule)
}
