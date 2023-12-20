package postgresql

import (
	"context"
	"fmt"
	"strings"

	"github.com/dsbasko/yandex-go-diploma-1/core/lib"
	"github.com/dsbasko/yandex-go-diploma-1/services/auth/internal/domain"
)

func (r *Repository) UpdateOnce(
	ctx context.Context,
	dto *domain.RepositoryAccountEntity,
) (*domain.RepositoryAccountEntity, error) {
	dtoKeysAndValues := lib.StructToKeysAndValues(dto, true, true)
	setMap := map[string]any{}
	for i, key := range dtoKeysAndValues.Keys {
		setMap[key] = dtoKeysAndValues.Values[i]
	}

	query, args, err := r.builder.
		Update("accounts").
		SetMap(setMap).
		Where("id = ?", dto.ID).
		Suffix(fmt.Sprintf("RETURNING %s", strings.Join(
			lib.StructToKeysAndValues(
				&domain.RepositoryAccountEntity{},
				false, false,
			).Keys,
			",",
		))).
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