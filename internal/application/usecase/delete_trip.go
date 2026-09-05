package usecase

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/riolivre/nautical_logistics/internal/application/ports/input"
	"github.com/riolivre/nautical_logistics/internal/application/ports/output"
	"github.com/riolivre/nautical_logistics/internal/logger"
	"go.uber.org/zap"
)
type DeleteTripConfigUseCase struct {
	tripConfigRepo output.TripConfigRepository
}
var _ input.DeleteTripConfigUseCase = (*DeleteTripConfigUseCase)(nil)

// NewDeleteTripUseCase creates a new instance of DeleteTripUseCase.
func NewDeleteTripUseCase(tripConfigRepo output.TripConfigRepository) input.DeleteTripConfigUseCase {
	return &DeleteTripConfigUseCase{tripConfigRepo: tripConfigRepo}
}

// Execute implements input.DeleteTripUseCase.
func (uc *DeleteTripConfigUseCase) Execute(ctx context.Context, id string) error {
	// Parse the terminal ID
	tripID, err := uuid.Parse(id)
	if err != nil {
		logger.Error("Invalid trip ID format", zap.String("id", id), zap.Error(err))
		return fmt.Errorf("invalid trip ID format: %w", err)
	}
	// Call the repository to delete the trip
	if err := uc.tripConfigRepo.DeleteTripConfig(ctx, tripID); err != nil {
		logger.Error("Failed to delete trip config", zap.String("id", id), zap.Error(err))
		return fmt.Errorf("failed to delete trip config: %w", err)
	}
	logger.Info("Trip deleted successfully", zap.String("trip_id", id))
	return nil
}
