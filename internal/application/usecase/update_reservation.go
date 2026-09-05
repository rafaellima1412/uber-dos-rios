package usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/rafaellima1412/uber-dos-rios/internal/application/domain"
	"github.com/rafaellima1412/uber-dos-rios/internal/application/ports/input"
	"github.com/rafaellima1412/uber-dos-rios/internal/application/ports/output"
	"github.com/rafaellima1412/uber-dos-rios/internal/logger"
	"go.uber.org/zap"
)

type updateReservationUseCase struct {
	reservationRepo output.ReservationRepository

}

func NewUpdateReservationUseCase(reservationRepo output.ReservationRepository) *updateReservationUseCase {
	return &updateReservationUseCase{
		reservationRepo: reservationRepo,
	}
}

// Execute implements input.UpdateReservationUseCase
func (uc *updateReservationUseCase) Execute(ctx context.Context, inputDTO *input.UpdateReservationInput, id uuid.UUID) error {
	reservation, err := uc.reservationRepo.GetReservation(ctx, id)
	if err != nil {
		return err
	}
	if reservation == nil {
		return fmt.Errorf("ship not found: %w", err)
	}

	reservationDomian := &domain.Reservation{
		ID:              reservation.ID,
		UserID:          reservation.UserID,
		ReservationDate: inputDTO.ReservationDate,
		Status:          domain.ReservationStatus(inputDTO.Status),
		UpdatedAt:       func() *time.Time { t := time.Now(); return &t }(),
	}
	
	err = uc.reservationRepo.UpdateReservation(ctx, reservationDomian)
	if err != nil {
		return err
	}
	logger.Info("Reservation updated successfully", zap.String("reservation_id",reservation.ID.String()))
	return nil
}
