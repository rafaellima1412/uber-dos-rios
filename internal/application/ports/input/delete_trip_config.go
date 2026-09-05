// implements the input port for deleting a trip.
package input

import (
	"context"
)

type DeleteTripConfigUseCase interface {
	Execute(ctx context.Context, id string) error
}	

