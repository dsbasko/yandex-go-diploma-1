package postgresql

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/dsbasko/yandex-go-diploma-1/core/postgresql"
	"github.com/dsbasko/yandex-go-diploma-1/services/planner/internal/config"
	"github.com/dsbasko/yandex-go-diploma-1/services/planner/internal/interfaces"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct {
	conn    *pgxpool.Pool
	builder squirrel.StatementBuilderType
}

var _ interfaces.Repository = (*Repository)(nil)

func NewRepository(ctx context.Context) (*Repository, error) {
	conn, err := postgresql.Connect(ctx, config.GetPsqlConnectingString(), postgresql.Config{
		MaxConns: config.GetPsqlMaxPools(),
	})
	if err != nil {
		return nil, fmt.Errorf("postgresql.Connect: %w", err)
	}

	builder := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)

	repo := Repository{
		conn:    conn,
		builder: builder,
	}

	return &repo, nil
}
