package postgresql

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Config struct {
	MaxConns        int32
	MaxConnIdleTime time.Duration
}

const MaxRetries int = 10
const RetryTimeOut time.Duration = 3 * time.Second

func Connect(ctx context.Context, dsn string, options Config) (*pgxpool.Pool, error) {
	cfg, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return nil, fmt.Errorf("pgxpool.ParseConfig: %w", err)
	}

	cfg.MaxConns = options.MaxConns
	cfg.MaxConnIdleTime = options.MaxConnIdleTime

	connect, err := pgxpool.NewWithConfig(ctx, cfg)
	if err != nil {
		return nil, fmt.Errorf("pgxpool.NewWithConfig: %w", err)
	}

	for i := 0; i < MaxRetries; i++ {
		if err = connect.Ping(ctx); err != nil {
			if err == nil {
				break
			} else if i == MaxRetries-1 {
				return nil, fmt.Errorf("conn.Ping: %w", err)
			}
			time.Sleep(RetryTimeOut)
		}
	}

	return connect, nil
}
