package usecase

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/riolivre/nautical_logistics/internal/application/domain"
	"github.com/riolivre/nautical_logistics/internal/application/ports/input"
	"github.com/riolivre/nautical_logistics/internal/application/ports/output"
)

type createScheduleUseCase struct {
	scheduleRepo output.ScheduleRepository
	terminalRepo output.TerminalRepository
	routeRepo    output.RouteRepository
}

func NewCreateScheduleUseCase(scheduleRepo output.ScheduleRepository, terminalRepo output.TerminalRepository, routeRepo output.RouteRepository) input.CreateScheduleUseCase {
	return &createScheduleUseCase{
		scheduleRepo: scheduleRepo,
		terminalRepo: terminalRepo,
		routeRepo:    routeRepo,
	}
}

// Execute implements input.CreateScheduleUseCase.
func (uc *createScheduleUseCase) Execute(ctx context.Context, inputSchedule input.CreateScheduleInput) error {
	scheduleID, err := uuid.NewUUID()
	if err != nil {
		return err
	}

	terminalUUID, err := uuid.Parse(inputSchedule.TerminalID)
	if err != nil {
		return err
	}

	terminal, err := uc.terminalRepo.GetTerminal(ctx, terminalUUID)
	if err != nil {
		return err
	}
	if terminal == nil {
		return err
	}

	routeUUID, err := uuid.Parse(inputSchedule.RouteID)
	if err != nil {
		return err
	}

	route, err := uc.routeRepo.GetRoute(ctx, routeUUID)
	if err != nil {
		return err
	}
	if route == nil {
		return err
	}

	schedule := &domain.Schedule{
		ID:         scheduleID,
		TerminalID: terminal.ID,
		RouteID:    route.ID,
		Value:      inputSchedule.Value,
		Active:     true,
		StopOrder:  inputSchedule.StopOrder,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	if err := uc.scheduleRepo.CreateSchedule(ctx, schedule); err != nil {
		return err
	}

	return nil
}
