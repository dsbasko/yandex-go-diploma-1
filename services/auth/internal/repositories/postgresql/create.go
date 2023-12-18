package postgresql

import (
	"context"
	"fmt"

	"github.com/dsbasko/yandex-go-diploma-1/services/auth/internal/domain"
	"github.com/dsbasko/yandex-go-diploma-1/services/auth/pkg/api"
)

func (r *Repository) CreateOnce(
	ctx context.Context,
	dto *api.RegisterRequestV1,
) (*domain.RepositoryAccountEntity, error) {
	query, args, err := r.builder.
		Insert("accounts").
		Columns("username", "password").
		Values(dto.Username, dto.Password).
		Suffix("RETURNING id, username, password").
		ToSql()

	if err != nil {
		return nil, fmt.Errorf("squirrel.ToSql: %w", err)
	}

	var response domain.RepositoryAccountEntity
	row := r.conn.QueryRow(ctx, query, args...)
	if err = row.Scan(
		&response.ID,
		&response.Username,
		&response.Password,
	); err != nil {
		return nil, fmt.Errorf("conn.QueryRow: row.Scan: %w", err)
	}

	return &response, nil
}
