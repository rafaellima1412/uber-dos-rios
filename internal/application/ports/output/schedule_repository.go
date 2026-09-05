package output

import (
	"context"

	"github.com/rafaellima1412/uber-dos-rios/internal/application/domain"
)

type ScheduleRepository interface {
	CreateSchedule(ctx context.Context, schedule *domain.Schedule) error
	GetSchedule(ctx context.Context, id string) (*domain.Schedule, error)
	UpdateSchedule(ctx context.Context, schedule *domain.Schedule) error
	DeleteSchedule(ctx context.Context, id string) error
	ListSchedules(ctx context.Context, limit int, offset int) ([]*domain.Schedule, error)
	SearchSchedules(ctx context.Context, query string, limit int, offset int) ([]*domain.Schedule, error)
}
