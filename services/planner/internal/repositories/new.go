package repositories

import (
	"context"
	"fmt"
	"testing"

	"github.com/dsbasko/yandex-go-diploma-1/services/planner/internal/interfaces"
	"github.com/dsbasko/yandex-go-diploma-1/services/planner/internal/repositories/mock"
	"github.com/dsbasko/yandex-go-diploma-1/services/planner/internal/repositories/postgresql"
	"github.com/golang/mock/gomock"
)

func NewRepository(ctx context.Context) (interfaces.Repository, error) {
	repo, err := postgresql.NewRepository(ctx)
	if err != nil {
		return nil, fmt.Errorf("postgresql.NewRepository: %w", err)
	}
	return repo, nil
}

func NewMock(t *testing.T) *mock.MockRepository {
	controller := gomock.NewController(t)
	defer controller.Finish()
	return mock.NewMockRepository(controller)
}