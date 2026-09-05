package input

import (
	"context"
)

// SearchTerminalInput is the input data transfer object for searching terminals.
type SearchTerminalInput struct {
	Page    int
	PerPage int
}

// SearchTerminalSummary is a summary representation of a terminal.
type SearchTerminalSummary struct {
	TerminalID string
	CityID     int
	CityName	 string
	UF         string
	Name       string
	Latitude   float64
	Longitude  float64
	Active     bool
}

// SearchTerminalOutput is the output data transfer object for searching terminals.
type SearchTerminalOutput struct {
	Terminals   []SearchTerminalSummary
	TotalCount  int
	TotalPages  int
	CurrentPage int
}

// SearchTerminalUseCase defines the interface for the use case of searching terminals.
type SearchTerminalUseCase interface {
	Execute(ctx context.Context, query string, inputSearch SearchTerminalInput) (*SearchTerminalOutput, error)
}
