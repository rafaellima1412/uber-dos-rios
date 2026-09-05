package usecase

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/rafaellima1412/uber-dos-rios/internal/application/ports/output"
	"github.com/rafaellima1412/uber-dos-rios/internal/logger"
	"go.uber.org/zap"
)

type DeleteSeatUseCase struct {
	seatRepo output.ShipConfigRepository
}

func NewDeleteSeatUseCase(seatRepo output.ShipConfigRepository) *DeleteSeatUseCase {
	return &DeleteSeatUseCase{
		seatRepo: seatRepo,
	}
}

func (uc *DeleteSeatUseCase) Execute(ctx context.Context, id uuid.UUID) error {
	if err := uc.seatRepo.DeleteSeat(ctx, id); err != nil {
		logger.Error("Error deleting seat", zap.Error(err))
		return fmt.Errorf("failed to delete seat: %w", err)
	}
	return nil

}
