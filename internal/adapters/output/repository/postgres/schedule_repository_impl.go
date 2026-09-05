package postgres

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rafaellima1412/uber-dos-rios/internal/application/domain"
	"github.com/rafaellima1412/uber-dos-rios/internal/application/ports/output"
)

type scheduleRepositoryImpl struct {
	db *pgxpool.Pool
}

// Ensure scheduleRepositoryImpl implements output.ScheduleRepository.
var _ output.ScheduleRepository = (*scheduleRepositoryImpl)(nil)

// NewScheduleRepositoryImpl creates a new instance of scheduleRepositoryImpl.
func NewScheduleRepositoryImpl(db *pgxpool.Pool) output.ScheduleRepository {
	return &scheduleRepositoryImpl{
		db: db,
	}
}

// CreateSchedule implements output.ScheduleRepository.
func (s *scheduleRepositoryImpl) CreateSchedule(ctx context.Context, schedule *domain.Schedule) error {
	tx, err := s.db.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	query := `
		INSERT INTO schedules (
			id,
			route_id,
			terminal_id,
			value,
			active,
			stop_order,
			created_at
		) VALUES ($1, $2, $3, $4, $5, $6, $7)
	`
	_, err = tx.Exec(ctx, query,
		schedule.ID,
		schedule.RouteID,
		schedule.TerminalID,
		schedule.Value,
		schedule.Active,
		schedule.StopOrder,
		schedule.CreatedAt,
	)
	if err != nil {
		return err
	}

	if err := tx.Commit(ctx); err != nil {
		return err
	}

	return nil
}

// GetSchedule implements output.ScheduleRepository.
func (s *scheduleRepositoryImpl) GetSchedule(ctx context.Context, id string) (*domain.Schedule, error) {
	var schedule domain.Schedule

	query := `
		SELECT 
			id,
			route_id,
			terminal_id,
			value,
			active
		FROM schedules
		WHERE id = $1
	`
	err := s.db.QueryRow(ctx, query, id).Scan(
		&schedule.ID,
		&schedule.RouteID,
		&schedule.TerminalID,
		&schedule.Value,
		&schedule.Active,
	)
	if err != nil {
		return nil, err
	}

	return &schedule, nil
}

// UpdateSchedule implements output.ScheduleRepository.
func (s *scheduleRepositoryImpl) UpdateSchedule(ctx context.Context, schedule *domain.Schedule) error {
	query := `UPDATE schedules SET
		route_id = $1,
		terminal_id = $2,
		value = $3
	WHERE id = $4
	`
	_, err := s.db.Exec(ctx, query,
		schedule.RouteID,
		schedule.TerminalID,
		schedule.Value,
		schedule.ID,
	)
	if err != nil {
		return err
	}
	return nil
}

// DeleteSchedule implements output.ScheduleRepository.
func (s *scheduleRepositoryImpl) DeleteSchedule(ctx context.Context, id string) error {
	query := `DELETE FROM schedules WHERE id = $1`
	_, err := s.db.Exec(ctx, query, id)
	if err != nil {
		return err
	}

	return nil
}

// ListSchedules implements output.ScheduleRepository.
func (s *scheduleRepositoryImpl) ListSchedules(ctx context.Context, limit int, offset int) ([]*domain.Schedule, error) {
	query := `
		SELECT 
			id,
			route_id,
			terminal_id,
			value,
			active
		FROM schedules
		LIMIT $1 OFFSET $2
	`
	rows, err := s.db.Query(ctx, query, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var schedules []*domain.Schedule
	for rows.Next() {
		var schedule domain.Schedule
		err := rows.Scan(
			&schedule.ID,
			&schedule.RouteID,
			&schedule.TerminalID,
			&schedule.Value,
			&schedule.Active,
		)
		if err != nil {
			return nil, err
		}
		schedules = append(schedules, &schedule)
	}

	if rows.Err() != nil {
		return nil, rows.Err()
	}

	return schedules, nil
}

// SearchSchedules implements output.ScheduleRepository.
func (s *scheduleRepositoryImpl) SearchSchedules(ctx context.Context, query string, limit int, offset int) ([]*domain.Schedule, error) {
	searchQuery := `
		SELECT 
			id,
			route_id,
			terminal_id,
			value,
			active
		FROM schedules
		WHERE value::text ILIKE '%' || $1 || '%'
		LIMIT $2 OFFSET $3
	`
	rows, err := s.db.Query(ctx, searchQuery, query, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var schedules []*domain.Schedule
	for rows.Next() {
		var schedule domain.Schedule
		err := rows.Scan(
			&schedule.ID,
			&schedule.RouteID,
			&schedule.TerminalID,
			&schedule.Value,
			&schedule.Active,
		)
		if err != nil {
			return nil, err
		}
		schedules = append(schedules, &schedule)
	}

	if rows.Err() != nil {
		return nil, rows.Err()
	}

	return schedules, nil
}
