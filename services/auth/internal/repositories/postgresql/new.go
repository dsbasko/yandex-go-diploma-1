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

	if err = conn.Ping(ctx); err != nil {
		return nil, fmt.Errorf("conn.Ping: %w", err)
	}

	repo := &Repository{
		conn: conn,
	}

	if err = repo.TableInit(ctx); err != nil {
		return nil, fmt.Errorf("repo.TableInit: %w", err)
	}

	return &repo, nil
}

func (r *Repository) TableInit(ctx context.Context) error {
	if _, err := r.conn.Exec(ctx, `CREATE EXTENSION IF NOT EXISTS "uuid-ossp"`); err != nil {
		return fmt.Errorf("conn.Exec: create extension uuid: %w", err)
	}

	if _, err := r.conn.Exec(ctx, `
		CREATE TABLE IF NOT EXISTS accounts (
		    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
			username VARCHAR(255) UNIQUE,
			password VARCHAR(255),
		    last_login TIMESTAMPTZ,
		    created_at TIMESTAMPTZ DEFAULT NOW()
		)
	`); err != nil {
		return fmt.Errorf("conn.Exec: create table \"accounts\": %w", err)
	}

	if _, err := r.conn.Exec(ctx, `CREATE INDEX IF NOT EXISTS username ON accounts (username)`); err != nil {
		return fmt.Errorf("conn.Exec: create index \"username\" on table \"accounts\": %w", err)
	}

	return nil
}