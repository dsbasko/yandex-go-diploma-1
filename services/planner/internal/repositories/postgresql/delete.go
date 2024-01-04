package postgresql

import (
	"context"
	"fmt"
)

func (r *Repository) DeleteByID(
	ctx context.Context,
	userID, id string,
) error {
	query, args, err := r.builder.
		Delete("task").
		Where("user_id = ? AND id = ?", userID, id).
		ToSql()
	if err != nil {
		return fmt.Errorf("builder: %w", err)
	}

	if _, err = r.conn.Exec(ctx, query, args...); err != nil {
		return fmt.Errorf("row.Scan: %w", err)
	}

	return nil
}
