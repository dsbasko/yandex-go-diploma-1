package domain

import (
	"context"
	"time"

	"github.com/dsbasko/yandex-go-diploma-1/services/planner/pkg/api"
)

//go:generate ../../../../bin/mockgen -destination=../repositories/mock/mock.go -package=mock github.com/dsbasko/yandex-go-diploma-1/services/planner/internal/domain Repository

type Repository interface {
	Ping(ctx context.Context) error
	CreateTask(ctx context.Context, dto *api.CreateTaskRequestV1) (*RepositoryTaskEntity, error)
}

type RepositoryTaskEntity struct {
	ID          string    `json:"id"`
	UserID      string    `json:"user_id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
