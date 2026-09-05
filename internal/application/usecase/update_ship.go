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

type updateShipUseCase struct {
	shipRepo output.ShipRepository
}

func NewUpdateShipUseCase(shipRepo output.ShipRepository) *updateShipUseCase {
	return &updateShipUseCase{
		shipRepo: shipRepo,
	}
}

// Execute implements input.UpdateShipUseCase
func (uc *updateShipUseCase) Execute(ctx context.Context, inputDTO *input.UpdateShipInput, id uuid.UUID) error {
	findShip, err := uc.shipRepo.GetShip(ctx, id, "all")
	if err != nil {
		return fmt.Errorf(common.ErrGetShip, err)
	}
	if findShip == nil {
		return fmt.Errorf("ship not found: %w", err)
	}
	// configID := findShip.ConfigurationID
	// if inputDTO.ConfigurationsID != "" {
	// 	parsedUUID, err := uuid.Parse(inputDTO.ConfigurationsID)
	// 	if err != nil {
	// 		return fmt.Errorf("invalid configuration_id format: %w", err)
	// 	}
	// 	configID = &[]uuid.UUID{parsedUUID}
		
	// }

	shipDomain := &domain.Ship{
		ID:                findShip.ID,
		TypeShip:          domain.ShipType(inputDTO.TypeShip),
		StatusShip:        domain.ShipStatus(inputDTO.Status),
		Name:              inputDTO.Name,
		IMO:               inputDTO.IMO,
		TotalSeats:        inputDTO.TotalSeats,
		TotalCabins:       inputDTO.TotalCabins,
		PassengerCapacity: inputDTO.PassengerCapacity,
		WeightCapacity:    inputDTO.WeightCapacity,
		UpdatedAt:         func() *time.Time { t := time.Now(); return &t }(),
		//ConfigurationID:   configID,
		OrganizationID:    findShip.OrganizationID,
		ImageUrl:          findShip.ImageUrl,
	}

	err = uc.shipRepo.UpdateShip(ctx, shipDomain)
	if err != nil {
		return fmt.Errorf(common.ErrUpdateShip, err)
	}

	logger.Info(common.LogInfoUpdateShip, zap.String("ship_id", shipDomain.ID.String()))
	return nil
}
