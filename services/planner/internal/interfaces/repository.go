package interfaces

import (
	"context"
	"time"

	"github.com/dsbasko/yandex-go-diploma-1/services/planner/internal/entities"
	"github.com/dsbasko/yandex-go-diploma-1/services/planner/pkg/api"
)

//go:generate ../../../../bin/mockgen -destination=../repositories/mock/mock.go -package=mock github.com/dsbasko/yandex-go-diploma-1/services/planner/internal/interfaces Repository

type Repository interface {
	Ping(ctx context.Context) error
	Create(ctx context.Context, dto *api.CreateTaskRequestV1) (*entities.RepositoryTaskEntity, error)
	FindByID(ctx context.Context, userID, id string) (*entities.RepositoryTaskEntity, error)
	FindByUserIDAndDate(ctx context.Context, userID string, dateStart, dateEnd *time.Time) (*[]entities.RepositoryTaskEntity, error)
	FindArchive(ctx context.Context, userID string) (*[]entities.RepositoryTaskEntity, error)
	UpdateOnce(ctx context.Context, userID, id string, dto *api.UpdateTaskRequestV1) (*entities.RepositoryTaskEntity, error)
	UpdateIsArchive(ctx context.Context, userID, id string, isArchive bool) (*entities.RepositoryTaskEntity, error)
	UpdateDueDate(ctx context.Context, userID, id string, dueDate time.Time) (*entities.RepositoryTaskEntity, error)
	DeleteByID(ctx context.Context, userID, id string) error
}