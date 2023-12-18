package jwt

import (
	"fmt"
	"time"

	"github.com/dsbasko/yandex-go-diploma-1/services/auth/internal/config"
	"github.com/dsbasko/yandex-go-diploma-1/services/auth/internal/domain"
	"github.com/golang-jwt/jwt/v4"
)

type Claims struct {
	jwt.RegisteredClaims
	UserID    string `json:"user_id"`
	Username  string `json:"username"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func (s *Service) Generate(
	accountEntity *domain.RepositoryAccountEntity,
) (string, error) {
	if accountEntity == nil {
		return "", ErrArgumentsNotFilled
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(config.GetJwtExp())),
		},
		UserID:    accountEntity.ID,
		Username:  accountEntity.Username,
		FirstName: accountEntity.FirstName,
		LastName:  accountEntity.LastName,
	})

	tokenString, err := token.SignedString([]byte(config.GetJwtSecretKey()))
	if err != nil {
		return "", fmt.Errorf("token.SignedString: %w", err)
	}

	return tokenString, nil
}
