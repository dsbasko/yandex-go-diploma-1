package account

import (
	"context"
	"fmt"

	"github.com/dsbasko/yandex-go-diploma-1/services/auth/internal/services/jwt"
	"github.com/dsbasko/yandex-go-diploma-1/services/auth/pkg/api"
)

func (s *Service) Login(ctx context.Context, dto *api.LoginRequestV1) (*api.LoginResponseV1, error) {
	if ctx == nil || dto == nil {
		return nil, ErrArgumentsNotFilled
	}

	foundUser, err := s.repo.FindByUsername(ctx, dto.Username)
	if err != nil {
		return nil, fmt.Errorf("repo.FindByUsername: %w", err)
	}

	if err = passwordCompare(foundUser.Password, dto.Password); err != nil {
		return nil, fmt.Errorf("bcrypt.CompareHashAndPassword: %w", err)
	}

	jwtService := jwt.NewService(s.log, s.repo)
	token, err := jwtService.Generate(foundUser)
	if err != nil {
		return nil, fmt.Errorf("jwtService.Generate: %w", err)
	}

	return &api.LoginResponseV1{
		UUID:  foundUser.ID,
		Token: token,
	}, nil
}
