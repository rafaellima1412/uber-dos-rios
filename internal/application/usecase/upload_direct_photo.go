package usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/riolivre/nautical_logistics/internal/application/domain"
	"github.com/riolivre/nautical_logistics/internal/application/ports/input"
	"github.com/riolivre/nautical_logistics/internal/application/ports/output"
)

type UploadDirectMultipleUseCase struct {
	shipRepo output.ShipRepository // Repo: repository para salvar urls no banco
}

func NewUploadDirectMultipleUseCase(
	shipRepo output.ShipRepository,
) *UploadDirectMultipleUseCase {
	return &UploadDirectMultipleUseCase{
		shipRepo: shipRepo,
	}
}

// UploadDirectPhotos
func (uc *UploadDirectMultipleUseCase) Execute(ctx context.Context, photoInput *input.UploadShipImagesInput, shipID string) error {
	ID := uuid.MustParse(shipID)
	findShip, err := uc.shipRepo.GetShip(ctx, ID, "All")
	if err != nil {
		return fmt.Errorf("error get ship by id: %w", err)
	}
	if findShip == nil {
		return fmt.Errorf("ship not found: %w", err)
	}

	shipDomain := &domain.Ship{
		ID:                findShip.ID,
		TypeShip:          findShip.TypeShip,
		StatusShip:        findShip.StatusShip,
		Name:              findShip.Name,
		IMO:               findShip.IMO,
		TotalSeats:        findShip.TotalSeats,
		PassengerCapacity: findShip.PassengerCapacity,
		WeightCapacity:    findShip.WeightCapacity,
		UpdatedAt:         func() *time.Time { t := time.Now(); return &t }(),
		ConfigurationID:   findShip.ConfigurationID,
		OrganizationID:    findShip.OrganizationID,
		ImageUrl:          photoInput.Urls,
	}

	err = uc.shipRepo.UpdateShip(ctx, shipDomain)
	if err != nil {
		return fmt.Errorf("error update ship: %w", err)
	}

	return nil
}
