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

// routeRepositoryImpl implements the RouteRepository interface.
type routeRepositoryImpl struct {
	db *pgxpool.Pool
}

// Ensure routeRepositoryImpl implements output.RouteRepository.
var _ output.RouteRepository = (*routeRepositoryImpl)(nil)

// NewRouteRepositoryImpl creates a new instance of routeRepositoryImpl.
func NewRouteRepositoryImpl(db *pgxpool.Pool) output.RouteRepository {
	return &routeRepositoryImpl{
		db: db,
	}
}

// CreateRoute implements output.RouteRepository.
func (r *routeRepositoryImpl) CreateRoute(ctx context.Context, route *domain.Route) error {
	tx, err := r.db.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	query := `
		INSERT INTO routes (
			id, 
			name, 
			active, 
			organization_id, 
			created_at, 
			updated_at
		) VALUES ($1, $2, $3, $4, $5, $6)
	`
	_, err = tx.Exec(ctx, query,
		route.ID,
		route.Name,
		route.Active,
		route.OrganizationID,
		route.CreatedAt,
		route.UpdatedAt,
	)
	if err != nil {
		return err
	}

	if err := tx.Commit(ctx); err != nil {
		return err
	}

	return nil
}

// DeleteRoute implements output.RouteRepository.
func (r *routeRepositoryImpl) DeleteRoute(ctx context.Context, id uuid.UUID) error {
	query := `DELETE FROM routes WHERE id = $1`
	_, err := r.db.Exec(ctx, query, id)
	if err != nil {
		return err
	}
	return nil
}

// GetRoute implements output.RouteRepository.
func (r *routeRepositoryImpl) GetRoute(ctx context.Context, id uuid.UUID) (*domain.Route, error) {
	query := `
		SELECT 
			id, 
			name, 
			active, 
			organization_id
		FROM routes 
		WHERE id = $1`
	row := r.db.QueryRow(ctx, query, id)

	var route domain.Route
	err := row.Scan(
		&route.ID,
		&route.Name,
		&route.Active,
		&route.OrganizationID,
	)
	if err != nil {
		if err == pgx.ErrNoRows {
			logger.Error("Route not found", zap.Error(err))
			return nil, fmt.Errorf("route not found: %w", err)
		}
		logger.Error("Error finding route by ID", zap.Error(err))
		return nil, fmt.Errorf("error finding route by ID: %w", err)
	}

	return &route, nil
}

// ListRoutes implements output.RouteRepository.
func (r *routeRepositoryImpl) ListRoutes(ctx context.Context, limit int, offset int) ([]*domain.Route, error) {
	query := `
		SELECT 
			id, 
			name, 
			active, 
			organization_id
		FROM routes 
		LIMIT $1 OFFSET $2`
	rows, err := r.db.Query(ctx, query, limit, offset)
	if err != nil {
		logger.Error("Error listing routes", zap.Error(err))
		return nil, fmt.Errorf("error listing routes: %w", err)
	}
	defer rows.Close()

	var routes []*domain.Route
	for rows.Next() {
		var route domain.Route
		err := rows.Scan(
			&route.ID,
			&route.Name,
			&route.Active,
			&route.OrganizationID,
		)
		if err != nil {
			logger.Error("Error scanning route", zap.Error(err))
			return nil, fmt.Errorf("error scanning route: %w", err)
		}
		routes = append(routes, &route)
	}

	if rows.Err() != nil {
		logger.Error("Error iterating over routes", zap.Error(rows.Err()))
		return nil, fmt.Errorf("error iterating over routes: %w", rows.Err())
	}

	return routes, nil
}

// SearchRoutes implements output.RouteRepository.
func (r *routeRepositoryImpl) SearchRoutes(ctx context.Context, query string, limit int, offset int) ([]*domain.Route, error) {
	searchQuery := `
		SELECT 
			id, 
			name, 
			active, 
			organization_id
		FROM routes 
		WHERE name ILIKE '%' || $1 || '%'
		LIMIT $2 OFFSET $3`
	rows, err := r.db.Query(ctx, searchQuery, query, limit, offset)
	if err != nil {
		logger.Error("Error searching routes", zap.Error(err))
		return nil, fmt.Errorf("error searching routes: %w", err)
	}
	defer rows.Close()

	var routes []*domain.Route
	for rows.Next() {
		var route domain.Route
		err := rows.Scan(
			&route.ID,
			&route.Name,
			&route.Active,
			&route.OrganizationID,
		)
		if err != nil {
			logger.Error("Error scanning route", zap.Error(err))
			return nil, fmt.Errorf("error scanning route: %w", err)
		}
		routes = append(routes, &route)
	}

	if rows.Err() != nil {
		logger.Error("Error iterating over searched routes", zap.Error(rows.Err()))
		return nil, fmt.Errorf("error iterating over searched routes: %w", rows.Err())
	}

	return routes, nil
}

// UpdateRoute implements output.RouteRepository.
func (r *routeRepositoryImpl) UpdateRoute(ctx context.Context, route *domain.Route) error {
	query := `
		UPDATE routes 
		SET name = $1, active = $2, updated_at = $3
		WHERE id = $4
	`
	_, err := r.db.Exec(ctx, query,
		route.Name,
		route.Active,
		route.UpdatedAt,
		route.ID,
	)
	if err != nil {
		logger.Error("Error updating route", zap.Error(err))
		return fmt.Errorf("error updating route: %w", err)
	}
	return nil
}
