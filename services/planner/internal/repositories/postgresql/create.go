package postgresql

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/dsbasko/yandex-go-diploma-1/core/structs"
	"github.com/dsbasko/yandex-go-diploma-1/services/planner/internal/entities"
	"github.com/dsbasko/yandex-go-diploma-1/services/planner/pkg/api"
)

func (r *Repository) Create(
	ctx context.Context,
	dto *api.CreateTaskRequestV1,
) (*entities.RepositoryTaskEntity, error) {
	dtoKeys, dtoValues, err := structs.ToKeysAndValues(dto, true, &[]string{"id"})
	if err != nil {
		return nil, fmt.Errorf("structs.ToKeysAndValues: %w", err)
	}

	entityKeys, _, err := structs.ToKeysAndValues(entities.RepositoryTaskEntity{}, true, &[]string{"id"})
	if err != nil {
		return nil, fmt.Errorf("structs.ToKeysAndValues: %w", err)
	}

	query, args, err := r.builder.
		Insert("task").
		Columns(dtoKeys...).
		Values(dtoValues...).
		Suffix(fmt.Sprintf("RETURNING %s", strings.Join(entityKeys, ","))).
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("squirrel.ToSql: %w", err)
	}

	var response entities.RepositoryTaskEntity
	var dueDate sql.NullTime
	row := r.conn.QueryRow(ctx, query, args...)
	if err = row.Scan(
		&response.ID,
		&response.UserID,
		&response.Name,
		&response.Description,
		&dueDate,
		&response.IsArchive,
		&response.CreatedAt,
		&response.UpdatedAt,
	); err != nil {
		return nil, fmt.Errorf("conn.QueryRow: row.Scan: %w", err)
	}
	response.DueDate = dueDate.Time

	return &response, nil
}
