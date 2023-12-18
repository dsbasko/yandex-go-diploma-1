package account

import (
	"context"
	"errors"
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

func TestService_Login(t *testing.T) {
	ctx := context.Background()
	log, _ := logger.NewMock()
	repo := repositories.NewMock(t)
	service := NewService(log, repo)

	tests := []struct {
		name       string
		dto        *api.AuthRequestV1
		wantRes    *api.AuthResponseV1
		wantErr    error
		mockConfig func()
	}{
		{
			name: "Success",
			dto: &api.AuthRequestV1{
				Username: "username",
				Password: "password",
			},
			wantRes: &api.AuthResponseV1{UUID: "42", Token: "42"},
			wantErr: nil,
			mockConfig: func() {
				repo.EXPECT().
					FindByUsername(gomock.Any(), gomock.Any()).
					Return(&domain.RepositoryAccountEntity{ID: "42", Password: "password"}, nil)
			},
		},
		{
			name:       "Empty DTO",
			dto:        nil,
			wantRes:    nil,
			wantErr:    ErrArgumentsNotFilled,
			mockConfig: func() {},
		},
		{
			name: "User Not Found",
			dto: &api.AuthRequestV1{
				Username: "username",
				Password: "password",
			},
			wantRes: nil,
			wantErr: errors.New("not found"),
			mockConfig: func() {
				repo.EXPECT().FindByUsername(gomock.Any(), gomock.Any()).Return(nil, errors.New("not found"))
			},
		},
		{
			name: "Password Not Valid",
			dto: &api.AuthRequestV1{
				Username: "username",
				Password: "password",
			},
			wantRes: nil,
			wantErr: bcrypt.ErrMismatchedHashAndPassword,
			mockConfig: func() {
				repo.EXPECT().
					FindByUsername(gomock.Any(), gomock.Any()).
					Return(&domain.RepositoryAccountEntity{ID: "42", Password: "password2"}, nil)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockConfig()

			if tt.dto != nil {
				passHash, err := passwordEncode(tt.dto.Password)
				if err != nil {
					t.Fatalf("passwordEncode: %v", err)
				}

				tt.dto.Password = passHash
			}

			response, err := service.Login(ctx, tt.dto)
			if err != nil || tt.wantErr != nil {
				assert.Equal(t, lib.ErrorsUnwrap(err), tt.wantErr)
			} else {
				assert.Equal(t, response.UUID, tt.wantRes.UUID)
				assert.NotEmpty(t, response.Token)
			}
		})
	}
}
