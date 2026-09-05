package usecase

import (
	"context"

	"github.com/rafaellima1412/uber-dos-rios/internal/application/ports/input"
	"github.com/rafaellima1412/uber-dos-rios/internal/application/ports/output"
)

type searchTerminalUseCase struct {
	terminalRepo output.TerminalRepository
}

// NewSearchTerminalUseCase creates a new instance of SearchTerminalUseCase.
func NewSearchTerminalUseCase(terminalRepo output.TerminalRepository) input.SearchTerminalUseCase {
	return &searchTerminalUseCase{
		terminalRepo: terminalRepo,
	}
}

// Execute searches terminals based on the provided input parameters.
func (uc *searchTerminalUseCase) Execute(ctx context.Context, query string, inputSearch input.SearchTerminalInput) (*input.SearchTerminalOutput, error) {
	limit := inputSearch.PerPage
	offset := (inputSearch.Page - 1) * inputSearch.PerPage

	terminals, err := uc.terminalRepo.SearchTerminals(ctx, query, limit, offset)
	if err != nil {
		return nil, err
	}

	var terminalSummaries []input.SearchTerminalSummary
	for _, terminal := range terminals {
		terminalSummaries = append(terminalSummaries, input.SearchTerminalSummary{
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

	return &input.SearchTerminalOutput{
		Terminals:   terminalSummaries,
		TotalCount:  totalCount,
		TotalPages:  totalPages,
		CurrentPage: inputSearch.Page,
	}, nil
}
