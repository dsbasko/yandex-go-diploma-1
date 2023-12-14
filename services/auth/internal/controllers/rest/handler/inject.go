package handler

import (
	"github.com/dsbasko/yandex-go-diploma-1/core/logger"
	"github.com/go-chi/chi/v5"
)

type Handler struct {
	log *logger.Logger
}

func Inject(
	log *logger.Logger,
	handler *chi.Mux,
) {
	h := &Handler{
		log: log,
	}

	handler.Get("/ping", h.Ping)

	routes := handler.Routes()
	for _, route := range routes {
		for handle := range route.Handlers {
			log.Debugf("Mapped [%v] %v route", handle, route.Pattern)
		}
	}
}
