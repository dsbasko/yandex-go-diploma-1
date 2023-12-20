package account

import (
	"github.com/dsbasko/yandex-go-diploma-1/core/logger"
	"github.com/dsbasko/yandex-go-diploma-1/services/auth/internal/domain"
)

const (
	UsernameMinLength int = 4
	UsernameMaxLength int = 16
	PasswordMinLength int = 8
	PasswordMaxLength int = 24
)

type Service struct {
	log  *logger.Logger
	repo domain.Repository
}

func NewService(log *logger.Logger, repo domain.Repository) *Service {
	return &Service{
		log:  log,
		repo: repo,
	}
}
