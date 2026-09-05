package output

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/riolivre/nautical_logistics/internal/application/domain"
)

// SeatRepository defines the interface for ship data persistence.
type ShipConfigRepository interface {
	CreateSeat(ctx context.Context, seat domain.Seat) (*domain.Seat, error)
	UpdateSeat(ctx context.Context, seat domain.Seat, id uuid.UUID) error
	ListSeats(ctx context.Context, limit int, offset int) ([]*domain.Seat, error)
	GetSeat(ctx context.Context, id uuid.UUID) (*domain.Seat, error)
	DeleteSeat(ctx context.Context, id uuid.UUID) error
	SearchSeats(ctx context.Context, isactive bool, prefix bool, limit int, offset int) ([]*domain.Seat, error)
	CreateCabins(ctx context.Context, cabin []domain.Cabin) error
	ListCabins(ctx context.Context, limit int, offset int) ([]*domain.CabinDetails, error)
	DeleteCabin(ctx context.Context, id uuid.UUID) error
	UpdateCabin(ctx context.Context, cabin domain.CabinDetails, id uuid.UUID) error
	GetCabin(ctx context.Context, id uuid.UUID) (*domain.CabinDetails, error)
	SearchSeatByShipAndDate(ctx context.Context, id uuid.UUID, dateOrigin, dateDest time.Time) (*domain.UnitsAvailable, error)
}
