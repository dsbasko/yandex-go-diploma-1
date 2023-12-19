package handler

import (
	"github.com/dsbasko/yandex-go-diploma-1/core/logger"
	"github.com/dsbasko/yandex-go-diploma-1/services/auth/internal/services/account"
	"github.com/dsbasko/yandex-go-diploma-1/services/auth/internal/services/jwt"
	"github.com/go-chi/chi/v5"
)

type Handler struct {
	log            *logger.Logger
	accountService *account.Service
	jwtService     *jwt.Service
}

func Inject(
	log *logger.Logger,
	handler *chi.Mux,
	accountService *account.Service,
	jwtService *jwt.Service,
) {
	h := &Handler{
		log:            log,
		accountService: accountService,
		jwtService:     jwtService,
	}

	handler.Get("/ping", h.Ping)
	handler.Post("/register", h.Register)
	handler.Post("/login", h.Login)

	routes := handler.Routes()
	for _, route := range routes {
		for handle := range route.Handlers {
			log.Debugf("Mapped [%v] %v route", handle, route.Pattern)
		}
	}
}
