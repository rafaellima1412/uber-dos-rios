package input

import "context"

// UpdateTerminalInput represents the input data required to update a terminal.
type UpdateTerminalInput struct {
	ID        string
	Name      string
	UF        string
	CityID    int
	Latitude  float64
	Longitude float64
}

// UpdateTerminalOutput represents the output data after updating a terminal.
type UpdateTerminalOutput struct {
	ID   string
	Name string
}

type UpdateTerminalUseCase interface {
	Execute(ctx context.Context, id string, inputTerminal UpdateTerminalInput) error
}
