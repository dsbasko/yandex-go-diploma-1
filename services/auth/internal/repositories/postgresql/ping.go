package postgresql

import (
	"context"
	"fmt"
)

func (r *Repository) Ping(ctx context.Context) error {
	if err := r.conn.Ping(ctx); err != nil {
		return fmt.Errorf("conn.Ping: %w", err)
	}

	return nil
}
