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

// deleteTerminalUseCase is the implementation of DeleteTerminalUseCase.
type deleteTerminalUseCase struct {
	terminalRepo output.TerminalRepository
}

// Ensure deleteTerminalUseCase implements the DeleteTerminalUseCase interface.
var _ input.DeleteTerminalUseCase = (*deleteTerminalUseCase)(nil)

// NewDeleteTerminalUseCase creates a new instance of DeleteTerminalUseCase.
func NewDeleteTerminalUseCase(terminalRepo output.TerminalRepository) input.DeleteTerminalUseCase {
	return &deleteTerminalUseCase{
		terminalRepo: terminalRepo,
	}
}

// Execute implements input.DeleteTerminalUseCase
func (d *deleteTerminalUseCase) Execute(ctx context.Context, id string) error {
	// Parse the terminal ID
	terminalID, err := uuid.Parse(id)
	if err != nil {
		logger.Error("Invalid terminal ID format", zap.String("id", id), zap.Error(err))
		return fmt.Errorf("invalid terminal ID format: %w", err)
	}

	// Call the repository to delete the terminal
	if err := d.terminalRepo.DeleteTerminal(ctx, terminalID); err != nil {
		logger.Error("Failed to delete terminal", zap.String("id", id), zap.Error(err))
		return fmt.Errorf("failed to delete terminal: %w", err)
	}

	logger.Info("Terminal deleted successfully", zap.String("terminal_id", id))
	return nil
}
