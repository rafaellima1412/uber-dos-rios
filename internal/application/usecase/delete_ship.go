package usecase

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/riolivre/nautical_logistics/internal/application/ports/output"
	common "github.com/riolivre/nautical_logistics/internal/common"

	"github.com/riolivre/nautical_logistics/internal/logger"
	"go.uber.org/zap"
)

type DeleteShipUseCase struct {
	shipRepo output.ShipRepository
}

func NewDeleteShipUseCase(shipRepo output.ShipRepository) *DeleteShipUseCase {
	return &DeleteShipUseCase{
		shipRepo: shipRepo,
	}
}

// Execute implements input.DeleteShipUseCase
func (d *DeleteShipUseCase) Execute(ctx context.Context, id uuid.UUID) error {
	if err := d.shipRepo.DeleteShip(ctx, id); err != nil {
		logger.Error(common.LogFailedToDeleteShip, zap.Error(err))
		return fmt.Errorf("failed to delete ship: %w", err)
	}
	return nil

}
