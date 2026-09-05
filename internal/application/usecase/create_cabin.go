package usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/riolivre/nautical_logistics/internal/application/domain"
	"github.com/riolivre/nautical_logistics/internal/application/ports/input"
	"github.com/riolivre/nautical_logistics/internal/application/ports/output"
	common "github.com/riolivre/nautical_logistics/internal/common"
	"github.com/riolivre/nautical_logistics/internal/logger"
	"go.uber.org/zap"
)

type CreateCabinUsecase struct {
	cabinrepo output.ShipConfigRepository
}

func NewCreateCabinUsecase(cabinrepo output.ShipConfigRepository) *CreateCabinUsecase {
	return &CreateCabinUsecase{
		cabinrepo: cabinrepo,
	}
}

func (uc *CreateCabinUsecase) Execute(ctx context.Context, inputDTO *input.CreateCabinInput) error {
	cabinsDomain := make([]domain.Cabin, 0, len(inputDTO.Cabins))
	for _, cabin := range inputDTO.Cabins {
		id, err := uuid.NewV7()
		if err != nil {
			logger.Error("Error generating UUID", zap.Error(err))
			return fmt.Errorf("error generating UUID: %w", err)
		}
		detail := domain.CabinDetails{
			ID:          id,
			Name:        cabin.Name,
			BedType:     cabin.BedType,
			Capacity:    cabin.Capacity,
			Description: cabin.Description,
			CreatedAt:   time.Now(),
		}
		cabinEntity := domain.Cabin{
        Cabins: []domain.CabinDetails{detail},
    }
		cabinsDomain = append(cabinsDomain, cabinEntity)
	}

	err := uc.cabinrepo.CreateCabins(ctx, cabinsDomain)
	if err != nil {
		return err
	}

	logger.Info(common.LogInfoCreateSeat)
	return nil

}
