package db

import (
	"context"
	"time"

	"github.com/Mohd-Sayeedul-Hoda/tinypath/internal/config"
	"github.com/jackc/pgx/v5/pgxpool"
)

func OpenDB(ctx context.Context, cfg *config.Config) (*pgxpool.Pool, error) {

	poolConfig, err := pgxpool.ParseConfig(cfg.DB.DSN)
	if err != nil {
		return nil, err
	}

	poolConfig.MaxConns = int32(cfg.DB.MaxOpenConns)
	poolConfig.MinConns = int32(cfg.DB.MaxIdleConns)

	duration, err := time.ParseDuration(cfg.DB.MaxIdleTime)
	if err != nil {
		return nil, err
	}
	poolConfig.MaxConnIdleTime = duration

	conn, err := pgxpool.NewWithConfig(ctx, poolConfig)
	if err != nil {
		return nil, err
	}

	err = conn.Ping(ctx)
	if err != nil {
		return nil, err
	}

	return conn, nil
}
