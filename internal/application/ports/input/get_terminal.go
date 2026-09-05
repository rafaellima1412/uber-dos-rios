package input

import (
	"context"

	"github.com/google/uuid"
)

// GetTerminalOutput represents the output data after retrieving a terminal.
type GetTerminalOutput struct {
	ID        uuid.UUID
	Name      string
	UF        string
	CityID    int
	Latitude  float64
	Longitude float64
	Active    bool
}

// GetTerminalUseCase defines the interface for the use case to get a terminal by ID.
type GetTerminalUseCase interface {
	Execute(ctx context.Context, id string) (*GetTerminalOutput, error)
}
