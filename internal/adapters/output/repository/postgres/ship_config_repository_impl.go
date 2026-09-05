package postgres

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/riolivre/nautical_logistics/internal/application/domain"
	"github.com/riolivre/nautical_logistics/internal/application/ports/output"
	"github.com/riolivre/nautical_logistics/internal/common"
	"github.com/riolivre/nautical_logistics/internal/logger"
	"go.uber.org/zap"
)

type configRepositoryImpl struct {
	db *pgxpool.Pool
}

var _ output.ShipConfigRepository = (*configRepositoryImpl)(nil)

func NewSeatRepository(db *pgxpool.Pool) *configRepositoryImpl {
	return &configRepositoryImpl{db: db}
}

// CreateSeat implements output.SeatRepository.
func (s *configRepositoryImpl) CreateSeat(ctx context.Context, seat domain.Seat) (*domain.Seat, error) {
	tx, err := s.db.Begin(ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback(ctx)

	query := `
	INSERT INTO seats (
		id,
		is_active,
		left_columns_count,
		right_columns_count,
		prefix,
		created_at,
		updated_at
	)
	VALUES ($1, $2, $3, $4, $5, $6, $7)
	RETURNING id;
	`

	row := tx.QueryRow(ctx,
		query,
		seat.ID,
		seat.IsActive,
		seat.LeftColumnsCount,
		seat.RightColumnsCount,
		seat.Prefix,
		seat.CreatedAt,
		seat.UpdatedAt,
	)

	if err := row.Scan(&seat.ID); err != nil {
		return nil, err
	}

	if err := tx.Commit(ctx); err != nil {
		return nil, err
	}

	return &seat, nil
}

// UpdateSeat
func (s *configRepositoryImpl) UpdateSeat(ctx context.Context, seat domain.Seat, id uuid.UUID) error {
	var seatUpdate domain.Seat
	updateSQL := `
	UPDATE seats 
	SET 
		is_active = $2, 
		left_columns_count = $3, 
		right_columns_count = $4,
		prefix = $5,
		updated_at = $6
	WHERE id =	$1
	RETURNING id;
	`
	err := s.db.QueryRow(ctx, updateSQL,
		seat.ID,
		seat.IsActive,
		seat.LeftColumnsCount,
		seat.RightColumnsCount,
		seat.Prefix,
		seat.UpdatedAt,
	).Scan(
		&seatUpdate.ID,
	)
	if err != nil {
		logger.Error(common.LogFailedToExecuteSeatUpdate, zap.Error(err))
		return fmt.Errorf(common.ErrExecuteSeatUpdate, err)
	}
	logger.Info(common.LogInfoUpdateSeat, zap.String("seat_id", seatUpdate.ID.String()))
	return nil
}

// GetSeat
func (s *configRepositoryImpl) GetSeat(ctx context.Context, id uuid.UUID) (*domain.Seat, error) {
	var seat domain.Seat
	query := `
	SELECT 
		id, 
		is_active, 
		left_columns_count, 
		right_columns_count,  
		prefix,
		created_at, 
		updated_at
	FROM seats 
	WHERE id = $1
	`
	err := s.db.QueryRow(ctx, query, id).Scan(
		&seat.ID,
		&seat.IsActive,
		&seat.LeftColumnsCount,
		&seat.RightColumnsCount,
		&seat.Prefix,
		&seat.CreatedAt,
		&seat.UpdatedAt,
	)
	if err != nil {
		return nil, fmt.Errorf(common.ErrExecuteSeatSelect, err)
	}
	return &seat, nil
}

// ListSeats
func (s *configRepositoryImpl) ListSeats(ctx context.Context, limit, offset int) ([]*domain.Seat, error) {
	query := `
	SELECT 
		id, 
		is_active, 
		left_columns_count, 
		right_columns_count, 		
		prefix,
		created_at, 
		updated_at
	FROM seats 
	ORDER BY created_at DESC
	LIMIT $1 OFFSET $2
	`
	rows, err := s.db.Query(ctx, query, limit, offset)
	if err != nil {
		logger.Error(common.LogFailedToExecuteSeatList, zap.Error(err))
		return nil, fmt.Errorf(common.ErrExecuteSeatList, err)
	}
	defer rows.Close()

	var seats []*domain.Seat
	for rows.Next() {
		var seat domain.Seat
		if err := rows.Scan(
			&seat.ID,
			&seat.IsActive,
			&seat.LeftColumnsCount,
			&seat.RightColumnsCount,
			&seat.Prefix,
			&seat.CreatedAt,
			&seat.UpdatedAt,
		); err != nil {
			logger.Error(common.LogFailedToScanSeatRow, zap.Error(err))
			return nil, fmt.Errorf(common.ErrScanSeatRow, err)
		}
		seats = append(seats, &seat)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return seats, nil
}

// SearchSeats by active status or prefix
func (s *configRepositoryImpl) SearchSeats(ctx context.Context, isActive bool, prefix bool, limit, offset int) ([]*domain.Seat, error) {
	query := `
		SELECT 
				id, 
				is_active, 
				left_columns_count, 
				right_columns_count, 
				prefix,
				created_at, 
				updated_at		
		FROM seats 
		WHERE is_active = $1 AND prefix = $2
		ORDER BY created_at DESC
		LIMIT $3 OFFSET $4
	`
	rows, err := s.db.Query(ctx, query, isActive, prefix, limit, offset)
	if err != nil {
		logger.Error("Error searching seats", zap.Error(err))
		return nil, fmt.Errorf("error searching seats: %w", err)
	}
	defer rows.Close()
	var seats []*domain.Seat
	for rows.Next() {
		var seat domain.Seat
		if err := rows.Scan(
			&seat.ID,
			&seat.IsActive,
			&seat.LeftColumnsCount,
			&seat.RightColumnsCount,
			&seat.Prefix,
			&seat.CreatedAt,
			&seat.UpdatedAt,
		); err != nil {
			logger.Error(common.LogFailedToScanSeatRow, zap.Error(err))
			return nil, fmt.Errorf(common.ErrScanSeatRow, err)
		}
		seats = append(seats, &seat)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return seats, nil
}

// DeleteSeats
func (r *configRepositoryImpl) DeleteSeat(ctx context.Context, id uuid.UUID) error {
	query := `DELETE FROM seats WHERE id = $1`
	_, err := r.db.Exec(ctx, query, id)
	if err != nil {
		logger.Error("Error deleting seat", zap.Error(err))
		return fmt.Errorf("error deleting seat: %w", err)
	}

	logger.Info("Seat deleted successfully", zap.String("seat_id", id.String()))
	return nil
}

// CreateCabin
func (s *configRepositoryImpl) CreateCabins(ctx context.Context, cabins []domain.Cabin) error {
	tx, err := s.db.Begin(ctx)
	if err != nil {
		return nil
	}
	defer tx.Rollback(ctx)

	query := `
	INSERT INTO cabins (
		id,
		name,
		bed_type,
		capacity,
		description,
		created_at,
		updated_at
	)
	VALUES ($1, $2, $3, $4, $5, $6, $7);
	`
	for _, c := range cabins {
		for _, detail := range c.Cabins {
			_, err := tx.Exec(ctx, query,
				detail.ID,
				detail.Name,
				detail.BedType,
				detail.Capacity,
				detail.Description,
				detail.CreatedAt,
				detail.UpdatedAt,
			)
			if err != nil {
				return err
			}
		}
	}

	if err := tx.Commit(ctx); err != nil {
		return fmt.Errorf("falha ao finalizar transação: %w", err)
	}

	return nil
}

// ListCabins
func (s *configRepositoryImpl) ListCabins(ctx context.Context, limit, offset int) ([]*domain.CabinDetails, error) {
	query := `
	SELECT 
		id, 
		name, 
		bed_type, 
		description, 
		capacity,
		created_at, 
		updated_at
	FROM cabins 
	ORDER BY created_at DESC
	LIMIT $1 OFFSET $2
	`
	rows, err := s.db.Query(ctx, query, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var cabins []*domain.CabinDetails
	for rows.Next() {
		var cabin domain.CabinDetails
		if err := rows.Scan(
			&cabin.ID,
			&cabin.Name,
			&cabin.BedType,
			&cabin.Description,
			&cabin.Capacity,
			&cabin.CreatedAt,
			&cabin.UpdatedAt,
		); err != nil {

			return nil, err
		}
		cabins = append(cabins, &cabin)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return cabins, nil
}

// deleteCabins
func (s *configRepositoryImpl) DeleteCabin(ctx context.Context, id uuid.UUID) error {
	query := `DELETE FROM cabins WHERE id = $1`
	_, err := s.db.Exec(ctx, query, id)
	if err != nil {
		logger.Error("Error deleting cabin", zap.Error(err))
		return fmt.Errorf("error deleting cabin: %w", err)
	}
	logger.Info("Cabin deleted successfully", zap.String("cabin_id", id.String()))
	return nil
}

// UpdateCabin
func (s *configRepositoryImpl) UpdateCabin(ctx context.Context, cabin domain.CabinDetails, id uuid.UUID) error {
	var cabinUpdate domain.CabinDetails
	updateSQL := `
	UPDATE cabins 
	SET 
		name = $1, 
		bed_type = $2, 
		description = $3,
		capacity = $4,
		updated_at = $5
	WHERE id = $6
	RETURNING id;
	`
	err := s.db.QueryRow(ctx, updateSQL,
		cabin.Name,
		cabin.BedType,
		cabin.Description,
		cabin.Capacity,
		cabin.UpdatedAt,
		id,
	).Scan(
		&cabinUpdate.ID,
	)
	if err != nil {
		logger.Error("Error updating cabin", zap.Error(err))
		return fmt.Errorf("error updating cabin: %w", err)
	}
	logger.Info("Cabin updated successfully", zap.String("cabin_id", cabinUpdate.ID.String()))
	return nil
}

// GetCabin
func (s *configRepositoryImpl) GetCabin(ctx context.Context, id uuid.UUID) (*domain.CabinDetails, error) {
	var cabin domain.CabinDetails
	query := `
	SELECT 
		id, 
		name, 
		bed_type, 
		description, 
		capacity,
		created_at, 
		updated_at
	FROM cabins 
	WHERE id = $1
	`
	err := s.db.QueryRow(ctx, query, id).Scan(
		&cabin.ID,
		&cabin.Name,
		&cabin.BedType,
		&cabin.Description,
		&cabin.Capacity,
		&cabin.CreatedAt,
		&cabin.UpdatedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("error getting cabin: %w", err)
	}
	return &cabin, nil
}

// SearchSeatByDate implements [output.ShipConfigRepository].
func (s *configRepositoryImpl) SearchSeatByShipAndDate(ctx context.Context, shipID uuid.UUID, dateOrigin, dateDest time.Time) (*domain.UnitsAvailable, error) {
	querySql := `
		WITH seat_data AS (
				SELECT 
						ti.ship_id,
						COUNT(DISTINCT seat) AS occupied_seats_count,
						ARRAY_AGG(DISTINCT seat ORDER BY seat) FILTER (WHERE seat IS NOT NULL) AS occupied_seats_list
				FROM trips ti
				JOIN reservations_trips rti ON rti.trip_instances_id = ti.id
				LEFT JOIN LATERAL jsonb_array_elements_text(rti.occupied_seats) AS seat ON true
				WHERE ti.ship_id = $1
					AND ti.departure_at >= $2
					AND ti.arrival_at   <= $3
				GROUP BY ti.ship_id
		),
		cabin_data AS (
				SELECT 
						ti.ship_id,
						COUNT(DISTINCT cabin) AS occupied_cabins_count,
						ARRAY_AGG(DISTINCT cabin ORDER BY cabin) FILTER (WHERE cabin IS NOT NULL) AS occupied_cabins_list
				FROM trips ti
				JOIN reservations_trips rti ON rti.trip_instances_id = ti.id
				LEFT JOIN LATERAL jsonb_array_elements_text(rti.occupied_cabins) AS cabin ON true
				WHERE ti.ship_id = $1
					AND ti.departure_at >= $2
					AND ti.arrival_at   <= $3
				GROUP BY ti.ship_id
		)
		SELECT
				s.id,
				s.name,
				s.total_seats,
				s.total_seats - COALESCE(sd.occupied_seats_count, 0) AS available_seats,
				COALESCE(sd.occupied_seats_list, '{}') AS occupied_seats,
				s.total_cabins,
				s.total_cabins - COALESCE(cd.occupied_cabins_count, 0) AS available_cabins,
				COALESCE(cd.occupied_cabins_list, '{}') AS occupied_cabins
		FROM ships s
		LEFT JOIN seat_data  sd ON sd.ship_id = s.id
		LEFT JOIN cabin_data cd ON cd.ship_id = s.id
		WHERE s.id = $1;
			`

	var units domain.UnitsAvailable

	// Usamos QueryRow quando esperamos apenas um resultado
	err := s.db.QueryRow(ctx, querySql, shipID, dateOrigin, dateDest).Scan(
		&units.ShipID,
		&units.ShipName,
		&units.Seats.TotalSeats,
		&units.Seats.AvailableSeats,
		&units.Seats.OccupiedUnits,
		&units.Cabins.TotalCabins,
		&units.Cabins.AvailableCabins,
		&units.Cabins.OccupiedUnits,
	)

	if err != nil {
		return nil, err
	}

	return &units, nil
}
