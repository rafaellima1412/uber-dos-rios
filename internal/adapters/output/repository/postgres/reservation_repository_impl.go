package postgres

import (
	"context"
	"fmt"
	"strings"

	// "time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rafaellima1412/uber-dos-rios/internal/application/domain"
	"github.com/rafaellima1412/uber-dos-rios/internal/application/ports/output"
	"github.com/rafaellima1412/uber-dos-rios/internal/logger"
	"go.uber.org/zap"
)

type reservationRepositoryImpl struct {
	db *pgxpool.Pool
}

var _ output.ReservationRepository = (*reservationRepositoryImpl)(nil)

func NewReservationRepository(db *pgxpool.Pool) *reservationRepositoryImpl {
	return &reservationRepositoryImpl{db: db}
}

func (r *reservationRepositoryImpl) CreateReservation(ctx context.Context, reservation *domain.Reservation) (string, error) {

	tx, err := r.db.Begin(ctx)
	if err != nil {
		return "", err
	}
	defer tx.Rollback(ctx)

	insertReservation := `
	INSERT INTO reservations (
		id,
		user_id,
		reservation_date,
		status,
		created_at
	) VALUES ($1, $2, $3, $4, $5)	RETURNING id;
	`

	_, err = tx.Exec(ctx, insertReservation,
		reservation.ID,
		reservation.UserID,
		reservation.ReservationDate,
		reservation.Status,
		reservation.CreatedAt,
	)
	if err != nil {
		return "", fmt.Errorf("error inserting reservation: %w", err)
	}

	if err := tx.Commit(ctx); err != nil {
		return "", fmt.Errorf("error committing transaction: %w", err)
	}

	return reservation.ID.String(), nil
}

// updateReservation implements output.ReservationRepository
func (r *reservationRepositoryImpl) UpdateReservation(ctx context.Context, reservation *domain.Reservation) error {
	var reservationUpdate domain.Reservation
	updateSQL := `
	UPDATE reservations 
	SET 
		user_id = $1, 
		trip_id = $2, 
		reservation_date = $3, 
		status = $4
		updated_at = $5
	WHERE id = $6
	`
	_, err := r.db.Exec(ctx, updateSQL,
		&reservation.UserID,
		&reservation.ReservationDate,
		&reservation.Status,
		&reservation.UpdatedAt,
		&reservation.ID,
	)

	if err != nil {
		logger.Error("failed to execute reservation update", zap.Error(err))
		// (L) quem chama a interface trate os erros
		return fmt.Errorf("failed to execute reservation update: %w", err)
	}
	logger.Info("reservation updated successfully", zap.String("reservation_id", reservationUpdate.ID.String()))
	return nil

}

// GetReservation implements output.ReservationRepository.
func (r *reservationRepositoryImpl) GetReservation(ctx context.Context, id uuid.UUID) (*domain.Reservation, error) {
	var reservation domain.Reservation
	query := `
	SELECT 
		id,
		user_id,
		reservation_date,
		status,
		created_at,
		updated_at
	FROM reservations 
	WHERE id = $1
	`

	err := r.db.QueryRow(ctx, query, id).Scan(
		&reservation.ID,
		&reservation.UserID,
		&reservation.ReservationDate,
		&reservation.Status,
		&reservation.CreatedAt,
		&reservation.UpdatedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("error getting reservation: %w", err)
	}
	return &reservation, nil

}

// ListReservation
func (r *reservationRepositoryImpl) ListReservations(ctx context.Context, limit, offset int) ([]*domain.Reservation, error) {

	query := `
		select
			r.id,
			r.user_id,
			r.status,
			r.reservation_date,
			r.created_at,
			r.updated_at
		FROM reservations r 
		ORDER BY r.created_at DESC
		LIMIT $1 OFFSET $2
	`

	rows, err := r.db.Query(ctx, query, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("query error: %w", err)
	}
	defer rows.Close()

	var reservations []*domain.Reservation

	for rows.Next() {
		var reservation domain.Reservation

		if err := rows.Scan(
			&reservation.ID,
			&reservation.UserID,
			&reservation.Status,
			&reservation.ReservationDate,
			&reservation.CreatedAt,
			&reservation.UpdatedAt,
		); err != nil {
			logger.Error("Error scanning reservations", zap.Error(err))
			return nil, fmt.Errorf("scan error: %w", err)
		}

		reservations = append(reservations, &reservation)
	}

	return reservations, nil
}

// delete
func (r *reservationRepositoryImpl) DeleteReservation(ctx context.Context, id uuid.UUID) error {
	query := `DELETE FROM reservations WHERE id = $1`
	_, err := r.db.Exec(ctx, query, id)
	if err != nil {
		logger.Error("Error deleting reservation", zap.Error(err))
		return fmt.Errorf("error deleting reservation: %w", err)
	}
	logger.Info("Reservation deleted successfully", zap.String("reservation_id", id.String()))
	return nil
}

func (r *reservationRepositoryImpl) CreateTripReservation(ctx context.Context, tripInstance *domain.Trip, reservation *domain.TripToReservation) error {
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
		tripInstance.ID,
		tripInstance.TripConfigurationID,
		tripInstance.RouteID,
		tripInstance.ShipID,
		tripInstance.DepartureAt,
		tripInstance.ArrivalAt,
		tripInstance.CreatedAt,
	).Scan(&tripInstanceID)

	if err != nil {
		return fmt.Errorf("error inserting trip_instance: %w", err)
	}
	if reservation != nil && len(reservation.OccupiedUnits) > 0 {
		var seats []string
		var cabins []string
		for _, unit := range reservation.OccupiedUnits {
			if strings.HasPrefix(unit, "CA") {
				cabins = append(cabins, unit)
			} else {
				seats = append(seats, unit)
			}
		}
		insertJoin := `
        INSERT INTO reservations_trips (
            reservation_id,
            trip_instances_id,
            occupied_seats,
            occupied_cabins
        ) VALUES ($1, $2, $3, $4);
    `

		_, err = tx.Exec(ctx, insertJoin,
			reservation.ReservationID,
			tripInstanceID,
			seats,
			cabins,
		)

		if err != nil {
			return fmt.Errorf("error inserting mixed units: %w", err)
		}
	}
	if err := tx.Commit(ctx); err != nil {
		return fmt.Errorf("error committing transaction: %w", err)
	}
	return nil

}
