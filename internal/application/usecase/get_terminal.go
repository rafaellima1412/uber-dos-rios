package usecase

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/riolivre/nautical_logistics/internal/application/ports/input"
	"github.com/riolivre/nautical_logistics/internal/application/ports/output"
	"github.com/riolivre/nautical_logistics/internal/logger"
	"go.uber.org/zap"
)

// getTerminal is the implementation of GetTerminalUseCase.
type getTerminalUseCase struct {
	terminalRepo output.TerminalRepository
}

// Ensure getTerminal implements the GetTerminalUseCase interface.
var _ input.GetTerminalUseCase = (*getTerminalUseCase)(nil)

// NewGetTerminal creates a new instance of GetTerminalUseCase.
func NewGetTerminal(terminalRepo output.TerminalRepository) input.GetTerminalUseCase {
	return &getTerminalUseCase{terminalRepo: terminalRepo}
}

// Execute retrieves a terminal by its ID.
func (g *getTerminalUseCase) Execute(ctx context.Context, id string) (*input.GetTerminalOutput, error) {
	terminalID, err := uuid.Parse(id)
	if err != nil {
		logger.Error("Invalid terminal ID format", zap.String("id", id), zap.Error(err))
		return nil, fmt.Errorf("invalid terminal ID format: %w", err)
	}

	terminal, err := g.terminalRepo.GetTerminal(ctx, terminalID)
	if err != nil {
		logger.Error("Error retrieving terminal", zap.String("id", id), zap.Error(err))
		return nil, err
	}

	logger.Info("Terminal retrieved successfully", zap.String("terminal_id", terminal.ID.String()))

	return &input.GetTerminalOutput{
		ID:        terminal.ID,
		Name:      terminal.Name,
		UF:        terminal.UF,
		CityID:      terminal.CityID,
		Latitude:  terminal.Latitude,
		Longitude: terminal.Longitude,
		Active:    terminal.Activity,
	}, nil
}
