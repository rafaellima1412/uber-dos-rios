package usecase

import (
	"context"

	"github.com/riolivre/nautical_logistics/internal/application/ports/input"
	"github.com/riolivre/nautical_logistics/internal/application/ports/output"
)

type listTerminalUseCase struct {
	terminalRepo output.TerminalRepository
}

// NewListTerminalUseCase creates a new instance of ListTerminalUseCase.
func NewListTerminalUseCase(terminalRepo output.TerminalRepository) input.ListTerminalUseCase {
	return &listTerminalUseCase{
		terminalRepo: terminalRepo,
	}
}

// Execute lists terminals based on the provided input parameters.
func (uc *listTerminalUseCase) Execute(ctx context.Context, inputTerminal input.ListTerminalInput) (*input.ListTerminalOutput, error) {
	limit := inputTerminal.PerPage
	offset := (inputTerminal.Page - 1) * inputTerminal.PerPage

	terminals, err := uc.terminalRepo.ListTerminals(ctx, limit, offset)
	if err != nil {
		return nil, err
	}

	var terminalSummaries []input.TerminalSummary
	for _, terminal := range terminals {
		terminalSummaries = append(terminalSummaries, input.TerminalSummary{
			TerminalID: terminal.ID.String(),
			CityID:     terminal.CityID,
			CityName: terminal.CityName,
			UF:         terminal.UF,
			Name:       terminal.Name,
			Latitude:   terminal.Latitude,
			Longitude:  terminal.Longitude,
			Active:     terminal.Activity,
		})
	}

	// For simplicity, assuming total count is the length of the current page
	totalCount := len(terminals)
	totalPages := (totalCount + limit - 1) / limit

	return &input.ListTerminalOutput{
		Terminals:   terminalSummaries,
		TotalCount:  totalCount,
		TotalPages:  totalPages,
		CurrentPage: inputTerminal.Page,
	}, nil
}
