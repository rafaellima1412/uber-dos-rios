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

type CreateReservationUseCase struct {
	reservationRepo output.ReservationRepository
	tripRepo        output.TripConfigRepository
}

func NewCreateReservationUseCase(reservationRepo output.ReservationRepository) *CreateReservationUseCase {
	return &CreateReservationUseCase{
		reservationRepo: reservationRepo,
	}
}

func (uc *CreateReservationUseCase) Execute(ctx context.Context, input *input.CreateReservationInput) (string, error) {

	reservationID, err := uuid.NewV7()
	if err != nil {
		return "", fmt.Errorf("error generating UUID: %w", err)
	}
	//criação do heeader da viagen
	reservation := &domain.Reservation{
		ID:              reservationID,
		UserID:          input.UserID,
		ReservationDate: input.ReservationDate,
		Status:          domain.ReservationStatus("PENDING"),
		CreatedAt:       time.Now(),
	}

	 id, err := uc.reservationRepo.CreateReservation(ctx, reservation)
	 if err != nil {
		return "", fmt.Errorf("failed to create reservation: %w", err)
	}

	logger.Info("Reservation created successfully", zap.String("reservation_id", reservationID.String()))

	return id, nil
}
