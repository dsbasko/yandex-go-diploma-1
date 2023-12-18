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
	FindByUsername(ctx context.Context, username string) (*RepositoryAccountEntity, error)
}

type RepositoryAccountEntity struct {
	ID        string    `json:"id"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	LastLogin time.Time `json:"last_login"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
