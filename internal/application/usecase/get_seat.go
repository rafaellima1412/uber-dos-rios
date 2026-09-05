// get seat
package usecase

import (
	"context"

	"github.com/google/uuid"
	"github.com/rafaellima1412/uber-dos-rios/internal/application/ports/input"
	"github.com/rafaellima1412/uber-dos-rios/internal/application/ports/output"
)

type GetSeatUseCase struct {
	seatRepo output.ShipConfigRepository
}

func NewGetSeatUseCase(seatRepo output.ShipConfigRepository) *GetSeatUseCase {
	return &GetSeatUseCase{	
		seatRepo: seatRepo,
	}
}

func (uc *GetSeatUseCase) Execute(ctx context.Context, id uuid.UUID) (*input.GetSeatOutput, error) {
	seat, err := uc.seatRepo.GetSeat(ctx, id)
	if err != nil {
		return nil, err
	}
	return &input.GetSeatOutput{
		ID:                seat.ID,
		IsActive:          seat.IsActive,
		LeftColumnsCount:  seat.LeftColumnsCount,
		RightColumnsCount: seat.RightColumnsCount,
	}, nil
}