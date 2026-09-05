package input

import "context"

type DeleteTerminalUseCase interface {
	Execute(ctx context.Context, id string) error
}
