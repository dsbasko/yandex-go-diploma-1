package handles

import (
	"github.com/dsbasko/yandex-go-diploma-1/core/logger"
	"github.com/dsbasko/yandex-go-diploma-1/services/auth/internal/interfaces"
	"github.com/dsbasko/yandex-go-diploma-1/services/auth/internal/services/account"
	"github.com/dsbasko/yandex-go-diploma-1/services/auth/internal/services/jwt"
)

type Handler struct {
	log            *logger.Logger
	repo           interfaces.Repository
	accountService *account.Service
	jwtService     *jwt.Service
}

func New(log *logger.Logger, repo interfaces.Repository, accountService *account.Service, jwtService *jwt.Service) *Handler {
	h := &Handler{
		log:            log,
		repo:           repo,
		accountService: accountService,
		jwtService:     jwtService,
	}

	return h
}
