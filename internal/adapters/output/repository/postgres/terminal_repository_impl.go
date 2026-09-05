package postgres

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/riolivre/nautical_logistics/internal/application/domain"
	"github.com/riolivre/nautical_logistics/internal/application/ports/output"
	"github.com/riolivre/nautical_logistics/internal/logger"
	"go.uber.org/zap"
)

// terminalRepositoryImpl implements the TerminalRepository interface.
type terminalRepositoryImpl struct {
	db *pgxpool.Pool
}

// Ensure terminalRepositoryImpl implements output.TerminalRepository.
var _ output.TerminalRepository = (*terminalRepositoryImpl)(nil)

// NewTerminalRepositoryImpl creates a new instance of terminalRepositoryImpl.
func NewTerminalRepositoryImpl(db *pgxpool.Pool) output.TerminalRepository {
	return &terminalRepositoryImpl{
		db: db,
	}
}

// CreateTerminal implements output.TerminalRepository.
func (t *terminalRepositoryImpl) CreateTerminal(ctx context.Context, terminal *domain.Terminal) error {
	tx, err := t.db.Begin(ctx)
	if err != nil {
		logger.Error("Error beginning transaction", zap.Error(err))
		return fmt.Errorf("error beginning transaction: %w", err)
	}
	defer tx.Rollback(ctx)

	query := `
		INSERT INTO terminals (
			id, 
			name, 
			uf,
			city_id, 
			activity,
			latitude, 
			longitude,
			created_at
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
	`
	_, err = tx.Exec(ctx, query,
		terminal.ID,
		terminal.Name,
		terminal.UF,
		terminal.CityID,
		terminal.Activity,
		terminal.Latitude,
		terminal.Longitude,
		terminal.CreatedAt,
	)
	if err != nil {
		logger.Error("Error inserting terminal", zap.Error(err))
		return fmt.Errorf("error inserting terminal: %w", err)
	}

	if err := tx.Commit(ctx); err != nil {
		logger.Error("Error committing transaction", zap.Error(err))
		return fmt.Errorf("error committing transaction: %w", err)
	}

	logger.Info("Terminal created successfully", zap.String("terminal_id", terminal.ID.String()))
	return nil
}

// DeleteTerminal implements output.TerminalRepository.
func (t *terminalRepositoryImpl) DeleteTerminal(ctx context.Context, id uuid.UUID) error {
	query := `DELETE FROM terminals WHERE id = $1`
	_, err := t.db.Exec(ctx, query, id)
	if err != nil {
		logger.Error("Error deleting terminal", zap.Error(err))
		return fmt.Errorf("error deleting terminal: %w", err)
	}

	logger.Info("Terminal deleted successfully", zap.String("terminal_id", id.String()))
	return nil
}

// GetTerminal implements output.TerminalRepository.
func (t *terminalRepositoryImpl) GetTerminal(ctx context.Context, id uuid.UUID) (*domain.Terminal, error) {
	var terminal domain.Terminal
	query := `
		SELECT 
			id, 
			name, 
			uf, 
			city_id, 
			activity, 
			latitude, 
			longitude
		FROM terminals 
		WHERE id = $1`
	err := t.db.QueryRow(ctx, query, id).Scan(
		&terminal.ID,
		&terminal.Name,
		&terminal.UF,
		&terminal.CityID,
		&terminal.Activity,
		&terminal.Latitude,
		&terminal.Longitude,
	)
	if err != nil {
		if err == pgx.ErrNoRows {
			logger.Error("Terminal not found", zap.Error(err))
			return nil, fmt.Errorf("terminal not found: %w", err)
		}
		logger.Error("Error finding terminal by ID", zap.Error(err))
		return nil, fmt.Errorf("error finding terminal by ID: %w", err)
	}

	logger.Info("Terminal retrieved successfully", zap.String("terminal_id", terminal.ID.String()))
	return &terminal, nil
}

// ListTerminals implements output.TerminalRepository.
func (t *terminalRepositoryImpl) ListTerminals(ctx context.Context, limit int, offset int) ([]*domain.Terminal, error) {
	query := `
  SELECT 
			t.id, 
			t.name, 
			t.uf, 
			t.city_id, 
			c."name", 
			t.activity, 
			t.latitude, 
			t.longitude 
		FROM terminals t
		join cities c on c.id = t.city_id 
		ORDER BY created_at DESC
		LIMIT $1 OFFSET $2
	`
	rows, err := t.db.Query(ctx, query, limit, offset)
	if err != nil {
		logger.Error("Error listing terminals", zap.Error(err))
		return nil, fmt.Errorf("error listing terminals: %w", err)
	}
	defer rows.Close()

	var terminals []*domain.Terminal
	for rows.Next() {
		var terminal domain.Terminal
		if err := rows.Scan(
			&terminal.ID,
			&terminal.Name,
			&terminal.UF,
			&terminal.CityID,
			&terminal.CityName,
			&terminal.Activity,
			&terminal.Latitude,
			&terminal.Longitude,
		); err != nil {
			logger.Error("Error scanning terminal row", zap.Error(err))
			return nil, fmt.Errorf("error scanning terminal row: %w", err)
		}
		terminals = append(terminals, &terminal)
	}

	if rows.Err() != nil {
		logger.Error("Error iterating terminal rows", zap.Error(rows.Err()))
		return nil, fmt.Errorf("error iterating terminal rows: %w", rows.Err())
	}

	logger.Info("Terminals listed successfully", zap.Int("count", len(terminals)))
	return terminals, nil
}

// SearchTerminals implements output.TerminalRepository.
func (t *terminalRepositoryImpl) SearchTerminals(ctx context.Context, search string, limit int, offset int) ([]*domain.Terminal, error) {
	sqlQuery := `
		SELECT 
			t.id, 
			t.name as terminal_name, 
			t.uf, 
			t.city_id,
			c.name as city_name,
			t.activity, 
			t.latitude, 
			t.longitude
		FROM terminals t
		join cities c on t.city_id = c.id 
		WHERE city_id = $1 
		ORDER BY created_at DESC
		LIMIT $2 OFFSET $3
	`
	rows, err := t.db.Query(ctx, sqlQuery, search, limit, offset)
	if err != nil {
		logger.Error("Error searching terminals", zap.Error(err))
		return nil, fmt.Errorf("error searching terminals: %w", err)
	}
	defer rows.Close()

	var terminals []*domain.Terminal
	for rows.Next() {
		var terminal domain.Terminal
		if err := rows.Scan(
			&terminal.ID,
			&terminal.Name,
			&terminal.UF,
			&terminal.CityID,
			&terminal.CityName,
			&terminal.Activity,
			&terminal.Latitude,
			&terminal.Longitude,
		); err != nil {
			logger.Error("Error scanning terminal row", zap.Error(err))
			return nil, fmt.Errorf("error scanning terminal row: %w", err)
		}
		terminals = append(terminals, &terminal)
	}

	if rows.Err() != nil {
		logger.Error("Error iterating terminal rows", zap.Error(rows.Err()))
		return nil, fmt.Errorf("error iterating terminal rows: %w", rows.Err())
	}

	logger.Info("Terminals searched successfully", zap.Int("count", len(terminals)))
	return terminals, nil
}

// UpdateTerminal implements output.TerminalRepository.
func (t *terminalRepositoryImpl) UpdateTerminal(ctx context.Context, terminal *domain.Terminal) error {
	query := `
		UPDATE terminals 
		SET name = $1, uf = $2, city_id = $3, latitude = $4, longitude = $5, updated_at = $6
		WHERE id = $7
	`
	_, err := t.db.Exec(ctx, query,
		terminal.Name,
		terminal.UF,
		terminal.CityID,
		terminal.Latitude,
		terminal.Longitude,
		terminal.UpdatedAt,
		terminal.ID,
	)
	if err != nil {
		logger.Error("Error updating terminal", zap.Error(err))
		return fmt.Errorf("error updating terminal: %w", err)
	}

	logger.Info("Terminal updated successfully", zap.String("terminal_id", terminal.ID.String()))
	return nil
}

// ListCities implements output.TerminalRepository.
func (t *terminalRepositoryImpl) ListCities(ctx context.Context) ([]*domain.City, error) {
	query := `
		select 
			c.id,
			c.name,
			c.state 
		from cities c `
	rows, err := t.db.Query(ctx, query)
	if err != nil {
		logger.Error("Error listing cities", zap.Error(err))
		return nil, fmt.Errorf("error listing cities: %w", err)
	}
	defer rows.Close()

	var cities []*domain.City
	for rows.Next() {
		var city domain.City
		if err := rows.Scan(
			&city.ID,
			&city.Name,
			&city.State,
		); err != nil {
			logger.Error("Error scanning city row", zap.Error(err))
			return nil, fmt.Errorf("error scanning city row: %w", err)
		}
		cities = append(cities, &city)
	}

	if rows.Err() != nil {
		logger.Error("Error iterating city rows", zap.Error(rows.Err()))
		return nil, fmt.Errorf("error iterating city rows: %w", rows.Err())
	}
	logger.Info("Cities listed successfully", zap.Int("count", len(cities)))
	return cities, nil
}
