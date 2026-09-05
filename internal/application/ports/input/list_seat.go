package input

import (
	"context"
)	
import (
	"github.com/google/uuid"
)
type ListSeatInput struct {
	Page  int
	PerPage int
}	
type SeatSummary struct {
	ID               uuid.UUID
	IsActive         bool
	LeftColumnsCount int
	RightColumnsCount int
	Nivel            int	
	Prefix           bool
}	
type ListSeatOutput struct {
	Seats      []SeatSummary
	TotalCount int
	TotalPages int	
	CurrentPage int
}
type ListSeatUseCase interface {
	Execute(ctx context.Context, input ListSeatInput) (ListSeatOutput, error)
}