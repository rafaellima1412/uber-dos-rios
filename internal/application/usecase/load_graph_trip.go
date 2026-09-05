package usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/riolivre/nautical_logistics/internal/application/ports/input"
	"github.com/riolivre/nautical_logistics/internal/application/ports/output"
)

type LoadGraphTripUseCase struct {
	tripRepo output.TripConfigRepository
}

func NewLoadGraphTripUseCase(tripRepo output.TripConfigRepository) *LoadGraphTripUseCase {
	return &LoadGraphTripUseCase{
		tripRepo: tripRepo,
	}
}

func (uc *LoadGraphTripUseCase) Execute(ctx context.Context, DateOrigin string, DateDest string) (input.Graph, error) {

	start, err := time.Parse(time.RFC3339, DateOrigin)
	if err != nil {
		return nil, fmt.Errorf("invalid date_origin format, expected YYYY-MM-DD")
	}
	formatted_start := start.Format("2006-01-02")

	end,  err := time.Parse(time.RFC3339,  DateDest)
	if err != nil {
		return nil,fmt.Errorf("invalid date_dest format, expected YYYY-MM-DD")
	}
	formatted_end := end.Format("2006-01-02")

	return uc.tripRepo.LoadGraph(ctx, formatted_start, formatted_end)
}
