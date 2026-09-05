package usecase

import (
	"context"

	"github.com/riolivre/nautical_logistics/internal/application/domain"
	"github.com/riolivre/nautical_logistics/internal/application/ports/input"
	"github.com/riolivre/nautical_logistics/internal/application/ports/output"
)

type SearchTrips struct {
	searchTriprepo output.TripRepository
}

func NewFilterTripUseCase(searchTriprepo output.TripRepository) *SearchTrips {
	return &SearchTrips{
		searchTriprepo: searchTriprepo,
	}
}

// Execute implements input.TripSearchUseCase
func (uc *SearchTrips) Execute(ctx context.Context, inputFilter *input.TripFilterInput, inputSearch input.SearchTripInput) (input.TripFilterOutput, error) {
	limit := inputSearch.PerPage
	offset := (inputSearch.Page - 1) * inputSearch.PerPage

	inputUsecase := &domain.TripFilter{
		ShipID:              inputFilter.ShipID,
		RouteID:             inputFilter.RouteID,
		TripConfigurationID: inputFilter.TripConfigurationsID,
		DepartureAfter:      inputFilter.DepartureAfter,
		DepartureBefore:     inputFilter.DepartureBefore,
	}

	trips, err := uc.searchTriprepo.FilterTrips(ctx, inputUsecase, limit, offset)
	if err != nil {
		return input.TripFilterOutput{}, err
	}
	var tripSummaries []input.TripSummary
	for _, trip := range trips {
		tripSummaries = append(tripSummaries, input.TripSummary{
			ID:                  trip.ID,
			ShipID:              trip.ShipID,
			RouteID:             trip.RouteID,
			TripConfigurationID: trip.TripConfigurationID,
			DepartureAt:         trip.DepartureAt,
			ArrivalAt:           trip.ArrivalAt,
			CreatedAt:           trip.CreatedAt,
			UpdatedAt:           trip.UpdatedAt,
		})
	}
	totalCount := len(trips)
	totalPages := (totalCount + limit - 1) / limit
	return input.TripFilterOutput{
		Trips:       tripSummaries,
		TotalCount:  totalCount,
		TotalPages:  totalPages,
		CurrentPage: inputSearch.Page,
	}, nil

}
