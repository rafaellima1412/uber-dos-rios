package usecase

//listSeatUseCase.go
import (
	"context"

	"github.com/rafaellima1412/uber-dos-rios/internal/application/ports/input"
	"github.com/rafaellima1412/uber-dos-rios/internal/application/ports/output"
)

type listSeatUseCase struct {
	seatRepo output.ShipConfigRepository
}

// NewListSeatUseCase creates a new instance of ListSeatUseCase.
func NewListSeatUseCase(seatRepo output.ShipConfigRepository) input.ListSeatUseCase {
	return &listSeatUseCase{seatRepo: seatRepo}
}

// Execute lists seats based on the provided input parameters.
func (uc *listSeatUseCase) Execute(ctx context.Context, inputSeat input.ListSeatInput) (input.ListSeatOutput, error) {

	limit := inputSeat.PerPage
	offset := (inputSeat.Page - 1) * inputSeat.PerPage
	seats, err := uc.seatRepo.ListSeats(ctx, limit, offset)
	if err != nil {
		return input.ListSeatOutput{}, err
	}

	seatSummaries := make([]input.SeatSummary, 0)
	for _, seat := range seats {
		seatSummaries = append(seatSummaries, input.SeatSummary{
			ID:                seat.ID,
			IsActive:          seat.IsActive,
			LeftColumnsCount:  seat.LeftColumnsCount,
			RightColumnsCount: seat.RightColumnsCount,
			Prefix:            seat.Prefix,
		})
	}
	totalCount := len(seats)
	totalPages := (totalCount + inputSeat.PerPage - 1) / inputSeat.PerPage
	return input.ListSeatOutput{
		Seats:       seatSummaries,
		TotalCount:  totalCount,
		TotalPages:  totalPages,
		CurrentPage: inputSeat.Page,
	}, nil
}
