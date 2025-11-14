package postgres

import (
	"context"
	"log/slog"

	"github.com/alonsoF100/shotlink/internal/config"
	"github.com/jackc/pgx/v5/pgxpool"
)

func NewPool(ctx context.Context, cfg config.DatabaseConfig) (*pgxpool.Pool, error) {
	connString := cfg.ConStr()

	poolConfig, err := pgxpool.ParseConfig(connString)
	if err != nil {
		slog.Error("failed to parse pool config", "connString", connString, "error", err)
		return nil, err
	}

	poolConfig.MaxConns = int32(cfg.MaxOpenConns)
	poolConfig.MinConns = int32(cfg.MaxIdleConns)
	poolConfig.MaxConnLifetime = cfg.ConnMaxLifetime

	pool, err := pgxpool.NewWithConfig(ctx, poolConfig)
	if err != nil {
		slog.Error("failed to connect", "connString", connString, "error", err)
		return nil, err
	}

	if err := pool.Ping(ctx); err != nil {
		slog.Error("failed to ping", "connString", connString, "error", err)
		return nil, err
	}

	slog.Info("Successfully connected to PostgreSQL",
		"max_conns", cfg.MaxOpenConns,
		"min_conns", cfg.MaxIdleConns,
		"conn_life_time", cfg.ConnMaxLifetime,
	)
	return pool, nil
}
