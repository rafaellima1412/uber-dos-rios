package usecase

import (
	"context"

	"github.com/riolivre/nautical_logistics/internal/application/ports/input"
	"github.com/riolivre/nautical_logistics/internal/application/ports/output"
)

type searchSeatsUseCase struct {
	seatRepo output.ShipConfigRepository
}

// NewSearchSeatsUseCase defines the interface for searching trips
func NewSearchSeatsUseCase(seatRepo output.ShipConfigRepository) input.SearchSeatUseCase {
	return &searchSeatsUseCase{
		seatRepo: seatRepo,
	}
}

// SearchSeatsUseCase defines the interface for searching trips
func (uc *searchSeatsUseCase) Execute(ctx context.Context, isactive bool, prefix bool, inputSearch input.SearchSeatInput) (*input.SearchSeatOutput, error) {
	limit := inputSearch.PerPage
	offset := (inputSearch.Page - 1) * inputSearch.PerPage
	seats, err := uc.seatRepo.SearchSeats(ctx, isactive, prefix, limit, offset)
	if err != nil {
		return nil, err
	}
	var seatSummaries []input.SearchSeatSummary
	for _, seat := range seats {
		seatSummaries = append(seatSummaries, input.SearchSeatSummary{
			SeatID:            seat.ID.String(),
			Prefix:            seat.Prefix,
			IsActive:          seat.IsActive,
			LeftColumnsCount:  seat.LeftColumnsCount,
			RightColumnsCount: seat.RightColumnsCount,
		})
	}
	// For simplicity, assuming total count is the length of the current page
	totalCount := len(seats)
	totalPages := (totalCount + limit - 1) / limit
	return &input.SearchSeatOutput{
		Seats:       seatSummaries,
		TotalCount:  totalCount,
		TotalPages:  totalPages,
		CurrentPage: inputSearch.Page,
	}, nil
}
