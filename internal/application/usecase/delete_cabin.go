package usecase

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/rafaellima1412/uber-dos-rios/internal/application/ports/output"
	"github.com/rafaellima1412/uber-dos-rios/internal/logger"
	"go.uber.org/zap"
)

type DeleteCabinUsecase struct {
	cabinRepo output.ShipConfigRepository
}

func NewDeleteCabinUsecase(cabinRepo output.ShipConfigRepository) *DeleteCabinUsecase {
	return &DeleteCabinUsecase{
		cabinRepo: cabinRepo,
	}
}

func (uc *DeleteCabinUsecase) Execute(ctx context.Context, id uuid.UUID) error {
	if err := uc.cabinRepo.DeleteCabin(ctx, id); err != nil {
		logger.Error("Error deleting cabin", zap.Error(err))
		return fmt.Errorf("failed to delete cabin: %w", err)
	}
	return nil
}