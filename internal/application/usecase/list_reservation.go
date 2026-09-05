package usecase

import (
	"context"

	"github.com/rafaellima1412/uber-dos-rios/internal/application/ports/input"
	"github.com/rafaellima1412/uber-dos-rios/internal/application/ports/output"
)

type ListReservationUseCase struct {
	reservationRepo output.ReservationRepository
}

func NewListReservationUseCase(reservationRepo output.ReservationRepository) *ListReservationUseCase {
	return &ListReservationUseCase{
		reservationRepo: reservationRepo,
	}
}

func (uc *ListReservationUseCase) Execute(ctx context.Context, inputReservation input.ListReservationInput) (*input.ListReservationOutput, error) {

	page := inputReservation.Page
	if page < 1 {
		page = 1
	}

	perPage := inputReservation.PerPage
	if perPage <= 0 {
		perPage = 10
	}

	offset := (page - 1) * perPage

	reservations, err := uc.reservationRepo.ListReservations(ctx, perPage, offset)
	if err != nil {
		return nil, err
	}

	summaries := make([]input.ReservationSummary, 0, len(reservations))

	for _, reservation := range reservations {

		// configs := make([]input.ReservationTripConfigurationSummary, 0, len(reservation.TripInstance))

		// for _, tc := range reservation.TripInstance {
		// 	configs = append(configs, input.ReservationTripConfigurationSummary{
		// 		ID:                 tc.ReservationID.String(),
		// 		TripConfigurationID: tc.TripInstanceID.String(),
		// 		OccupiedSeats:       tc.OccupiedSeats,
		// 	})
		// }

		summaries = append(summaries, input.ReservationSummary{
			ID:                 reservation.ID.String(),
			UserID:             reservation.UserID.String(),
			// UserFullName:       reservation.UserFullName,
			// RouteID:            reservation.RouteID,
			// RouteName:          reservation.RouteName,
			// ShipID:             reservation.ShipID,
			// ShipName:           reservation.ShipName,
			// ReservationName:    reservation.ReservationName,
			ReservationDate:    reservation.ReservationDate,
			Status:             reservation.Status,
			CreatedAt:          reservation.CreatedAt,
			UpdatedAt:          reservation.UpdatedAt,
			// TripConfigurations: configs,
		})
	}

	totalCount := offset + len(reservations)
	totalPages := (totalCount + perPage - 1) / perPage

	return &input.ListReservationOutput{
		Reservations: summaries,
		TotalCount:   totalCount,
		TotalPages:   totalPages,
		CurrentPage:  page,
	}, nil
}
