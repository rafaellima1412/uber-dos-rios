package output

import (
	"context"

	"github.com/google/uuid"
	"github.com/riolivre/nautical_logistics/internal/application/domain"
)
type GPSRepository interface {
	CreateGPS(ctx context.Context, gps *domain.GPS) error
	GetGPS(ctx context.Context, id uuid.UUID) (*domain.GPS, error)
}