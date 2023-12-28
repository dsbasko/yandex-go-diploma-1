package task

import (
	"github.com/dsbasko/yandex-go-diploma-1/core/logger"
	"github.com/dsbasko/yandex-go-diploma-1/services/planner/internal/domain"
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
