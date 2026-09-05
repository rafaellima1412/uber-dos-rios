package input

import (
	"context"

	"github.com/google/uuid"
)

type CreateSeatInput struct {
	IsActive          bool
	LeftColumnsCount  int
	RightColumnsCount int
	Prefix            bool
}

// DTO de saída após criar a Seat
type CreateSeatOutput struct {
	ID                uuid.UUID
	IsActive          bool
	LeftColumnsCount  int
	RightColumnsCount int
	Prefix            bool
}

type CreateSeatUseCase interface {
	Execute(ctx context.Context, input *CreateSeatInput) (*CreateSeatOutput, error)

}