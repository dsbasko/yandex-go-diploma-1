package postgresql

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/dsbasko/yandex-go-diploma-1/core/structs"
	"github.com/dsbasko/yandex-go-diploma-1/services/planner/internal/entities"
	"github.com/jackc/pgx/v5"
)

func (r *Repository) FindByID(
	ctx context.Context,
	userID, id string,
) (*entities.RepositoryTaskEntity, error) {
	entityKeys, _, err := structs.ToKeysAndValues(entities.RepositoryTaskEntity{}, false, nil)
	if err != nil {
		return nil, fmt.Errorf("structs.ToKeysAndValues: %w", err)
	}

	query, args, err := r.builder.
		Select(entityKeys...).
		From("task").
		Where("user_id = ? AND id = ?", userID, id).
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("builder: %w", err)
	}

	var response entities.RepositoryTaskEntity
	row := r.conn.QueryRow(ctx, query, args...)

	var dueDate sql.NullTime
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
		if err == pgx.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("row.Scan: %w", err)
	}
	response.DueDate = dueDate.Time

	return &response, nil
}

func (r *Repository) FindByUserIDAndDate(
	ctx context.Context,
	userID string, dateStart, dateEnd *time.Time,
) (*[]entities.RepositoryTaskEntity, error) {
	var whereQuery any
	var whereArgs []any

	switch {
	case dateStart == nil && dateEnd == nil:
		whereQuery = "user_id = ? AND due_date IS NULL"
		whereArgs = []any{userID}
	case dateStart == nil:
		whereQuery = "user_id = ? AND due_date <= ?"
		whereArgs = []any{userID, dateEnd}
	case dateEnd == nil:
		whereQuery = "user_id = ? AND due_date >= ?"
		whereArgs = []any{userID, dateStart}
	default:
		whereQuery = "user_id = ? AND due_date >= ? AND due_date <= ?"
		whereArgs = []any{userID, dateStart, dateEnd}
	}

	entityKeys, _, err := structs.ToKeysAndValues(entities.RepositoryTaskEntity{}, false, nil)
	if err != nil {
		return nil, fmt.Errorf("structs.ToKeysAndValues: %w", err)
	}

	query, args, err := r.builder.
		Select(entityKeys...).
		From("task").
		Where(whereQuery, whereArgs...).
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
		var dueDate sql.NullTime
		if err = rows.Scan(
			&res.ID,
			&res.UserID,
			&res.Name,
			&res.Description,
			&dueDate,
			&res.IsArchive,
			&res.CreatedAt,
			&res.UpdatedAt,
		); err != nil {
			return nil, fmt.Errorf("rows.Scan: %w", err)
		}
		res.DueDate = dueDate.Time
		response = append(response, res)
	}

	return &response, nil
}

func (r *Repository) FindArchive(
	ctx context.Context,
	userID string,
) (*[]entities.RepositoryTaskEntity, error) {
	entityKeys, _, err := structs.ToKeysAndValues(entities.RepositoryTaskEntity{}, false, nil)
	if err != nil {
		return nil, fmt.Errorf("structs.ToKeysAndValues: %w", err)
	}

	query, args, err := r.builder.
		Select(entityKeys...).
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
		var dueDate sql.NullTime
		if err = rows.Scan(
			&res.ID,
			&res.UserID,
			&res.Name,
			&res.Description,
			&dueDate,
			&res.IsArchive,
			&res.CreatedAt,
			&res.UpdatedAt,
		); err != nil {
			return nil, fmt.Errorf("rows.Scan: %w", err)
		}
		res.DueDate = dueDate.Time
		response = append(response, res)
	}

	return &response, nil
}
