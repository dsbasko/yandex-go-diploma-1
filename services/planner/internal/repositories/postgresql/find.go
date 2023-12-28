package postgresql

import (
	"context"
	"fmt"
	"time"

	"github.com/dsbasko/yandex-go-diploma-1/core/lib"
	"github.com/dsbasko/yandex-go-diploma-1/services/planner/internal/domain"
)

func (r *Repository) FindByUserIDAndDate(
	ctx context.Context,
	userID string, dateStart, dateEnd time.Time,
) (*[]domain.RepositoryTaskEntity, error) {
	query, args, err := r.builder.
		Select(lib.StructToKeysAndValues(
			&domain.RepositoryTaskEntity{}, false, false,
		).Keys...).
		From("task").
		Where(
			"user_id = ? AND due_date >= ? AND due_date <= ?",
			userID, dateStart, dateEnd,
		).
		OrderBy("due_date ASC").
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("builder: %w", err)
	}

	var response []domain.RepositoryTaskEntity
	rows, err := r.conn.Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("conn.Query: %w", err)
	}

	for rows.Next() {
		var res domain.RepositoryTaskEntity
		if err = rows.Scan(
			&res.ID,
			&res.UserID,
			&res.Name,
			&res.Description,
			&res.DueDate,
			&res.CreatedAt,
			&res.UpdatedAt,
		); err != nil {
			return nil, fmt.Errorf("rows.Scan: %w", err)
		}
		response = append(response, res)
	}

	return &response, nil
}
