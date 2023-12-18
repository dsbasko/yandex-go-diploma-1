package domain

import (
	"context"
	"time"

	"github.com/dsbasko/yandex-go-diploma-1/services/auth/pkg/api"
)

//go:generate mockgen -destination=../repositories/mock/mock.go -package=mock github.com/dsbasko/yandex-go-diploma-1/services/auth/internal/domain Repository

type Repository interface {
	Ping(ctx context.Context) error
	CreateOnce(ctx context.Context, dto *api.RegisterRequestV1) (*RepositoryAccountEntity, error)
}

type RepositoryAccountEntity struct {
	ID        string        `json:"id"`
	Username  string        `json:"username"`
	Password  string        `json:"password"`
	LastLogin time.Duration `json:"last_login"`
	CreatedAt time.Duration `json:"created_at"`
	UpdatedAt time.Duration `json:"updated_at"`
}
