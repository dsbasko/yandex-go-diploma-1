package account

import (
	"context"
	"fmt"

	"github.com/dsbasko/yandex-go-diploma-1/services/auth/internal/controllers/rest/middleware"
	"github.com/dsbasko/yandex-go-diploma-1/services/auth/internal/domain"
	"github.com/dsbasko/yandex-go-diploma-1/services/auth/pkg/api"
)

func (s *Service) ChangePassword(ctx context.Context, dto *api.ChangePasswordRequestV1) (*api.ChangePasswordResponseV1, error) {
	if ctx == nil || dto == nil {
		return nil, ErrArgumentsNotFilled
	}

	switch {
	case len(dto.NewPassword) < PasswordMinLength:
		return nil, ErrPasswordMinLength
	case len(dto.NewPassword) > PasswordMaxLength:
		return nil, ErrPasswordMaxLength
	default:
	}

	authPayload := middleware.GetAuthPayload(ctx)
	if authPayload == nil {
		return nil, ErrUnauthorized
	}

	foundUser, err := s.repo.FindByID(ctx, authPayload.UserID)
	if err != nil {
		return nil, fmt.Errorf("repo.FindByID: %w", err)
	}

	if err = passwordCompare(foundUser.Password, dto.OldPassword); err != nil {
		return nil, fmt.Errorf("passwordCompare: %w", err)
	}

	passHash, err := passwordEncode(dto.NewPassword)
	if err != nil {
		return nil, fmt.Errorf("passwordEncode: %w", err)
	}

	response, err := s.repo.UpdateOnce(ctx, &domain.RepositoryAccountEntity{
		ID:       authPayload.UserID,
		Password: passHash,
	})
	if err != nil {
		return nil, fmt.Errorf("repo.UpdateOnce: %w", err)
	}

	return &api.ChangePasswordResponseV1{
		UUID: response.ID,
	}, nil
}
