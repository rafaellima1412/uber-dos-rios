package usecase

import (
	"context"

	"github.com/rafaellima1412/uber-dos-rios/internal/application/ports/input"
	"github.com/rafaellima1412/uber-dos-rios/internal/application/ports/output"
)

type listShipCabinUseCase struct {
	cabinRepo output.ShipConfigRepository
}

func NewListCabinUseCase(cabinRepo output.ShipConfigRepository) input.ListShipCabinUseCase {
	return &listShipCabinUseCase{
		cabinRepo: cabinRepo,
	}
}

func (uc *listShipCabinUseCase) Execute(ctx context.Context,  inputCabin input.ListCabinInput) (input.ListCabinOutput, error) {

	limit := inputCabin.PerPage
	offset := (inputCabin.Page - 1) * inputCabin.PerPage

	cabins, err := uc.cabinRepo.ListCabins(ctx, limit, offset)
	if err != nil {
		return input.ListCabinOutput{}, err
	}
	outputCabins := make([]input.CabinOutput, 0, len(cabins))
	for _, sc := range cabins {
		outputCabins  = append(outputCabins, input.CabinOutput{
			ID:          sc.ID,
			Name:        sc.Name,
			BedType:     sc.BedType,
			Capacity:    sc.Capacity,
			Description: sc.Description,
		})
	}
	totalCount := len(cabins)
	totalPages := (totalCount + inputCabin.PerPage - 1) / inputCabin.PerPage

	return input.ListCabinOutput{
		Cabins:      outputCabins,
		TotalCount:  totalCount,
		TotalPages:  totalPages,
		CurrentPage: inputCabin.Page,
	}, nil
}
