package postgresql

import (
	"context"
	"fmt"
	"strings"

	"github.com/dsbasko/yandex-go-diploma-1/core/structs"
	"github.com/dsbasko/yandex-go-diploma-1/services/auth/internal/domain"
)

func (r *Repository) UpdateOnce(
	ctx context.Context,
	dto *domain.RepositoryAccountEntity,
) (*domain.RepositoryAccountEntity, error) {
	dtoKeys, dtoValues, err := structs.ToKeysAndValues(dto, true, &[]string{"id"})
	if err != nil {
		return nil, fmt.Errorf("structs.ToKeysAndValues: %w", err)
	}

	setMap := map[string]any{}
	for i, key := range dtoKeys {
		setMap[key] = dtoValues[i]
	}

	entityKeys, _, err := structs.ToKeysAndValues(domain.RepositoryAccountEntity{}, true, nil)
	if err != nil {
		return nil, fmt.Errorf("structs.ToKeysAndValues: %w", err)
	}

	query, args, err := r.builder.
		Update("accounts").
		SetMap(setMap).
		Where("id = ?", dto.ID).
		Suffix(fmt.Sprintf("RETURNING %s", strings.Join(entityKeys, ","))).
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
