package repositories

import (
	"context"
	"fmt"

	"github.com/dsbasko/yandex-go-diploma-1/services/auth/internal/domain"
	"github.com/dsbasko/yandex-go-diploma-1/services/auth/internal/repositories/postgresql"
)

func NewRepository(ctx context.Context) (domain.Repository, error) {
	repo, err := postgresql.NewRepository(ctx)
	return repo, fmt.Errorf("postgresql.NewRepository: %w", err)
}
