package output

import (
	"context"

	"github.com/google/uuid"
	"github.com/riolivre/nautical_logistics/internal/application/domain"
)

// TerminalRepository defines the interface for interacting with the terminal repository.
type TerminalRepository interface {
	CreateTerminal(ctx context.Context, terminal *domain.Terminal) error
	GetTerminal(ctx context.Context, id uuid.UUID) (*domain.Terminal, error)
	UpdateTerminal(ctx context.Context, terminal *domain.Terminal) error
	DeleteTerminal(ctx context.Context, id uuid.UUID) error
	ListTerminals(ctx context.Context, limit int, offset int) ([]*domain.Terminal, error)
	SearchTerminals(ctx context.Context, query string, limit int, offset int) ([]*domain.Terminal, error)
	ListCities(ctx context.Context) ([]*domain.City, error)
}
