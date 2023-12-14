package postgresql

import (
	"context"
	"fmt"

	"github.com/dsbasko/yandex-go-diploma-1/core/postgresql"
	"github.com/dsbasko/yandex-go-diploma-1/services/auth/internal/config"
	"github.com/dsbasko/yandex-go-diploma-1/services/auth/internal/domain"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct {
	conn *pgxpool.Pool
}

var _ domain.Repository = (*Repository)(nil)

func NewRepository(ctx context.Context) (domain.Repository, error) {
	conn, err := postgresql.Connect(ctx, config.GetPsqlConnectingString(), postgresql.Config{
		MaxConns: config.GetPsqlMaxPools(),
	})
	if err != nil {
		return nil, fmt.Errorf("postgresql.Connect: %w", err)
	}

	err = conn.Ping(ctx)
	if err != nil {
		return nil, fmt.Errorf("conn.Ping: %w", err)
	}

	return &Repository{
		conn: conn,
	}, nil
}
