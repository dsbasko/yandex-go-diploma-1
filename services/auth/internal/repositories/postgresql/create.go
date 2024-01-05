package postgresql

import (
	"context"
	"fmt"
	"strings"

	"github.com/dsbasko/yandex-go-diploma-1/core/structs"
	"github.com/dsbasko/yandex-go-diploma-1/services/auth/internal/domain"
	"github.com/dsbasko/yandex-go-diploma-1/services/auth/pkg/api"
)

func (r *Repository) CreateOnce(
	ctx context.Context,
	dto *api.RegisterRequestV1,
) (*domain.RepositoryAccountEntity, error) {
	dtoKeysAndValues := structs.ToKeysAndValues(dto, true, &[]string{"id"})

	query, args, err := r.builder.
		Insert("accounts").
		Columns(dtoKeysAndValues.Keys...).
		Values(dtoKeysAndValues.Values...).
		Suffix(fmt.Sprintf("RETURNING %s", strings.Join(
			structs.ToKeysAndValues(
				&domain.RepositoryAccountEntity{},
				false, nil,
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
