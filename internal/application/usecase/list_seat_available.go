// implements ListAvailableUseCase inteerface
package usecase

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/riolivre/nautical_logistics/internal/application/ports/input"
	"github.com/riolivre/nautical_logistics/internal/application/ports/output"
)

type listAvailableUseCase struct {
	shipRepo output.ShipConfigRepository
}

func NewListAvailableUseCase(shipRepo output.ShipConfigRepository) *listAvailableUseCase {
	return &listAvailableUseCase{
		shipRepo: shipRepo,
	}
}
func (uc *listAvailableUseCase) Execute(ctx context.Context, id uuid.UUID, dateOrigin, dateDest string) (*input.ListAvailableOutput, error) {
	dateOrigTime, err := time.Parse(time.RFC3339, dateOrigin)
	if err != nil {
		return nil, err
	}
	dateDestTime, err := time.Parse(time.RFC3339, dateDest)
	if err != nil {
		return nil, err
	}
	shipModel, err := uc.shipRepo.SearchSeatByShipAndDate(ctx, id, dateOrigTime, dateDestTime)
	if err != nil {
		return nil, err
	}
	seats := &input.Seats{

		TotalSeats:     shipModel.Seats.TotalSeats,
		OccupiedUnits:  shipModel.Seats.OccupiedUnits,
		AvailableSeats: shipModel.Seats.AvailableSeats,
	}
	cabins := &input.Cabins{
		TotalCabins:     shipModel.Cabins.TotalCabins,
		OccupiedUnits:   shipModel.Cabins.OccupiedUnits,
		AvailableCabins: shipModel.Cabins.AvailableCabins,
	}

	ship := &input.ListAvailableOutput{
		ShipID:   shipModel.ShipID,
		ShipName: shipModel.ShipName,
		Seats:    *seats,
		Cabins:   *cabins,
	}

	return ship, nil
}
