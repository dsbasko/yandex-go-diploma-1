package account

import (
	"context"
	"strings"
	"testing"

	"github.com/dsbasko/yandex-go-diploma-1/core/errors"
	"github.com/dsbasko/yandex-go-diploma-1/core/logger"
	"github.com/dsbasko/yandex-go-diploma-1/services/auth/internal/entities"
	"github.com/dsbasko/yandex-go-diploma-1/services/auth/internal/repositories"
	"github.com/dsbasko/yandex-go-diploma-1/services/auth/pkg/api"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestService_Register(t *testing.T) {
	ctx := context.Background()
	log := logger.NewMock()
	repo := repositories.NewMock(t)
	service := NewService(log, repo)

	tests := []struct {
		name       string
		dto        *api.RegisterRequestV1
		wantRes    *api.RegisterResponseV1
		wantErr    error
		mockConfig func()
	}{
		{
			name: "Success",
			dto: &api.RegisterRequestV1{
				Username: "test_username",
				Password: "test_password",
			},
			wantRes: &api.RegisterResponseV1{UUID: "42"},
			wantErr: nil,
			mockConfig: func() {
				repo.EXPECT().CreateOnce(gomock.Any(), gomock.Any()).Return(&entities.RepositoryAccountEntity{ID: "42"}, nil)
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
			name: "Short username",
			dto: &api.RegisterRequestV1{
				Username: "t",
				Password: "test_password",
			},
			wantRes:    nil,
			wantErr:    ErrUsernameMinLength,
			mockConfig: func() {},
		},
		{
			name: "Long username",
			dto: &api.RegisterRequestV1{
				Username: strings.Repeat("test_username", 42),
				Password: "test_password",
			},
			wantRes:    nil,
			wantErr:    ErrUsernameMaxLength,
			mockConfig: func() {},
		},
		{
			name: "Short password",
			dto: &api.RegisterRequestV1{
				Username: "test_username",
				Password: "t",
			},
			wantRes:    nil,
			wantErr:    ErrPasswordMinLength,
			mockConfig: func() {},
		},
		{
			name: "Long password",
			dto: &api.RegisterRequestV1{
				Username: "test_password",
				Password: strings.Repeat("test_password", 42),
			},
			wantRes:    nil,
			wantErr:    ErrPasswordMaxLength,
			mockConfig: func() {},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockConfig()
			response, err := service.Register(ctx, tt.dto)
			if err != nil || tt.wantErr != nil {
				assert.Equal(t, errors.Unwrap(err), tt.wantErr)
			} else {
				assert.Equal(t, response, tt.wantRes)
			}
		})
	}
}
