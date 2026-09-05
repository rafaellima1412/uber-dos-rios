package input

import (
	"context"

	"github.com/google/uuid"
)

// update seat
type UpdateSeatInput struct {
	ID                uuid.UUID
	IsActive          bool
	LeftColumnsCount  int
	RightColumnsCount int
	Nivel             int
	Prefix            bool
}

type UpdateSeatOutput struct {
	ID                uuid.UUID
	IsActive          bool
	LeftColumnsCount  int
	RightColumnsCount int
	Nivel             int
	Prefix            bool
}

type UpdateSeatUseCase interface {
	Execute(ctx context.Context, input *UpdateSeatInput, id uuid.UUID)  error
}
