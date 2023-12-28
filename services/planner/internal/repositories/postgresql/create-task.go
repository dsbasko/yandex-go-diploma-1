package postgresql

import (
	"context"
	"fmt"
	"strings"

	"github.com/dsbasko/yandex-go-diploma-1/core/lib"
	"github.com/dsbasko/yandex-go-diploma-1/services/planner/internal/domain"
	"github.com/dsbasko/yandex-go-diploma-1/services/planner/pkg/api"
)

func (r *Repository) CreateTask(
	ctx context.Context,
	dto *api.CreateTaskRequestV1,
) (*domain.RepositoryTaskEntity, error) {
	dtoKeysAndValues := lib.StructToKeysAndValues(dto, true, true)

	query, args, err := r.builder.
		Insert("task").
		Columns(dtoKeysAndValues.Keys...).
		Values(dtoKeysAndValues.Values...).
		Suffix(fmt.Sprintf("RETURNING %s", strings.Join(
			lib.StructToKeysAndValues(
				&domain.RepositoryTaskEntity{},
				false, false,
			).Keys,
			",",
		))).
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("squirrel.ToSql: %w", err)
	}

	var response domain.RepositoryTaskEntity
	row := r.conn.QueryRow(ctx, query, args...)
	if err = row.Scan(
		&response.ID,
		&response.UserID,
		&response.Name,
		&response.Description,
		&response.CreatedAt,
		&response.UpdatedAt,
	); err != nil {
		return nil, fmt.Errorf("conn.QueryRow: row.Scan: %w", err)
	}

	return &response, nil
}
