package usecase

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/riolivre/nautical_logistics/internal/application/ports/output"
	"github.com/riolivre/nautical_logistics/internal/logger"
	"go.uber.org/zap"
)

type DeleteReservationUseCase struct {
	reservationRepo output.ReservationRepository
}

func NewDeleteReservationUseCase(reservationRepo output.ReservationRepository) *DeleteReservationUseCase {
	return &DeleteReservationUseCase{
		reservationRepo: reservationRepo,
	}
}

func (uc *DeleteReservationUseCase) Execute(ctx context.Context, id uuid.UUID) error {
	if err := uc.reservationRepo.DeleteReservation(ctx, id); err != nil {
		logger.Error("Error deleting reservation", zap.Error(err))
		return fmt.Errorf("failed delete to reservation: %w",  err)

	}

	return nil
}
