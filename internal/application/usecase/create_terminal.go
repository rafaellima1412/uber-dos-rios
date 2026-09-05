package usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/riolivre/nautical_logistics/internal/application/domain"
	"github.com/riolivre/nautical_logistics/internal/application/ports/input"
	"github.com/riolivre/nautical_logistics/internal/application/ports/output"
	"github.com/riolivre/nautical_logistics/internal/logger"
	"go.uber.org/zap"
)

type createTerminal struct {
	terminalRepo output.TerminalRepository
}

var _ input.CreateTerminalUseCase = (*createTerminal)(nil)

func NewCreateTerminal(terminalRepo output.TerminalRepository) input.CreateTerminalUseCase {
	return &createTerminal{
		terminalRepo: terminalRepo,
	}
}

// Execute implements input.CreateTerminalUseCase.
func (c *createTerminal) Execute(ctx context.Context, input input.CreateTerminalInput) error {
	if input.Name == "" || input.UF == "" || input.CityID == 0 {
		logger.Error("Invalid input: Name, UF, and CityID are required fields")
		return fmt.Errorf("invalid input: Name, UF, and CityID are required fields")
	}

	id, err := uuid.NewV7()
	if err != nil {
		logger.Error("Error generating UUID", zap.Error(err))
		return fmt.Errorf("error generating UUID: %w", err)
	}

	activity := true

	terminal := &domain.Terminal{
		ID:        id,
		Name:      input.Name,
		UF:        input.UF,
		CityID:    input.CityID,
		Activity:  activity,
		Latitude:  input.Latitude,
		Longitude: input.Longitude,
		CreatedAt: time.Now(),
	}

	err = c.terminalRepo.CreateTerminal(ctx, terminal)
	if err != nil {
		logger.Error("Error creating terminal", zap.Error(err))
		return fmt.Errorf("error creating terminal: %w", err)
	}

	logger.Info("Terminal created successfully", zap.String("terminal_id", terminal.ID.String()))
	return nil
}
