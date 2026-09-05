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

type CreateShipUseCase struct {
	shipRepository output.ShipRepository
}

func NewCreateShipUseCase(shipRepo output.ShipRepository) *CreateShipUseCase {
	return &CreateShipUseCase{
		shipRepository: shipRepo,
	}
}

func (uc *CreateShipUseCase) Execute(ctx context.Context, inputDTO *input.CreateShipInput) error {
	organizationID, err := uuid.Parse(inputDTO.OrganizationID)
	if err != nil {
		return fmt.Errorf(common.ErrParseShipID, err)
	}

	// configurationsID, err := uuid.Parse(inputDTO.ConfigurationsID)
	// 	if err != nil {
	// 	return fmt.Errorf(common.ErrParseShipID, err)
	// }


	id, err := uuid.NewV7()
	if err != nil {
		logger.Error("Error generating UUID", zap.Error(err))
		return fmt.Errorf("error generating UUID: %w", err)
	}
	shipDomain := &domain.Ship{
		ID:                id,
		TypeShip:          domain.ShipType(inputDTO.TypeShip),
		StatusShip:        domain.ShipStatus(inputDTO.Status),
		Name:              inputDTO.Name,
		IMO:               inputDTO.IMO,
		TotalSeats:        inputDTO.TotalSeats,
		TotalCabins:       inputDTO.TotalCabins,
		PassengerCapacity: inputDTO.PassengerCapacity,
		WeightCapacity:    inputDTO.WeightCapacity,
		OrganizationID:    &organizationID,
		ConfigurationID:   &inputDTO.ConfigurationsID,
		CreatedAt:         time.Now(),
	}

	err = uc.shipRepository.CreateShip(ctx, shipDomain)
	if err != nil {
		return fmt.Errorf(common.ErrCreateShip, err)
	}

	logger.Info(common.LogInfoCreateShip, zap.String("ship_id", shipDomain.ID.String()))
	return nil
}
