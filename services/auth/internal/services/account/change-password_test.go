package account

import (
	"context"
	"errors"
	"strings"
	"testing"

	"github.com/dsbasko/yandex-go-diploma-1/core/lib"
	"github.com/dsbasko/yandex-go-diploma-1/core/logger"
	"github.com/dsbasko/yandex-go-diploma-1/services/auth/internal/domain"
	"github.com/dsbasko/yandex-go-diploma-1/services/auth/internal/repositories"
	"github.com/dsbasko/yandex-go-diploma-1/services/auth/pkg/api"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

func TestService_ChangePassword(t *testing.T) {
	ctx := context.Background()
	log, _ := logger.NewMock()
	repo := repositories.NewMock(t)
	service := NewService(log, repo)

	passHash, _ := passwordEncode("old_password")

	tests := []struct {
		name       string
		dto        *api.ChangePasswordRequestV1
		wantRes    *api.ChangePasswordResponseV1
		wantErr    error
		mockConfig func()
	}{
		{
			name: "Success",
			dto: &api.ChangePasswordRequestV1{
				OldPassword: "old_password",
				NewPassword: "new_password",
			},
			wantRes: &api.ChangePasswordResponseV1{UUID: "42"},
			wantErr: nil,
			mockConfig: func() {
				repo.EXPECT().FindByID(gomock.Any(), gomock.Any()).Return(&domain.RepositoryAccountEntity{ID: "42", Password: passHash}, nil)
				repo.EXPECT().UpdateOnce(gomock.Any(), gomock.Any()).Return(&domain.RepositoryAccountEntity{ID: "42"}, nil)
			},
		},
		{
			name: "Short password",
			dto: &api.ChangePasswordRequestV1{
				OldPassword: "old_password",
				NewPassword: "pass",
			},
			wantRes:    &api.ChangePasswordResponseV1{UUID: "42"},
			wantErr:    ErrPasswordMinLength,
			mockConfig: func() {},
		},
		{
			name: "Long password",
			dto: &api.ChangePasswordRequestV1{
				OldPassword: "old_password",
				NewPassword: strings.Repeat("new_password", 42),
			},
			wantRes:    &api.ChangePasswordResponseV1{UUID: "42"},
			wantErr:    ErrPasswordMaxLength,
			mockConfig: func() {},
		},
		{
			name: "User Not Found",
			dto: &api.ChangePasswordRequestV1{
				OldPassword: "old_password",
				NewPassword: "new_password",
			},
			wantRes: &api.ChangePasswordResponseV1{UUID: "42"},
			wantErr: errors.New("not found"),
			mockConfig: func() {
				repo.EXPECT().FindByID(gomock.Any(), gomock.Any()).Return(nil, errors.New("not found"))
			},
		},
		{
			name: "Wrong Old Password",
			dto: &api.ChangePasswordRequestV1{
				OldPassword: "old_password2",
				NewPassword: "new_password",
			},
			wantRes: &api.ChangePasswordResponseV1{UUID: "42"},
			wantErr: bcrypt.ErrMismatchedHashAndPassword,
			mockConfig: func() {
				repo.EXPECT().FindByID(gomock.Any(), gomock.Any()).Return(
					&domain.RepositoryAccountEntity{ID: "42", Password: passHash},
					nil,
				)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockConfig()
			response, err := service.ChangePassword(ctx, "42", tt.dto)
			if err != nil || tt.wantErr != nil {
				assert.Equal(t, lib.ErrorsUnwrap(err), tt.wantErr)
			} else {
				assert.Equal(t, response, tt.wantRes)
			}
		})
	}
}
