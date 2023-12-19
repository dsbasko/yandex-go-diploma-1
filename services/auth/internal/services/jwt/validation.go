package jwt

import (
	"context"
	"fmt"

	"github.com/dsbasko/yandex-go-diploma-1/services/auth/internal/config"
	"github.com/dsbasko/yandex-go-diploma-1/services/auth/pkg/api"
	"github.com/golang-jwt/jwt/v4"
)

func (s *Service) Validation(ctx context.Context, dto *api.JWTValidationRequestV1) (*api.JWTValidationResponseV1, error) {
	if dto == nil {
		return nil, ErrArgumentsNotFilled
	}

	claims := &Claims{}

	token, err := jwt.ParseWithClaims(
		dto.Token,
		claims,
		func(t *jwt.Token) (any, error) {
			return []byte(config.GetJwtSecretKey()), nil
		})
	if err != nil {
		return nil, fmt.Errorf("jwt.ParseWithClaims: %w", err)
	}

	if _, err = s.repo.FindByUsername(ctx, claims.JWTPayload.Username); err != nil {
		token.Valid = false
		claims.JWTPayload = nil
	}

	return &api.JWTValidationResponseV1{
		IsValid: token.Valid,
		Payload: claims.JWTPayload,
	}, nil
}
