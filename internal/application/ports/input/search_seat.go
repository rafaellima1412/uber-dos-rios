package input

import (
	"context"
)

// SearchSeatlInput is the input data transfer object for searching trip.
type SearchSeatInput struct {
	Page    int
	PerPage int
}


// SearchSeatSummary is a summary representation of a trip.
type SearchSeatSummary struct {
	SeatID string
	Nivel  int
	Prefix bool
	IsActive bool	
	LeftColumnsCount  int
	RightColumnsCount int
}

// SearchSeatOutput is the output data transfer object for searching trips.
type SearchSeatOutput struct {
	Seats       []SearchSeatSummary
	TotalCount  int
	TotalPages  int
	CurrentPage int
}

// SearchTripUseCase defines the interface for the use case of searching trips.
type SearchSeatUseCase interface {
	Execute(ctx context.Context, isactive bool, prefix bool,inputSearch SearchSeatInput) (*SearchSeatOutput, error)
}

