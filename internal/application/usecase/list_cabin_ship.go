package usecase

// import (
// 	"context"
// 	"fmt"

// 	"github.com/google/uuid"
// 	"github.com/rafaellima1412/uber-dos-rios/internal/application/ports/input"
// 	"github.com/rafaellima1412/uber-dos-rios/internal/application/ports/output"
// 	"github.com/rafaellima1412/uber-dos-rios/internal/logger"
// 	"go.uber.org/zap"
// )

// type getCabinShipUseCase struct {
// 	cabinShipRepo output.ShipConfigRepository
// }

// func NewGetCabinShipUseCase(cabinShipRepo output.ShipConfigRepository) *getCabinShipUseCase {
// 	return &getCabinShipUseCase{
// 		cabinShipRepo: cabinShipRepo,
// 	}
// }

// func (uc *getCabinShipUseCase) Execute(ctx context.Context, id string) (*input.GetShipCabinUseCase, error) {
// 	iduuID, err := uuid.Parse(id)
// 	if err != nil {
// 		return nil, err
// 	}
// 	cabin, err := uc.cabinShipRepo.ListCabinsByShipID(ctx, iduuID, 10, 0)
// 	if err != nil {
// 		logger.Error("Error retrieving cabin", zap.Error(err))
// 		return nil, fmt.Errorf("error get cabin: %w",err)
// 	}
// 	if cabin == nil {
// 		return nil, fmt.Errorf("cabin not found")
// 	}
// 	return nil,nil
// 	//  &input.GetShipCabinOutputDTO{
// 	// 	ID:              cabin.ID.String(),
// 	// 	ShipID:          cabin.ShipID.String(),
// 	// 	CabinID:         cabin.CabinID.String(),
// 	// 	Quantity:        cabin.Quantity,
// 	// 	Deck:            cabin.Deck,
// 	// 	Description:     cabin.Description,
// 	// 	CreatedAt:       cabin.CreatedAt,
// 	// 	UpdatedAt:       cabin.UpdatedAt,
// 	// }, 

// }