package postgresql

import (
	"context"
	"fmt"
	"time"

	"github.com/dsbasko/yandex-go-diploma-1/core/lib"
	"github.com/dsbasko/yandex-go-diploma-1/services/planner/internal/entities"
	"github.com/jackc/pgx/v5"
)

func (r *Repository) FindByID(
	ctx context.Context,
	userID, id string,
) (*entities.RepositoryTaskEntity, error) {
	query, args, err := r.builder.
		Select(lib.StructToKeysAndValues(
			&entities.RepositoryTaskEntity{}, false, false,
		).Keys...).
		From("task").
		Where("user_id = ? and id = ?", userID, id).
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("builder: %w", err)
	}

	var response entities.RepositoryTaskEntity
	row := r.conn.QueryRow(ctx, query, args...)

	if err = row.Scan(
		&response.ID,
		&response.UserID,
		&response.Name,
		&response.Description,
		&response.DueDate,
		&response.IsArchive,
		&response.CreatedAt,
		&response.UpdatedAt,
	); err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("row.Scan: %w", err)
	}

	return &response, nil
}

func (r *Repository) FindByUserIDAndDate(
	ctx context.Context,
	userID string, dateStart, dateEnd time.Time,
) (*[]entities.RepositoryTaskEntity, error) {
	query, args, err := r.builder.
		Select(lib.StructToKeysAndValues(
			&entities.RepositoryTaskEntity{}, false, false,
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

	var response []entities.RepositoryTaskEntity
	rows, err := r.conn.Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("conn.Query: %w", err)
	}

	for rows.Next() {
		var res entities.RepositoryTaskEntity
		if err = rows.Scan(
			&res.ID,
			&res.UserID,
			&res.Name,
			&res.Description,
			&res.DueDate,
			&res.IsArchive,
			&res.CreatedAt,
			&res.UpdatedAt,
		); err != nil {
			return nil, fmt.Errorf("rows.Scan: %w", err)
		}
		response = append(response, res)
	}

	return &response, nil
}

func (r *Repository) FindArchive(
	ctx context.Context,
	userID string,
) (*[]entities.RepositoryTaskEntity, error) {
	query, args, err := r.builder.
		Select(lib.StructToKeysAndValues(
			&entities.RepositoryTaskEntity{}, false, false,
		).Keys...).
		From("task").
		Where("user_id = ? AND is_archive = TRUE", userID).
		OrderBy("due_date DESC").
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("builder: %w", err)
	}

	var response []entities.RepositoryTaskEntity
	rows, err := r.conn.Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("conn.Query: %w", err)
	}

	for rows.Next() {
		var res entities.RepositoryTaskEntity
		if err = rows.Scan(
			&res.ID,
			&res.UserID,
			&res.Name,
			&res.Description,
			&res.DueDate,
			&res.IsArchive,
			&res.CreatedAt,
			&res.UpdatedAt,
		); err != nil {
			return nil, fmt.Errorf("rows.Scan: %w", err)
		}
		response = append(response, res)
	}

	return &response, nil
}
