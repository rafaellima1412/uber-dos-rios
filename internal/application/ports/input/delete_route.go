package input

import "context"

type DeleteRouteUseCase interface {
	Execute(ctx context.Context, id string) error
}
