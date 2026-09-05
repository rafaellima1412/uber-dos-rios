package usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/rafaellima1412/uber-dos-rios/internal/application/domain"
	"github.com/rafaellima1412/uber-dos-rios/internal/application/ports/input"
	"github.com/rafaellima1412/uber-dos-rios/internal/application/ports/output"
)

type UpdateCabinUseCase struct {
	cabinRepo output.ShipConfigRepository
}

func NewUpdateCabinUseCase(cabinRepo output.ShipConfigRepository) *UpdateCabinUseCase {
	return &UpdateCabinUseCase{
		cabinRepo: cabinRepo,
	}
}

// Execute implements input.UpdateCabinUseCase
func (uc *UpdateCabinUseCase) Execute(ctx context.Context, inputDTO *input.UpdateCabinInput, id uuid.UUID) error {
	findCabin, err := uc.cabinRepo.GetCabin(ctx, id)
	if err != nil {
		return fmt.Errorf("failed to find cabin: %w", err)
	}
	if findCabin == nil {
		return fmt.Errorf("cabin not found: %w", err)
	}

	cabinDomain := domain.CabinDetails{
		ID:          findCabin.ID,
		Name:        inputDTO.Name,
		BedType:     inputDTO.BedType,
		Capacity:    inputDTO.Capacity,
		Description: inputDTO.Description,
		UpdatedAt:   func() *time.Time { t := time.Now(); return &t }(),
	}

	err = uc.cabinRepo.UpdateCabin(ctx, cabinDomain, id)
	if err != nil {
		return fmt.Errorf("failed to update cabin: %w", err)
	}

	return nil

}