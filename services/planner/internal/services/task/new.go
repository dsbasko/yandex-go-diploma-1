package task

import (
	"github.com/dsbasko/yandex-go-diploma-1/core/logger"
	"github.com/dsbasko/yandex-go-diploma-1/services/planner/internal/interfaces"
)

type Service struct {
	log  *logger.Logger
	repo interfaces.Repository
}

func NewService(log *logger.Logger, repo interfaces.Repository) *Service {
	return &Service{
		log:  log,
		repo: repo,
	}
}
