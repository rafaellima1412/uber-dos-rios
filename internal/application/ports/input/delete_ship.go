package input

import (
	"context"

	"github.com/google/uuid"
)

type DeleteShipUseCase interface {
	Execute(ctx context.Context, id  uuid.UUID) error
}