package postgres

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rafaellima1412/uber-dos-rios/internal/application/domain"
	"github.com/rafaellima1412/uber-dos-rios/internal/application/ports/input"
	"github.com/rafaellima1412/uber-dos-rios/internal/application/ports/output"
	"github.com/rafaellima1412/uber-dos-rios/internal/logger"
	"go.uber.org/zap"
)

// tripConfigRepositoryImpl implements the TripConfigRepository interface.
type tripConfigRepositoryImpl struct {
	db *pgxpool.Pool
}

// Ensure tripRepositoryImpl implements output.TripConfigRepository.
var _ output.TripConfigRepository = (*tripConfigRepositoryImpl)(nil)

func mergeDateAndTime(dateStr string, clock time.Time) (time.Time, error) {
	layout := "2006-01-02" // Ajuste conforme o formato da sua string
	d, err := time.Parse(layout, dateStr)
	if err != nil {
		return time.Time{}, err
	}

	return time.Date(
		d.Year(), d.Month(), d.Day(),
		clock.Hour(), clock.Minute(), clock.Second(),
		0, clock.Location(),
	), nil
}

// NewTripRepositoryImpl creates a new instance of tripRepositoryImpl.
func NewTripConfigRepositoryImpl(db *pgxpool.Pool) output.TripConfigRepository {
	return &tripConfigRepositoryImpl{
		db: db,
	}
}

// CreateTrip implements output.TripRepository.
func (r *tripConfigRepositoryImpl) CreateTripConfig(ctx context.Context, trip *domain.TripConfig) error {
	tx, err := r.db.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)
	query := `
		INSERT INTO trip_configurations (
			id, 
			ship_id, 
			route_id, 
			recurrence, 
			expiration_date, 
			created_at, 
			departure_time,
			arrival_time, 
			duration_days,
			start_date
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) RETURNING ID;
	`
	_, err = tx.Exec(ctx, query,
		trip.ID,
		trip.ShipID,
		trip.RouteID,
		trip.Recurrence,
		trip.ExpirationDate,
		trip.CreatedAt,
		trip.DepartureTime,
		trip.ArrivalTime,
		trip.DurationDays,
		trip.StartDate,
	)
	if err != nil {
		return err
	}
	if err := tx.Commit(ctx); err != nil {
		return err
	}
	return nil
}

// UpdateTrip implements output.TripRepository.
func (r *tripConfigRepositoryImpl) UpdateTripConfig(ctx context.Context, trip *domain.TripConfig) error {
	query := `
		UPDATE trip_configurations 
		SET 
			ship_id = $1, 
			route_id = $2, 
			recurrence = $3, 
			expiration_date = $4, 
			departure_time = $5, 
			start_date = $6,
			duration_days = $7,
			updated_at = $8
		WHERE id = $9
	`
	_, err := r.db.Exec(ctx, query,
		&trip.ShipID,
		&trip.RouteID,
		&trip.Recurrence,
		&trip.ExpirationDate,
		&trip.DepartureTime,
		&trip.StartDate,
		&trip.DurationDays,
		&trip.UpdatedAt,
		&trip.ID,
	)
	if err != nil {
		logger.Error("Error updating trip", zap.Error(err))
		return fmt.Errorf("error updating trip: %w", err)
	}
	logger.Info("Trip updated successfully", zap.String("trip_id", trip.ID.String()))
	return nil
}

// GetTrip implements output.TripRepository.
func (r *tripConfigRepositoryImpl) GetTripConfig(ctx context.Context, id uuid.UUID) (*domain.TripConfig, error) {
	var trip domain.TripConfig
	query := `
		SELECT
			tc.id, 
			tc.ship_id, 
			tc.route_id,
			s."name" as ship_name,
			r."name" as route_name,
			tc.recurrence, 
			tc.expiration_date,
			tc.start_date,
			tc.departure_time,
			tc.arrival_time,
			tc.duration_days,
			tc.created_at, 
			tc.updated_at
		FROM trip_configurations tc
		inner join ships s
		on tc.ship_id = s.id
		inner join routes r
		on tc.route_id = r.id
		WHERE tc.id = $1
	`
	err := r.db.QueryRow(ctx, query, id).Scan(
		&trip.ID,
		&trip.ShipID,
		&trip.RouteID,
		&trip.ShipName,
		&trip.RouteName,
		&trip.Recurrence,
		&trip.ExpirationDate,
		&trip.StartDate,
		&trip.DepartureTime,
		&trip.ArrivalTime,
		&trip.DurationDays,
		&trip.CreatedAt,
		&trip.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &trip, nil
}

// DeleteTrip implements output.TripRepository.
func (r *tripConfigRepositoryImpl) DeleteTripConfig(ctx context.Context, id uuid.UUID) error {
	query := `DELETE FROM trip_configurations WHERE id = $1`
	_, err := r.db.Exec(ctx, query, id)
	if err != nil {
		return err
	}
	return nil
}

// ListTrips implements output.TripRepository.
func (r *tripConfigRepositoryImpl) ListTripsConfig(ctx context.Context, limit int, offset int) ([]*domain.TripConfig, error) {
	query := `
		SELECT
			t.id, 
			ship_id,
			route_id,
			s."name" as ship_name,
			r."name" as route_name,
			t.recurrence, 
			t.expiration_date, 
			t.departure_time,
			t.arrival_time, 
			t.duration_days,
			t.start_date, 
			t.created_at, 
			t.updated_at
		FROM trip_configurations t 
		inner join routes r 
		on t.route_id = r.id
		inner join ships s
		on t.ship_id = s.id
		LIMIT $1 OFFSET $2
	`
	rows, err := r.db.Query(ctx, query, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var trips []*domain.TripConfig

	for rows.Next() {
		var trip domain.TripConfig
		err := rows.Scan(
			&trip.ID,
			&trip.ShipID,
			&trip.RouteID,
			&trip.ShipName,
			&trip.RouteName,
			&trip.Recurrence,
			&trip.ExpirationDate,
			&trip.DepartureTime,
			&trip.ArrivalTime,
			&trip.DurationDays,
			&trip.StartDate,
			&trip.CreatedAt,
			&trip.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		trips = append(trips, &trip)
	}
	return trips, nil
}

// LoadGraph implements [output.TripRepository].
func (r *tripConfigRepositoryImpl) LoadGraph(ctx context.Context, DateOrigin, DateDest string) (input.Graph, error) {
	queryStr := `
		SELECT
				t_from.city_id AS from_city,
				t_to.city_id   AS to_city,
				s.value        AS cost,
				r.organization_id,
				r.id           AS route_id,
				r."name" as router_name,
				tc.arrival_time as hora_chegada,
				tc.departure_time as hora_saida,
				tc.ship_id,
				s2."name" as ship_name,
				COALESCE(s2.image_url[1], '') AS image_url,
				tc.id
		FROM schedules s
		JOIN schedules s_next
			ON s.route_id = s_next.route_id
			AND s.stop_order + 1 = s_next.stop_order
			JOIN terminals t_from ON s.terminal_id = t_from.id
			JOIN terminals t_to   ON s_next.terminal_id = t_to.id
			JOIN routes r ON r.id = s.route_id
			JOIN trip_configurations tc ON tc.route_id = r.id
			join ships s2 on s2.id = tc.ship_id 
    WHERE tc.start_date <= $1
  		AND tc.expiration_date >= $2;
	`
	rows, err := r.db.Query(ctx, queryStr, DateOrigin, DateDest)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	graph := make(input.Graph)

	for rows.Next() {
		var from, to, cost int
		var orgID, routeID, shipID, tripID uuid.UUID
		var routerName, shipName, imageURL string
		var dateDepartureTime, dataArrivalTime time.Time
		if err := rows.Scan(
			&from,
			&to,
			&cost,
			&orgID,
			&routeID,
			&routerName,
			&dataArrivalTime,
			&dateDepartureTime,
			&shipID,
			&shipName,
			&imageURL,
			&tripID,
		); err != nil {
			return nil, err
		}

		if _, ok := graph[from]; !ok {
			graph[from] = []*input.Edge{}
		}

		if _, ok := graph[to]; !ok {
			graph[to] = []*input.Edge{}
		}
		dateDepartureTimeFinal, _ := mergeDateAndTime(DateOrigin, dateDepartureTime)
		dateArrivalTimefinal, _ := mergeDateAndTime(DateDest, dataArrivalTime)

		graph[from] = append(graph[from], &input.Edge{
			To:                  to,
			Cost:                cost,
			OrganizationID:      orgID,
			RouteID:             routeID,
			RouteName:           routerName,
			DateDepartureTime:   dateDepartureTimeFinal,
			DateArrivalTime:     dateArrivalTimefinal,
			ShipID:              shipID,
			ShipName:            shipName,
			ShipImageURL:        imageURL,
			TripConfigurationID: tripID,
		})
	}

	return graph, nil

}
