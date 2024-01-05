package handles

import (
	"github.com/dsbasko/yandex-go-diploma-1/core/logger"
	"github.com/dsbasko/yandex-go-diploma-1/services/planner/internal/interfaces"
	"github.com/dsbasko/yandex-go-diploma-1/services/planner/internal/services/task"
)

type Handler struct {
	log         *logger.Logger
	taskService *task.Service
	repo        interfaces.Repository
}

func New(log *logger.Logger, repo interfaces.Repository, taskService *task.Service) *Handler {
	h := &Handler{
		log:         log,
		repo:        repo,
		taskService: taskService,
	}

	return h
}
