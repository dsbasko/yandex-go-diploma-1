package postgresql

import (
	"context"
	"fmt"

	"github.com/dsbasko/yandex-go-diploma-1/core/lib"
	"github.com/dsbasko/yandex-go-diploma-1/services/auth/internal/domain"
)

func (r *Repository) FindByID(
	ctx context.Context,
	id string,
) (*domain.RepositoryAccountEntity, error) {
	query, args, err := r.builder.
		Select(lib.StructToKeysAndValues(
			&domain.RepositoryAccountEntity{}, false, false,
		).Keys...).
		From("accounts").
		Where("id = ?", id).
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
		&response.FirstName,
		&response.LastName,
		&response.LastLogin,
		&response.CreatedAt,
		&response.UpdatedAt,
	); err != nil {
		return nil, fmt.Errorf("conn.QueryRow: row.Scan: %w", err)
	}

	return &response, nil
}

func (r *Repository) FindByUsername(
	ctx context.Context,
	username string,
) (*domain.RepositoryAccountEntity, error) {
	query, args, err := r.builder.
		Select(lib.StructToKeysAndValues(
			&domain.RepositoryAccountEntity{}, false, false,
		).Keys...).
		From("accounts").
		Where("username = ?", username).
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
		&response.FirstName,
		&response.LastName,
		&response.LastLogin,
		&response.CreatedAt,
		&response.UpdatedAt,
	); err != nil {
		return nil, fmt.Errorf("conn.QueryRow: row.Scan: %w", err)
	}

	return &response, nil
}
