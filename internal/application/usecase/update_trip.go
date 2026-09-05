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

type UpdateTripUseCase struct {
	tripRepo output.TripConfigRepository
}

func NewUpdateTripUseCase(tripRepo output.TripConfigRepository) *UpdateTripUseCase {
	return &UpdateTripUseCase{
		tripRepo: tripRepo,
	}
}

// Execute implements input.UpdateTripUseCase
func (uc *UpdateTripUseCase) Execute(ctx context.Context, id string, inputDTO *input.UpdateTripConfigInput) error {
	tripID, err := uuid.Parse(id)
	if tripID == uuid.Nil {
		return fmt.Errorf("invalid TripID: %w", err)
	}

	findTrip, err := uc.tripRepo.GetTripConfig(ctx, tripID)
	if err != nil {
		return fmt.Errorf("failed to find trip: %w", err)
	}
	if findTrip == nil {
		return fmt.Errorf("trip not found: %w", err)
	}
	shipUUID, err := uuid.Parse(inputDTO.ShipID)
	if err != nil {
		return fmt.Errorf("invalid ShipID: %w", err)
	}

	routeUUID, err := uuid.Parse(inputDTO.RouteID)
	if err != nil {
		return fmt.Errorf("invalid RouteID: %w", err)
	}

	tripDomain := domain.TripConfig{
		ID:             findTrip.ID,
		ShipID:         shipUUID,
		RouteID:        routeUUID,
		Recurrence:     domain.TripRecurrence(inputDTO.Recurrence),
		ExpirationDate: inputDTO.ExpirationDate,
		DurationDays:   inputDTO.DurationDays,
		UpdatedAt:      func() *time.Time { t := time.Now(); return &t }(),
	}

	err = uc.tripRepo.UpdateTripConfig(ctx, &tripDomain)
	if err != nil {
		return fmt.Errorf("failed to update trip: %w", err)
	}
	logger.Info("Trip updated successfully", zap.String("trip_id", tripDomain.ID.String()))
	return nil

}
