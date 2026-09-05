package postgres

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rafaellima1412/uber-dos-rios/internal/application/domain"
	"github.com/rafaellima1412/uber-dos-rios/internal/common"
	"github.com/rafaellima1412/uber-dos-rios/internal/logger"
	"go.uber.org/zap"
)

type shipRepositoryImpl struct {
	db *pgxpool.Pool
}

// var _ output.ShipRepository = (*shipRepositoryImpl)(nil)
func NewShipRepository(db *pgxpool.Pool) *shipRepositoryImpl {
	return &shipRepositoryImpl{db: db}
}

// CreateShip implements output.ShipRepository.
func (s *shipRepositoryImpl) CreateShip(ctx context.Context, ship *domain.Ship) error {
	tx, err := s.db.Begin(ctx)
	if err != nil {
		logger.Error("Error beginning transaction", zap.Error(err))
		return fmt.Errorf("error beginning transaction: %w", err)
	}
	defer tx.Rollback(ctx)

	// Insert principal na tabela ships
	orgSQL := `INSERT INTO ships (
		id, 
		organization_id,
		name, 
		imo,
		type_ship,
		status_ship, 
		passenger_capacity, 
		weight_capacity,
		total_seats,
		total_cabins,
		created_at,
		updated_at
	) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)
	`
	_, err = tx.Exec(ctx, orgSQL,
		ship.ID,
		ship.OrganizationID,
		ship.Name,
		ship.IMO,
		ship.TypeShip,
		ship.StatusShip,
		ship.PassengerCapacity,
		ship.WeightCapacity,
		ship.TotalSeats,
		ship.TotalCabins,
		ship.CreatedAt,
		ship.UpdatedAt,
	)
	if err != nil {
		logger.Error("Error inserting ship", zap.Error(err))
		return fmt.Errorf("error inserting ship: %w", err)
	}
	configs := []uuid.UUID{}
	if ship.ConfigurationID != nil {
		configs = *ship.ConfigurationID
	}

	if ship.TypeShip == "BOAT" && len(configs) > 0 {
		seatSQL := `INSERT INTO ship_seats (id, ship_id, seat_id,  created_at, updated_at) VALUES ($1, $2, $3, $4, $5)`
		for _, configID := range configs {
			_, err = tx.Exec(ctx, seatSQL,
				uuid.New(),
				ship.ID,
				configID,
				ship.CreatedAt,
				ship.UpdatedAt,
			)
			if err != nil {
				logger.Error("Error inserting ship_seats", zap.Error(err))
				return err
			}
		}

	} else if (ship.TypeShip == "FERRYBOAT" || ship.TypeShip == "SHIP") && len(configs) > 0 {
		cabinSQL := ` INSERT INTO ship_cabins (id, ship_id, cabin_id, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6)`
		for _, configID := range configs {
			_, err = tx.Exec(ctx, cabinSQL,
				uuid.New(),
				ship.ID,
				configID,
				ship.CreatedAt,
				ship.UpdatedAt,
			)
			if err != nil {
				logger.Error("Error inserting ship_cabins", zap.Error(err))
				return err
			}
		}
	}

	if err := tx.Commit(ctx); err != nil {
		logger.Error("Error committing transaction", zap.Error(err))
		return fmt.Errorf("error committing transaction: %w", err)
	}

	logger.Info("Ship created successfully", zap.String("ship_id", ship.ID.String()))
	return nil
}

// UpdateShip implements output.ShipRepository.
func (s *shipRepositoryImpl) UpdateShip(ctx context.Context, ship *domain.Ship) error {
	tx, err := s.db.Begin(ctx)
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback(ctx)

	updateSQL := `
	UPDATE ships 
	SET 
		name = $1, 
		imo = $2, 
		type_ship = $3,
		status_ship = $4, 
		passenger_capacity = $5, 
		weight_capacity = $6,
		total_seats = $7,
		total_cabins = $8,
		updated_at = $9,
	 	organization_id= $10,
		image_url = $11
	WHERE id = $12
	`
	_, err = tx.Exec(ctx, updateSQL,
		ship.Name,
		ship.IMO,
		ship.TypeShip,
		ship.StatusShip,
		ship.PassengerCapacity,
		ship.WeightCapacity,
		ship.TotalSeats,
		ship.TotalCabins,
		ship.UpdatedAt,
		ship.OrganizationID,
		ship.ImageUrl,
		ship.ID,
	)

	if err != nil {
		logger.Error(common.LogFailedToExecuteShipUpdate, zap.Error(err))
		return fmt.Errorf(common.ErrExecuteShipUpdate, err)
	}

	// if ship.ConfigurationID != nil {
	// 	var query string
	// 	var tableName string

	// 	switch ship.TypeShip {
	// 	case "BOAT":
	// 		tableName = "ship_seats"
	// 		query = `INSERT INTO ship_seats (id, ship_id, seat_id, updated_at)
	//           VALUES ($1, $2, $3, $4)
	//           ON CONFLICT (ship_id)
	//           DO UPDATE SET seat_id = EXCLUDED.seat_id, updated_at = EXCLUDED.updated_at`
	// 	case "FERRYBOAT", "SHIP":
	// 		tableName = "ship_cabins"
	// 		query = `INSERT INTO ship_cabins (id, ship_id, cabin_id, updated_at)
	//           VALUES ($1, $2, $3, $4)
	//           ON CONFLICT (ship_id)
	//           DO UPDATE SET updated_at = EXCLUDED.updated_at`
	// 	}

	// 	if query != "" {
	// 		result, err := tx.Exec(ctx, query, uuid.New(), ship.ID, *ship.ConfigurationID, ship.UpdatedAt)
	// 		if err != nil {
	// 			logger.Error("Error updating "+tableName, zap.Error(err))
	// 			return fmt.Errorf("error updating %s: %w", tableName, err)
	// 		}

	// 		if result.RowsAffected() == 0 {
	// 			logger.Warn("No rows updated in "+tableName, zap.String("ship_id", ship.ID.String()))
	// 		}
	// 	}
	// }
	if err := tx.Commit(ctx); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	logger.Info(common.LogInfoUpdateShip, zap.String("ship_id", ship.ID.String()))
	return nil
}

// DeleteShip implements output.ShipRepository.
func (s *shipRepositoryImpl) DeleteShip(ctx context.Context, id uuid.UUID) error {
	shipSQl := `DELETE FROM ships WHERE id = $1`
	_, err := s.db.Exec(ctx, shipSQl, id)
	if err != nil {
		return err
	}
	return nil
}

// SearchShips implements output.ShipRepository.
func (s *shipRepositoryImpl) SearchShips(ctx context.Context, search string, limit, offset int, typeShip string) ([]*domain.Ship, error) {
	sqlQuery := `
	SELECT 
		s.id,
		%s 
		organization_id,
		name, 
		imo, 
		status_ship, 
		type_ship, 
		passenger_capacity, 
		weight_capacity, 
		total_seats,
		image_url, 
		s.created_at, 
		s.updated_at 
	FROM ships s
	WHERE s.name ILIKE $1 OR s.imo ILIKE $1
	%s 
	ORDER BY created_at DESC
	LIMIT $2 OFFSET $3			
	`
	var joinClause, selectSeat string
	switch typeShip {
	case "BOAT":
		joinClause = "LEFT JOIN ship_seats ss ON s.id = ss.ship_id"
		selectSeat = "ss.seat_id,"
	case "FERRYBOAT":
		joinClause = "LEFT JOIN ship_cabins sc ON s.id = sc.ship_id"
		selectSeat = "sc.cabin_id,"
	case "SHIP":
		joinClause = "LEFT JOIN ship_cabins sc ON s.id = sc.ship_id"
		selectSeat = "sc.cabin_id,"
	default:
		joinClause = ""
		selectSeat = "NULL as seat_id,"
	}

	sqlQuery = fmt.Sprintf(sqlQuery, selectSeat, joinClause)

	rows, err := s.db.Query(ctx, sqlQuery, "%"+search+"%", limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var ships []*domain.Ship
	for rows.Next() {
		var ship domain.Ship
		if err := rows.Scan(
			&ship.ID,
			&ship.ConfigurationID,
			&ship.OrganizationID,
			&ship.Name,
			&ship.IMO,
			&ship.StatusShip,
			&ship.TypeShip,
			&ship.PassengerCapacity,
			&ship.WeightCapacity,
			&ship.TotalSeats,
			&ship.ImageUrl,
			&ship.CreatedAt,
			&ship.UpdatedAt,
		); err != nil {
			return nil, err
		}
		ships = append(ships, &ship)
	}
	if rows.Err() != nil {
		logger.Error("Error iterating Ships rows", zap.Error(rows.Err()))
		return nil, fmt.Errorf("error iterating ships rows: %w", rows.Err())
	}

	logger.Info("Ships searched successfully", zap.Int("count", len(ships)))
	return ships, nil
}

// ListShip implements output.ShipRepository.
func (s *shipRepositoryImpl) ListShips(ctx context.Context, limit, offset int, typeShip string) ([]*domain.Ship, error) {
	querySQL := `
		SELECT 
				s.id, 
				s.organization_id,
				ARRAY_REMOVE(ARRAY_AGG(DISTINCT COALESCE(ss.seat_id, sc.cabin_id)), NULL) AS configuration_ids,
				s.name, 
				s.imo, 
				s.status_ship, 
				s.type_ship, 
				s.passenger_capacity, 
				s.weight_capacity, 
				s.total_seats,
				s.total_cabins,
				s.image_url,
				s.created_at,
				s.updated_at 
		FROM ships s
		LEFT JOIN ship_seats ss ON s.id = ss.ship_id AND s.type_ship = 'BOAT'
		LEFT JOIN ship_cabins sc ON s.id = sc.ship_id AND s.type_ship IN ('FERRYBOAT', 'SHIP')
		GROUP BY s.id
		ORDER BY s.created_at DESC
		LIMIT $1 OFFSET $2
	`
	// TODO verificar forma dinamica

	rows, err := s.db.Query(ctx, querySQL, limit, offset)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var ships []*domain.Ship

	for rows.Next() {
		var ship domain.Ship
		if err := rows.Scan(
			&ship.ID,
			&ship.OrganizationID,
			&ship.ConfigurationID,
			&ship.Name,
			&ship.IMO,
			&ship.StatusShip,
			&ship.TypeShip,
			&ship.PassengerCapacity,
			&ship.WeightCapacity,
			&ship.TotalSeats,
			&ship.TotalCabins,
			&ship.ImageUrl,
			&ship.CreatedAt,
			&ship.UpdatedAt,
		); err != nil {
			return nil, err
		}
		ships = append(ships, &ship)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return ships, nil
}

// GetShip implements output.ShipRepository.
func (s *shipRepositoryImpl) GetShip(ctx context.Context, id uuid.UUID, typeShip string) (*domain.Ship, error) {
	var ship domain.Ship
	shipSQL := `
		SELECT 
			s.id,
			ARRAY_REMOVE(ARRAY_AGG(DISTINCT COALESCE(ss.seat_id, sc.cabin_id)), NULL) AS configuration_ids,
			organization_id, 
			name, 
			imo, 
			status_ship, 
			type_ship, 
			passenger_capacity, 
			weight_capacity,	
			total_seats,
			image_url, 
			s.created_at,
			s.updated_at 
		FROM ships s
		LEFT JOIN ship_seats ss ON s.id = ss.ship_id AND s.type_ship = 'BOAT'
		LEFT JOIN ship_cabins sc ON s.id = sc.ship_id AND s.type_ship IN ('FERRYBOAT', 'SHIP')
		WHERE s.id = $1
		GROUP BY s.id
		LIMIT 1
		`
	//TODO criar de forma dinamica
	err := s.db.QueryRow(ctx, shipSQL, id).Scan(
		&ship.ID,
		&ship.ConfigurationID,
		&ship.OrganizationID,
		&ship.Name,
		&ship.IMO,
		&ship.StatusShip,
		&ship.TypeShip,
		&ship.PassengerCapacity,
		&ship.WeightCapacity,
		&ship.TotalSeats,
		&ship.ImageUrl,
		&ship.CreatedAt,
		&ship.UpdatedAt,
	)
	if err != nil {
		return nil, fmt.Errorf(common.ErrExecuteShipSelect, err)
	}
	return &ship, nil
}
