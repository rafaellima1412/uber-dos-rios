package usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/riolivre/nautical_logistics/internal/application/domain"
	"github.com/riolivre/nautical_logistics/internal/application/ports/input"
	"github.com/riolivre/nautical_logistics/internal/application/ports/output"
	"github.com/riolivre/nautical_logistics/internal/logger"
	"go.uber.org/zap"
)

type addTripsToReservationUseCase struct {
	reservationRepo output.ReservationRepository
	validateTripsUC input.ValidatedTripsUseCase 
}

func NewAddTripsToReservationUseCase(reservationRepo output.ReservationRepository, validateUC input.ValidatedTripsUseCase) *addTripsToReservationUseCase {
	return &addTripsToReservationUseCase{
		reservationRepo: reservationRepo,
		validateTripsUC: validateUC,
	}
}

func (uc *addTripsToReservationUseCase) Execute(
	ctx context.Context,
	triToReservation *input.TripToReservationInput,
) error {

	for _, trip := range triToReservation.SelectedTrips {
		
		departureStr := trip.DepartureAt.Format(time.RFC3339)
		arrivalStr := trip.ArrivalAt.Format(time.RFC3339)

		deptime, arrivaltime, err := uc.validateTripsUC.Execute(
			ctx,
			departureStr,
			arrivalStr,
			trip.OccupiedUnits,
			trip.ShipID,
		)
		if err != nil {
			return err
		}

		tripID, err := uuid.NewV7()
		if err != nil {
			return fmt.Errorf("error generating UUID: %w", err)
		}

		trips := &domain.Trip{
			ID:                  tripID,
			TripConfigurationID: trip.TripConfigurationID,
			ShipID:              trip.ShipID,
			RouteID:             trip.RouteID,
			DepartureAt:         *deptime,
			ArrivalAt:           *arrivaltime,
			CreatedAt:           time.Now(),
		}

		tripToReservation := &domain.TripToReservation{
			ReservationID:  triToReservation.ReservationID,
			TripInstanceID: tripID,
			OccupiedUnits:  trip.OccupiedUnits,
		}
		

		if err := uc.reservationRepo.CreateTripReservation(ctx, trips, tripToReservation); 
		err != nil {
			return err
		}
		
		logger.Info("Trip instance created successfully", zap.String("trip_instance_id", tripID.String()))
	}
	return nil

}
