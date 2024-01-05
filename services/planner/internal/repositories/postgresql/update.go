package postgresql

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/dsbasko/yandex-go-diploma-1/core/lib"
	"github.com/dsbasko/yandex-go-diploma-1/services/planner/internal/entities"
	"github.com/dsbasko/yandex-go-diploma-1/services/planner/pkg/api"
)

func (r *Repository) UpdateOnce(
	ctx context.Context,
	userID, id string,
	dto *api.UpdateTaskRequestV1,
) (*entities.RepositoryTaskEntity, error) {
	dtoKeysAndValues := lib.StructToKeysAndValues(dto, true, true)
	setMap := map[string]any{}
	for i, key := range dtoKeysAndValues.Keys {
		setMap[key] = dtoKeysAndValues.Values[i]
	}

	query, args, err := r.builder.
		Update("task").
		SetMap(setMap).
		Where("user_id = ? AND id = ?", userID, id).
		Suffix(fmt.Sprintf("RETURNING %s", strings.Join(
			lib.StructToKeysAndValues(
				&entities.RepositoryTaskEntity{},
				false, false,
			).Keys,
			",",
		))).
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

func (r *Repository) UpdateIsArchive(
	ctx context.Context,
	userID, id string,
	isArchive bool,
) (*entities.RepositoryTaskEntity, error) {
	query, args, err := r.builder.
		Update("task").
		Set("is_archive", isArchive).
		Where("user_id = ? AND id = ?", userID, id).
		Suffix(fmt.Sprintf("RETURNING %s", strings.Join(
			lib.StructToKeysAndValues(
				&entities.RepositoryTaskEntity{},
				false, false,
			).Keys,
			",",
		))).
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("squirrel.ToSql: %w", err)
	}

	fmt.Println("query", query)
	fmt.Println("args", args)

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
