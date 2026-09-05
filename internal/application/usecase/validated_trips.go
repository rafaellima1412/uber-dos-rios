package usecase

import (
	"context"
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/rafaellima1412/uber-dos-rios/internal/application/ports/output"
	common "github.com/rafaellima1412/uber-dos-rios/internal/common"
	"github.com/rafaellima1412/uber-dos-rios/internal/logger"
	"go.uber.org/zap"
)

type ValidatedTripsUseCase struct {
	shipRepo       output.ShipRepository
	shipConfigRepo output.ShipConfigRepository
}

func NewValidatedTripsUseCase(shipRepo output.ShipRepository, shipConfigRepo output.ShipConfigRepository) *ValidatedTripsUseCase {
	return &ValidatedTripsUseCase{
		shipRepo:       shipRepo,
		shipConfigRepo: shipConfigRepo,
	}
}

// regex para unidades "A3, CABIN-2"
var unitRegex = regexp.MustCompile(`^([A-Z][1-9]\d*|CA[1-9]\d*)$`)

func isValidUnitCode(code string) bool {
	code = strings.ToUpper(code)
	return unitRegex.MatchString(code)
}

func (s *ValidatedTripsUseCase) Execute(ctx context.Context,
	departure, arrival string, occupiedUnits []string, id uuid.UUID) (*time.Time, *time.Time, error) {
	// Parse strings into time.Time objects
	layout := time.RFC3339
	depTime, err := time.Parse(layout, departure)
	if err != nil {
		return nil, nil, err
	}
	arrTime, err := time.Parse(layout, arrival)
	if err != nil {
		return nil, nil, err
	}

	//---------validação do payload------------------
	if len(occupiedUnits) == 0 {
		return nil, nil, fmt.Errorf("occupied_units is required for trip_configuration")
	}

	seen := map[string]struct{}{}
	for _, seat := range occupiedUnits {
		if _, ok := seen[seat]; ok {
			return nil, nil, fmt.Errorf("duplicate unit %s in same trip", seat)
		}
		seen[seat] = struct{}{}
	}

	for _, seat := range occupiedUnits {
		if !isValidUnitCode(seat) {
			return nil, nil, fmt.Errorf("invalid unit code: %s", seat)
		}
	}
	//-------------------------------------------------------
	// ------------- Validação datas -----------------
	//não pode sair antes de chegar
	if !arrTime.After(depTime) {
		//return fmt.Errorf(common.ErrArrivalBeforeDeparture)
		logger.Info(common.ErrArrivalBeforeDeparture)

		// TODO ajustar conforme a recorencia ou quantidade de dias ....
		arrTime = arrTime.AddDate(0, 0, 1) // Ao inves de retornar erro ajusto a data
		logger.Info("Ajustando chegada para o dia seguinte", zap.Time("nova_chegada", arrTime))
	}

	//timezone
	if depTime.Location() != arrTime.Location() {
		logger.Info(common.ErrTimezoneMismatch, zap.String("departure_timezone", depTime.Location().String()),
			zap.String("arrival_timezone", arrTime.Location().String()))
	}

	// Validação  de duração
	duration := arrTime.UTC().Sub(depTime.UTC())
	minDuration := 8 * time.Hour
	if duration < minDuration {

		y1, m1, d1 := depTime.Date()
		arrivalInStartLoc := arrTime.In(depTime.Location())
		y2, m2, d2 := arrivalInStartLoc.Date()

		if y1 == y2 && m1 == m2 && d1 == d2 {
			return nil, nil, fmt.Errorf(common.ErrTripTooShortSameDay)
		}
		return nil, nil, fmt.Errorf(common.ErrTripTooShort)
	}
	//------------- Validação com Banco ----------------------
	//total de assentos e cabines
	ship, err := s.shipRepo.GetShip(ctx, id, "all")
	if err != nil {
		return nil, nil, err
	}
	if ship == nil {
		return nil, nil, fmt.Errorf("ship not found")
	}

	requestedCount := len(occupiedUnits)

	if requestedCount > ship.TotalSeats {
		return nil, nil, fmt.Errorf("passenger capacity exceeded: requested %d but only %d left",
			requestedCount, ship.TotalSeats)
	}

	resp, _ := s.shipConfigRepo.SearchSeatByShipAndDate(ctx, id, depTime, arrTime)
	if resp != nil {
		arrayOcupados := resp.Seats.OccupiedUnits
		//Comparar os assentos do payload com o do banco
		alreadyOccupied := make(map[string]struct{})
		for _, seat := range arrayOcupados {
			alreadyOccupied[seat] = struct{}{}
		}
		for _, seat := range occupiedUnits {
			// Se unidade existir no  banco
			if _, exists := alreadyOccupied[strings.ToUpper(seat)]; exists {
				return nil, nil, fmt.Errorf("unit %s is already occupied in the database for this trip", seat)
			}
		}
	}

	return &depTime, &arrTime, nil
}
