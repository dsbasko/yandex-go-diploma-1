package postgresql

import (
	"context"
	"fmt"

	"github.com/dsbasko/yandex-go-diploma-1/core/structs"
	"github.com/dsbasko/yandex-go-diploma-1/services/auth/internal/entities"
)

func (r *Repository) FindByID(
	ctx context.Context,
	id string,
) (*entities.RepositoryAccountEntity, error) {
	entityKeys, _, err := structs.ToKeysAndValues(entities.RepositoryAccountEntity{}, false, nil)
	if err != nil {
		return nil, fmt.Errorf("structs.ToKeysAndValues: %w", err)
	}

	query, args, err := r.builder.
		Select(entityKeys...).
		From("accounts").
		Where("id = ?", id).
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("squirrel.ToSql: %w", err)
	}

	var response entities.RepositoryAccountEntity
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
) (*entities.RepositoryAccountEntity, error) {
	entityKeys, _, err := structs.ToKeysAndValues(entities.RepositoryAccountEntity{}, false, nil)
	if err != nil {
		return nil, fmt.Errorf("structs.ToKeysAndValues: %w", err)
	}

	query, args, err := r.builder.
		Select(entityKeys...).
		From("accounts").
		Where("username = ?", username).
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("squirrel.ToSql: %w", err)
	}

	var response entities.RepositoryAccountEntity
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
