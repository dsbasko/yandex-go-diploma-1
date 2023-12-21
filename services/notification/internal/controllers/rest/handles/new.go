package handles

import (
	"github.com/dsbasko/yandex-go-diploma-1/core/logger"
)

type Handler struct {
	log *logger.Logger
}

func New(log *logger.Logger) *Handler {
	h := &Handler{
		log: log,
	}

	return h
}
