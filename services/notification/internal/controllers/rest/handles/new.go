package handles

import (
	"github.com/dsbasko/yandex-go-diploma-1/core/logger"
	"github.com/dsbasko/yandex-go-diploma-1/services/notification/internal/interfaces"
)

type Handler struct {
	log  *logger.Logger
	repo interfaces.Repository
}

func New(log *logger.Logger, repo interfaces.Repository) *Handler {
	h := &Handler{
		log:  log,
		repo: repo,
	}

	return h
}
