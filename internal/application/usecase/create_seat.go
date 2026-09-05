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
	"github.com/rafaellima1412/uber-dos-rios/internal/logger"
	"go.uber.org/zap"
)

type CreateSeatUseCase struct {
	seatRepo output.ShipConfigRepository
}

func NewCreateSeatUseCase(seatRepo output.ShipConfigRepository) *CreateSeatUseCase {
	return &CreateSeatUseCase{
		seatRepo: seatRepo,
	}
}

func (uc *CreateSeatUseCase) Execute(ctx context.Context, inputDTO *input.CreateSeatInput) (*input.CreateSeatOutput, error) {
	seatDomain := domain.Seat{
		IsActive:          inputDTO.IsActive,
		LeftColumnsCount:  inputDTO.LeftColumnsCount,
		RightColumnsCount: inputDTO.RightColumnsCount,
		Prefix:            inputDTO.Prefix,
		CreatedAt:         time.Now(),
	}

	seatDomain.ID, _ = uuid.NewV7()
	createSeat, err := uc.seatRepo.CreateSeat(ctx, seatDomain)
	if err != nil {
		return nil, fmt.Errorf(common.ErrCreateSeat, err)
	}
	outputDTO := &input.CreateSeatOutput{
		ID:                createSeat.ID,
		IsActive:          createSeat.IsActive,
		LeftColumnsCount:  createSeat.LeftColumnsCount,
		RightColumnsCount: createSeat.RightColumnsCount,
		Prefix:            createSeat.Prefix,
	}

	logger.Info(common.LogInfoCreateSeat, zap.String("seat_id", outputDTO.ID.String()))
	return outputDTO, nil
}
