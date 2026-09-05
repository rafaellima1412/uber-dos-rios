// get seat
package input

import (
	"context"

	"github.com/google/uuid"
)

type	GetSeatOutput struct {
	ID                uuid.UUID
	IsActive          bool
	LeftColumnsCount  int
	RightColumnsCount int
	Nivel             int
	Prefix            bool
}

type GetSeatUseCase interface {
	Execute(ctx context.Context, id uuid.UUID) (*GetSeatOutput, error)
}	
