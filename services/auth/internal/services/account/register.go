package account

import (
	"context"
	"fmt"

	"github.com/dsbasko/yandex-go-diploma-1/services/auth/pkg/api"
)

const (
	UsernameMinLength int = 4
	UsernameMaxLength int = 16
	PasswordMinLength int = 8
	PasswordMaxLength int = 24
)

func (s *Service) Register(ctx context.Context, dto *api.RegisterRequestV1) (*api.RegisterResponseV1, error) {
	if ctx == nil || dto == nil {
		return nil, ErrArgumentsNotFilled
	}

	switch {
	case len(dto.Username) < UsernameMinLength:
		return nil, ErrUsernameMinLength
	case len(dto.Username) > UsernameMaxLength:
		return nil, ErrUsernameMaxLength
	case len(dto.Password) < PasswordMinLength:
		return nil, ErrPasswordMinLength
	case len(dto.Password) > PasswordMaxLength:
		return nil, ErrPasswordMaxLength
	default:
	}

	passHash, err := passwordEncode(dto.Password)
	if err != nil {
		return nil, fmt.Errorf("passwordEncode: %w", err)
	}
	dto.Password = passHash

	createdUser, err := s.repo.CreateOnce(ctx, dto)
	if err != nil {
		return nil, fmt.Errorf("repo.Create: %w", err)
	}
	_ = createdUser

	return &api.RegisterResponseV1{
		UUID: createdUser.ID,
	}, nil
}
