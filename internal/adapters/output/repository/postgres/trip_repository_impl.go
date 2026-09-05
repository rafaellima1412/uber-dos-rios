package postgres

import (
	"context"
	"fmt"
	"strings"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/riolivre/nautical_logistics/internal/application/domain"
	"github.com/riolivre/nautical_logistics/internal/application/ports/output"
)

// tripRepositoryImpl implements the TripRepository interface.
type tripRepositoryImpl struct {
	db *pgxpool.Pool
}

// Ensure tripRepositoryImpl implements output.TripRepository.
var _ output.TripRepository = (*tripRepositoryImpl)(nil)

// NewTripRepositoryImpl creates a new instance of tripRepositoryImpl.
func NewTripRepositoryImpl(db *pgxpool.Pool) output.TripRepository {
	return &tripRepositoryImpl{
		db: db,
	}
}

// CreateTrip implements [output.TripRepository].
func (r *tripRepositoryImpl) CreateTrip(ctx context.Context, trip *domain.Trip) error {
	tx, err := r.db.Begin(ctx)
	if err != nil {
		return err
	}

	defer tx.Rollback(ctx)
	insertTripInstance := `
	INSERT INTO trips (
		id,
		trip_configurations_id,
		route_id,
		ship_id,
		departure_at,
		arrival_at,
		created_at
	)
	VALUES ($1, $2, $3, $4, $5, $6, $7)
	RETURNING id;
`

	var tripInstanceID uuid.UUID

	err = tx.QueryRow(ctx, insertTripInstance,
		trip.ID,
		trip.TripConfigurationID,
		trip.RouteID,
		trip.ShipID,
		trip.DepartureAt,
		trip.ArrivalAt,
		trip.CreatedAt,
	).Scan(&tripInstanceID)

	if err != nil {
		return fmt.Errorf("error inserting trip: %w", err)
	}

	if err := tx.Commit(ctx); err != nil {
		return fmt.Errorf("error committing transaction: %w", err)
	}
	return nil

}

// DeleteTrip implements [output.TripRepository].
func (t *tripRepositoryImpl) DeleteTrip(ctx context.Context, id uuid.UUID) error {
	panic("unimplemented")
}

// GetTrip implements [output.TripRepository].
func (t *tripRepositoryImpl) GetTrip(ctx context.Context, id uuid.UUID) (*domain.Trip, error) {
	panic("unimplemented")
}

// ListTrips implements [output.TripRepository].
func (t *tripRepositoryImpl) ListTrips(ctx context.Context, limit int, offset int) ([]*domain.Trip, error) {
	panic("unimplemented")
}

// UpdateTrip implements [output.TripRepository].
func (t *tripRepositoryImpl) UpdateTrip(ctx context.Context, trip *domain.Trip) error {
	panic("unimplemented")
}

// SearchTrips implements [output.TripRepository].
func (r *tripRepositoryImpl) FilterTrips(ctx context.Context, filter *domain.TripFilter, limit int, offset int) ([]*domain.Trip, error) {
    baseQuery := `
        SELECT 
            id, 
            route_id, 
            ship_id,
            trip_configurations_id, 
            departure_at, 
            arrival_at 
        FROM trips`
    
    var conditions []string
    var args []interface{}
    placeholderIdx := 1

    addCondition := func(condition string, value interface{}) {
        conditions = append(conditions, fmt.Sprintf("%s $%d", condition, placeholderIdx))
        args = append(args, value)
        placeholderIdx++
    }

    if filter.ShipID != nil {
        addCondition("ship_id =", filter.ShipID)
    }

    if filter.RouteID != nil {
        addCondition("route_id =", filter.RouteID)
    }

    if filter.DepartureAfter != nil && !filter.DepartureAfter.IsZero() {
        addCondition("departure_at >=", *filter.DepartureAfter)
    }

    if filter.DepartureBefore != nil && !filter.DepartureBefore.IsZero() {
        addCondition("departure_at <=", *filter.DepartureBefore)
    }

    sql := baseQuery
    if len(conditions) > 0 {
        sql += " WHERE " + strings.Join(conditions, " AND ")
    }
		
    sql += fmt.Sprintf(" ORDER BY departure_at DESC LIMIT $%d OFFSET $%d", placeholderIdx, placeholderIdx+1)
    args = append(args, limit, offset)

    rows, err := r.db.Query(ctx, sql, args...)
    if err != nil {
        return nil, fmt.Errorf("error executing query: %w", err)
    }
    defer rows.Close()

    var trips []*domain.Trip
    for rows.Next() {
        var trip domain.Trip
        err := rows.Scan(
            &trip.ID,
            &trip.RouteID,
            &trip.ShipID,
            &trip.TripConfigurationID,
            &trip.DepartureAt,
            &trip.ArrivalAt,
        )
        if err != nil {
            return nil, fmt.Errorf("error scanning trip: %w", err)
        }
        trips = append(trips, &trip)
    }

    if err := rows.Err(); err != nil {
        return nil, fmt.Errorf("error iterating trips rows: %w", err)
    }

    return trips, nil
}
