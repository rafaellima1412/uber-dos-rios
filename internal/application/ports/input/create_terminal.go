package input

import (
	"context"

	"github.com/google/uuid"
)

// CreateTerminalInput represents the input data required to create a new terminal.
type CreateTerminalInput struct {
	Name      string
	UF        string
	CityID    int
	Latitude  float64
	Longitude float64
}

// CreateTerminalOutput represents the output data after creating a new terminal.
type CreateTerminalOutput struct {
	ID   uuid.UUID
	Name string
}

// CreateTerminalUseCase defines the interface for the use case of creating a new terminal.
type CreateTerminalUseCase interface {
	Execute(ctx context.Context, input CreateTerminalInput) error
}
