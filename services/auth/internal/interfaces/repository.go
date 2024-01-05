package interfaces

import (
	"context"

	"github.com/dsbasko/yandex-go-diploma-1/services/auth/internal/entities"
	"github.com/dsbasko/yandex-go-diploma-1/services/auth/pkg/api"
)

//go:generate ../../../../bin/mockgen -destination=../repositories/mock/mock.go -package=mock github.com/dsbasko/yandex-go-diploma-1/services/auth/internal/interfaces Repository

type Repository interface {
	Ping(ctx context.Context) error
	CreateOnce(ctx context.Context, dto *api.RegisterRequestV1) (*entities.RepositoryAccountEntity, error)
	FindByID(ctx context.Context, id string) (*entities.RepositoryAccountEntity, error)
	FindByUsername(ctx context.Context, username string) (*entities.RepositoryAccountEntity, error)
	UpdateOnce(ctx context.Context, dto *entities.RepositoryAccountEntity) (*entities.RepositoryAccountEntity, error)
}
