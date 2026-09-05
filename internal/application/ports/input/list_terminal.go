package input

import "context"

// ListTerminalInput is the input data transfer object for listing terminals.
type ListTerminalInput struct {
	Page    int
	PerPage int
}

// TerminalSummary is a summary representation of a terminal.
type TerminalSummary struct {
	TerminalID string
	CityID     int
	CityName   string
	UF         string
	Name       string
	Latitude   float64
	Longitude  float64
	Active     bool
}

// ListTerminalOutput is the output data transfer object for listing terminals.
type ListTerminalOutput struct {
	Terminals   []TerminalSummary
	TotalCount  int
	TotalPages  int
	CurrentPage int
}

// ListTerminalUseCase defines the interface for the use case of listing terminals.
type ListTerminalUseCase interface {
	Execute(ctx context.Context, input ListTerminalInput) (*ListTerminalOutput, error)
}
