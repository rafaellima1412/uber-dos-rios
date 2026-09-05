package input

import "context"

type DeleteScheduleUseCase interface {
	Execute(ctx context.Context, id string) error
}
