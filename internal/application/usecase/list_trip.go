package usecase

import (
	"context"
	"time"

	"github.com/rafaellima1412/uber-dos-rios/internal/application/domain"
	"github.com/rafaellima1412/uber-dos-rios/internal/application/ports/input"
	"github.com/rafaellima1412/uber-dos-rios/internal/application/ports/output"
)

type ListTripConfigUseCase struct {
	tripRepo output.TripConfigRepository
}

func NewListTripConfigUseCase(tripRepo output.TripConfigRepository) input.ListTripConfigUseCase {
	return &ListTripConfigUseCase{tripRepo: tripRepo}
}

// Execute implements input.ListTripUseCase.
func (uc *ListTripConfigUseCase) Execute(ctx context.Context, inputTrip input.ListTripConfigInput) (*input.ListTripConfigOutput, error) {
	trips, err := uc.tripRepo.ListTripsConfig(ctx, inputTrip.PerPage, (inputTrip.Page-1)*inputTrip.PerPage)
	if err != nil {
		return nil, err
	}
	var summaries []input.TripConfigSummary
	for _, trip := range trips {
		var updatedAt time.Time
		if trip.UpdatedAt != nil {
			updatedAt = *trip.UpdatedAt
		}
		summaries = append(summaries, input.TripConfigSummary{
			ID:             trip.ID,
			ShipID:         trip.ShipID,
			RouteID:        trip.RouteID,
			ShipName:       trip.ShipName,
			RouteName:      trip.RouteName,
			Recurrence:     domain.TripRecurrence(trip.Recurrence),
			ExpirationDate: trip.ExpirationDate,
			DepartureTime:  trip.DepartureTime.Format("15:04"),
			ArrivalTime:    trip.ArrivalTime.Format("15:04"),
			StartDate:      trip.StartDate,
			DurationDays:   trip.DurationDays,
			CreatedAt:      trip.CreatedAt,
			UpdatedAt:      updatedAt,
		})
	}
	totalCount := len(trips) // For simplicity, assuming total count is the length of the current page
	totalpages := (totalCount + inputTrip.PerPage - 1) / inputTrip.PerPage
	return &input.ListTripConfigOutput{
		Trips:       summaries,
		TotalCount:  totalCount,
		TotalPages:  totalpages,
		CurrentPage: inputTrip.Page,
	}, nil
}
