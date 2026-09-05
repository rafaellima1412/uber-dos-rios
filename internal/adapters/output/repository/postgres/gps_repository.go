package postgres

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rafaellima1412/uber-dos-rios/internal/application/domain"
	"github.com/rafaellima1412/uber-dos-rios/internal/application/ports/output"
	"github.com/rafaellima1412/uber-dos-rios/internal/logger"
	"go.uber.org/zap"
)

type gpsRepositoryImpl struct {
	db *pgxpool.Pool
}
var _ output.GPSRepository = (*gpsRepositoryImpl)(nil)

// NewGPSRepositoryImpl creates a new instance of gpsRepositoryImpl.
func NewGPSRepository(db *pgxpool.Pool) output.GPSRepository {
	return &gpsRepositoryImpl{
		db: db,
	}
}

// CreateGPS implements output.GPSRepository.
func (r *gpsRepositoryImpl) CreateGPS(ctx context.Context, gps *domain.GPS) error {
	tx, err := r.db.Begin(ctx)
	if err != nil {
		logger.Error("Error beginning transaction", zap.Error(err))
		return fmt.Errorf("error beginning transaction: %w", err)
	}
	defer tx.Rollback(ctx)

	query := `
		INSERT INTO gps_positions  (
			id, 
			ship_name,
			latitude,
			longitude,
			speed,
			course,
			received_at
		) VALUES ($1, $2, $3, $4, $5, $6, $7)
	`
	_, err = tx.Exec(ctx, query,
		gps.ID,
		gps.ShipName,
		gps.Latitude,
		gps.Longitude,
		gps.Speed,
		gps.Course,
		gps.ReceivedAt,
	)
	if err != nil {
		logger.Error("Error inserting GPS", zap.Error(err))
		return fmt.Errorf("error inserting GPS: %w", err)
	}

	if err := tx.Commit(ctx); err != nil {
		return err
	}
	logger.Info("GPS POSITION created successfully", zap.String("gps_id", gps.ID.String()))
	return nil
}

// GetGPS implements output.GPSRepository.
func (r *gpsRepositoryImpl) GetGPS(ctx context.Context, id uuid.UUID) (*domain.GPS, error) {
	query := `
		SELECT 
			id, 
			ship_name,
			latitude,
			longitude,
			speed,
			course,
			received_at
		FROM gps 
		WHERE id = $1`
	row := r.db.QueryRow(ctx, query, id)

	var gps domain.GPS
	err := row.Scan(
		&gps.ID,
		&gps.ShipName,
		&gps.Latitude,
		&gps.Longitude,
		&gps.Speed,
		&gps.Course,
		&gps.ReceivedAt,

	)
	if err != nil {
		if err == pgx.ErrNoRows {
			logger.Error("GPS not found", zap.Error(err))
			return nil, fmt.Errorf("gps not found: %w", err)
		}
		logger.Error("Error finding GPS by ID", zap.Error(err))
		return nil,
			fmt.Errorf("error finding GPS by ID: %w", err)
	}

	return &gps, nil
}
