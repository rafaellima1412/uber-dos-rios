package usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/rafaellima1412/uber-dos-rios/internal/application/ports/input"
	"github.com/rafaellima1412/uber-dos-rios/internal/application/ports/output"
	"github.com/rafaellima1412/uber-dos-rios/internal/logger"
	"go.uber.org/zap"
)

type GetTripUseCase struct {
	tripRepo output.TripConfigRepository
}

func NewGetTripUseCase(tripRepo output.TripConfigRepository) *GetTripUseCase {
	return &GetTripUseCase{
		tripRepo: tripRepo,
	}
}

// Execute implements input.GetTripUseCase
func (uc *GetTripUseCase) Execute(ctx context.Context, id string) (*input.GetTripConfigOutput, error) {
	tripID, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}
	trip, err := uc.tripRepo.GetTripConfig(ctx, tripID)
	if err != nil {
		logger.Error("error retrieving trip config", zap.Error(err))
		return nil, fmt.Errorf("error get trip config: %w", err)
	}
	if trip == nil {
		return nil, fmt.Errorf("trip not found")
	}
	var updatedAt time.Time
	if trip.UpdatedAt != nil {
		updatedAt = *trip.UpdatedAt
	}

	tripDTO := &input.GetTripConfigOutput{
		ID:             trip.ID,
		ShipID:         trip.ShipID,
		RouteID:        trip.RouteID,
		ShipName:       trip.ShipName,
		RouteName:      trip.RouteName,
		Recurrence:     string(trip.Recurrence),
		ExpirationDate: trip.ExpirationDate,
		DepartureTime:  trip.DepartureTime.String(),
		StartDate:      trip.StartDate.String(),
		DurationDays:   trip.DurationDays,
		CreatedAt:      trip.CreatedAt,
		UpdatedAt:      updatedAt,
	}
	logger.Info("Trip retrieved successfully", zap.String("trip_id", trip.ID.String()))

	return tripDTO, nil
}
