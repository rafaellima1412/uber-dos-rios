package usecase

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/rafaellima1412/uber-dos-rios/internal/application/ports/input"
	"github.com/rafaellima1412/uber-dos-rios/internal/application/ports/output"
	"github.com/rafaellima1412/uber-dos-rios/internal/common"
	"github.com/rafaellima1412/uber-dos-rios/internal/logger"
	"go.uber.org/zap"
)

type GetShipUseCase struct {
	shipRepo output.ShipRepository
}

func NewGetShipUseCase(shipRepo output.ShipRepository) *GetShipUseCase {
	return &GetShipUseCase{shipRepo: shipRepo}
}

// Execute implements input.GetShipUseCase.
func (uc *GetShipUseCase) Execute(ctx context.Context, id uuid.UUID, typeShip string) (*input.GetShipOutput, error) {
	ship, err := uc.shipRepo.GetShip(ctx, id, typeShip)
	if err != nil {
		logger.Error(common.LogFailedToGetShip, zap.Error(err))
		return nil, fmt.Errorf(common.ErrGetShip, err)
	}

	if ship == nil {
		return nil, fmt.Errorf("ship not found")
	}

	return &input.GetShipOutput{
		ID:                ship.ID,
		ConfigurationID:   ship.ConfigurationID,
		OrganizationID:    ship.OrganizationID,
		Name:              ship.Name,
		IMO:               ship.IMO,
		Status:            string(ship.StatusShip),
		TotalSeats:        ship.TotalSeats,
		PassengerCapacity: ship.PassengerCapacity,
		WeightCapacity:    ship.WeightCapacity,
		TypeShip:          string(ship.TypeShip),
		ImageUrl:          ship.ImageUrl,
	}, nil
}
