package usecase

import (
	"context"

	"github.com/riolivre/nautical_logistics/internal/application/ports/input"
	"github.com/riolivre/nautical_logistics/internal/application/ports/output"
)

type deleteScheduleUseCase struct {
	scheduleRepository output.ScheduleRepository
}

func NewDeleteScheduleUseCase(scheduleRepo output.ScheduleRepository) input.DeleteScheduleUseCase {
	return &deleteScheduleUseCase{
		scheduleRepository: scheduleRepo,
	}
}

// Execute implements input.DeleteScheduleUseCase.
func (uc *deleteScheduleUseCase) Execute(ctx context.Context, id string) error {
	return uc.scheduleRepository.DeleteSchedule(ctx, id)
}
