package postgres

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/riolivre/nautical_logistics/internal/common"
	"github.com/riolivre/nautical_logistics/internal/config"
	"github.com/riolivre/nautical_logistics/internal/logger"
	"go.uber.org/zap"
)

func NewConnection(ctx context.Context, cfg *config.Config) *pgxpool.Pool {
	pool, err := pgxpool.New(ctx, cfg.Database.URL)
	if err != nil {
		logger.Fatal(common.FailedToCreateDBPool, zap.Error(err))
	}

	if err := pool.Ping(ctx); err != nil {
		logger.Fatal(common.FailedToCreateDBPool, zap.Error(err))
	}

	logger.Info(common.InfoConnectedToDB)

	return pool
}
