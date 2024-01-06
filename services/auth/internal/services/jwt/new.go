package jwt

import (
	"github.com/dsbasko/yandex-go-diploma-1/core/logger"
	"github.com/dsbasko/yandex-go-diploma-1/services/auth/internal/interfaces"
	"github.com/dsbasko/yandex-go-diploma-1/services/auth/pkg/api"
	"github.com/golang-jwt/jwt/v4"
)

type Service struct {
	log  *logger.Logger
	repo interfaces.Repository
}

type Claims struct {
	jwt.RegisteredClaims
	JWTPayload *api.JWTPayloadV1
}

func NewService(log *logger.Logger, repo interfaces.Repository) *Service {
	return &Service{
		log:  log,
		repo: repo,
	}
}
