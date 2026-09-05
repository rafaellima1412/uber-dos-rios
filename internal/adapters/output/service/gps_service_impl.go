package service

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/riolivre/nautical_logistics/internal/adapters/input/http/dto"
	"github.com/riolivre/nautical_logistics/internal/application/domain"
	"github.com/riolivre/nautical_logistics/internal/application/ports/output"
)

type GPSService struct {
	repo output.GPSRepository
}

func NewGPSService(repo output.GPSRepository) *GPSService {
	return &GPSService{
		repo: repo,
	}
}

func (s *GPSService) ProcessGPSData(
	ctx context.Context,
	data dto.GPSRequest,
) error {

	if data.ShipName == "" {
		return fmt.Errorf("ship_name obrigatório")
	}

	gps := &domain.GPS{
		ID:           uuid.New(),
		ShipName:     data.ShipName,
		Latitude:     data.Latitude,
		Longitude:    data.Longitude,
		Speed:        data.Speed,
		Course:       data.Course,
		ReceivedAt:   time.Now().UTC(),
	}

	return s.repo.CreateGPS(ctx, gps)
}
