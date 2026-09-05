package usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/riolivre/nautical_logistics/internal/application/domain"
	"github.com/riolivre/nautical_logistics/internal/application/ports/input"
	"github.com/riolivre/nautical_logistics/internal/application/ports/output"
	"github.com/riolivre/nautical_logistics/internal/logger"
	"go.uber.org/zap"
)

type createTripConfigUseCase struct {
	tripRepository  output.TripConfigRepository
	shipRepository  output.ShipRepository
	routeRepository output.RouteRepository
}

func NewCreateTripConfigUseCase(
	tripRepo output.TripConfigRepository,
	shipRepo output.ShipRepository,
	routeRepo output.RouteRepository,
) *createTripConfigUseCase {
	return &createTripConfigUseCase{
		tripRepository:  tripRepo,
		shipRepository:  shipRepo,
		routeRepository: routeRepo,
	}
}
func (uc *createTripConfigUseCase) Execute(ctx context.Context, inputTrip *input.CreateTripConfigInput) error {
	// paerse inputTrip.ShipID to uuid
	shipID, err := uuid.Parse(inputTrip.ShipID)
	if err != nil {
		return fmt.Errorf("invalid ShipID format: %w", err)
	}
	ship, err := uc.shipRepository.GetShip(ctx, shipID, "all")
	if err != nil {
		return fmt.Errorf("failed to find ship: %w", err)
	}

	// paerse inputTrip.ShipID to uuid
	routeID, err := uuid.Parse(inputTrip.RouteID)
	if err != nil {
		return fmt.Errorf("invalid RouteID format: %w", err)
	}
	route, err := uc.routeRepository.GetRoute(ctx, routeID)
	if err != nil {
		return fmt.Errorf("failed to find route: %w", err)
	}
	// Generate new Trip ID
	tripID, err := uuid.NewV7()
	if err != nil {
		logger.Error("Error generating UUID", zap.Error(err))
		return fmt.Errorf("error generating UUID: %w", err)
	}
	dept, err := time.Parse("15:04:05", inputTrip.DepartureTime)
	if err != nil {
		logger.Error("Error Parse Date", zap.Error(err))
		return fmt.Errorf("error parse departure_time: %w", err)
	}

	departureTime := time.Date(
		2000, 1, 1,
		dept.Hour(),
		dept.Minute(),
		dept.Second(),
		0,
		time.UTC,
	)
	arrt, err := time.Parse("15:04:05", inputTrip.ArrivalTime)
	if err != nil {
		logger.Error("Error Parse Date", zap.Error(err))
		return fmt.Errorf("error parse departure_time: %w", err)
	}
	arrivalTime := time.Date(
		2000, 1, 1,
		arrt.Hour(),
		arrt.Minute(),
		arrt.Second(),
		0,
		time.UTC,
	)
	// Validação de 8 horas
	diff := arrivalTime.Sub(departureTime)

	// Se a viagem vira o dia (ex: sai 23:00 e chega 07:00)
	// aqui so consigo validar a hora do template de configuração
	if diff < 0 {
		diff += 24 * time.Hour
	}

	if diff < 8*time.Hour {
		return fmt.Errorf("o intervalo deve ser de no mínimo 8 horas, mas foi de %v", diff)
	}
	
	// t2, err := time.Parse("2006-01-02", inputTrip.StartDate)
	// if err != nil {
	// 	logger.Error("Error Parse Date", zap.Error(err))
	// 	return fmt.Errorf("error parse start_date: %w", err)
	// }

	tripDomain := &domain.TripConfig{
		ID:             tripID,
		ShipID:         ship.ID,
		RouteID:        route.ID,
		Recurrence:     domain.TripRecurrence(inputTrip.Recurrence),
		ExpirationDate: inputTrip.ExpirationDate,
		DepartureTime:  departureTime,
		ArrivalTime:    arrivalTime,
		DurationDays:   inputTrip.DurationDays,
		StartDate:      inputTrip.StartDate,
		CreatedAt:      time.Now(),
		UpdatedAt:      nil,
	}

	err = uc.tripRepository.CreateTripConfig(ctx, tripDomain)
	if err != nil {
		return fmt.Errorf("failed to create trip: %w", err)
	}
	logger.Info("package usecase", zap.String("trip_id", tripDomain.ID.String()))
	return nil
}
