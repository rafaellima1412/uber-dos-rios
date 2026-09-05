package usecase

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/rafaellima1412/uber-dos-rios/internal/application/ports/input"
	"github.com/rafaellima1412/uber-dos-rios/internal/application/ports/output"
	"github.com/rafaellima1412/uber-dos-rios/internal/logger"
	"go.uber.org/zap"
)

type updateTerminalUseCase struct {
	terminalRepo output.TerminalRepository
}

// Ensure updateTerminalUseCase implements the UpdateTerminalUseCase interface.
var _ input.UpdateTerminalUseCase = (*updateTerminalUseCase)(nil)

// NewUpdateTerminal creates a new instance of UpdateTerminalUseCase.
func NewUpdateTerminal(terminalRepo output.TerminalRepository) input.UpdateTerminalUseCase {
	return &updateTerminalUseCase{terminalRepo: terminalRepo}
}

// Execute updates a terminal by its ID.
func (u *updateTerminalUseCase) Execute(ctx context.Context, id string, inputTerminal input.UpdateTerminalInput) error {
	terminalID, err := uuid.Parse(id)
	if err != nil {
		logger.Error("Invalid terminal ID format", zap.String("id", id), zap.Error(err))
		return fmt.Errorf("invalid terminal ID format: %w", err)
	}

	terminal, err := u.terminalRepo.GetTerminal(ctx, terminalID)
	if err != nil {
		logger.Error("Error retrieving terminal", zap.String("id", id), zap.Error(err))
		return fmt.Errorf("error retrieving terminal: %w", err)
	}

	terminal.Name = inputTerminal.Name
	terminal.UF = inputTerminal.UF
	terminal.CityID = inputTerminal.CityID
	terminal.Latitude = inputTerminal.Latitude
	terminal.Longitude = inputTerminal.Longitude

	if err := u.terminalRepo.UpdateTerminal(ctx, terminal); err != nil {
		logger.Error("Error updating terminal", zap.String("id", id), zap.Error(err))
		return fmt.Errorf("error updating terminal: %w", err)
	}

	logger.Info("Terminal updated successfully", zap.String("terminal_id", terminal.ID.String()))
	return nil
}
