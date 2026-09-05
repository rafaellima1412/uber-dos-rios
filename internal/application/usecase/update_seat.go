package usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/rafaellima1412/uber-dos-rios/internal/application/domain"
	"github.com/rafaellima1412/uber-dos-rios/internal/application/ports/input"
	"github.com/rafaellima1412/uber-dos-rios/internal/application/ports/output"
	common "github.com/rafaellima1412/uber-dos-rios/internal/common"
)

type UpdateSeatUseCase struct {
	seatRepo output.ShipConfigRepository
}

func NewUpdateSeatUseCase(seatRepo output.ShipConfigRepository) *UpdateSeatUseCase {
	return &UpdateSeatUseCase{
		seatRepo: seatRepo,
	}
}

// Execute implements input.UpdateShipUseCase
func (uc *UpdateSeatUseCase) Execute(ctx context.Context, inputDTO *input.UpdateSeatInput, id uuid.UUID) error{
	findSeat, err := uc.seatRepo.GetSeat(ctx, id)
	if err != nil {
		return fmt.Errorf("failed to find seat: %w", err)
	}
	if findSeat == nil {
		return fmt.Errorf("seat not found: %w", err)
	}

	seatDomain := domain.Seat{
		ID:                findSeat.ID,
		IsActive:          inputDTO.IsActive,
		LeftColumnsCount:  inputDTO.LeftColumnsCount,
		RightColumnsCount: inputDTO.RightColumnsCount,
		Prefix:            inputDTO.Prefix,
		UpdatedAt:         func() *time.Time { t := time.Now(); return &t }(),
	}

	err = uc.seatRepo.UpdateSeat(ctx, seatDomain, id)
	if err != nil {
		return fmt.Errorf(common.ErrUpdateSeat, err)
	}

	return nil

}
