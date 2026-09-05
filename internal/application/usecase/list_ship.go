package usecase

import (
	"context"

	"github.com/riolivre/nautical_logistics/internal/application/ports/input"
	"github.com/riolivre/nautical_logistics/internal/application/ports/output"
)

type ListShipUseCase struct {
	shipRepo output.ShipRepository
}

func NewListShipUseCase(shipRepo output.ShipRepository) input.ListShipUseCase {
	return &ListShipUseCase{shipRepo: shipRepo}
}

// Execute implements input.ListShipUseCase.
func (uc *ListShipUseCase) Execute(ctx context.Context, inputShip input.ListShipInput, typeship string) (input.ListShipOutput, error) {
	ships, err := uc.shipRepo.ListShips(ctx, inputShip.PerPage, (inputShip.Page-1)*inputShip.PerPage, typeship)
	if err != nil {
		return input.ListShipOutput{}, err
	}

	var summaries []input.ShipSummary
	for _, ship := range ships {
		summaries = append(summaries, input.ShipSummary{
			ID:                ship.ID,
			ConfigurationID:   ship.ConfigurationID,
			OrganizationID:    ship.OrganizationID,
			Name:              ship.Name,
			IMO:               ship.IMO,
			Status:            string(ship.StatusShip),
			TypeShip:          string(ship.TypeShip),
			PassengerCapacity: ship.PassengerCapacity,
			WeightCapacity:    ship.WeightCapacity,
			TotalSeats:        ship.TotalSeats,
			ImageUrl:          ship.ImageUrl,
		})
	}
	totalCount := len(ships) // For simplicity, assuming total count is the length of the current page
	totalpages := (totalCount + inputShip.PerPage - 1) / inputShip.PerPage

	return input.ListShipOutput{
		Ships:       summaries,
		TotalCount:  totalCount,
		TotalPages:  totalpages,
		CurrentPage: inputShip.Page,
	}, nil
}
