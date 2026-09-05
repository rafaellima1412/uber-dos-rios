package usecase

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	
	"github.com/rafaellima1412/uber-dos-rios/internal/application/ports/input"
	"github.com/rafaellima1412/uber-dos-rios/internal/application/ports/output"
	"github.com/rafaellima1412/uber-dos-rios/internal/logger"
	"go.uber.org/zap"
)

// getReservation
type getReservationUseCase struct {
	reeservationRepo output.ReservationRepository
}

func NewGetReservationUseCase(
	reeservationRepo output.ReservationRepository) *getReservationUseCase {
	return &getReservationUseCase{
		reeservationRepo: reeservationRepo,
	}
}
// Execute implements input.GetReservationUseCase
func (uc *getReservationUseCase) Execute(ctx context.Context, id string) (*input.GetReservationOutput, error) {
	iduuID, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}
	
	reservation, err := uc.reeservationRepo.GetReservation(ctx, iduuID)
	if err != nil {
		logger.Error("Error retrieving reservation", zap.Error(err))
		return nil, fmt.Errorf("error get reservation: %w",err)
	}
	if reservation == nil {
		return nil, fmt.Errorf("reservation not found")
	}
	return &input.GetReservationOutput{
		ID:              reservation.ID.String(),
		UserID:          reservation.UserID.String(),
		ReservationDate: reservation.ReservationDate.String(),
		Status:          string(reservation.Status),
	}, nil
}

