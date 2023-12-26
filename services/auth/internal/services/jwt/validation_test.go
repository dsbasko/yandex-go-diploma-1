package jwt

import (
	"context"
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/dsbasko/yandex-go-diploma-1/core/lib"
	"github.com/dsbasko/yandex-go-diploma-1/core/logger"
	"github.com/dsbasko/yandex-go-diploma-1/services/auth/internal/config"
	"github.com/dsbasko/yandex-go-diploma-1/services/auth/internal/domain"
	"github.com/dsbasko/yandex-go-diploma-1/services/auth/internal/repositories"
	"github.com/dsbasko/yandex-go-diploma-1/services/auth/pkg/api"
	"github.com/golang-jwt/jwt/v4"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestService_Validation(t *testing.T) {
	ctx := context.Background()
	log, _ := logger.NewMock()
	repo := repositories.NewMock(t)
	service := NewService(log, repo)

	config.SetJwtExp(time.Hour)
	validToken, _ := service.Generate(&domain.RepositoryAccountEntity{
		ID:        "42",
		Username:  "test",
		Password:  "test",
		FirstName: "test",
		LastName:  "test",
		LastLogin: time.Now(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})

	tests := []struct {
		name       string
		dto        *api.JWTValidationRequestV1
		wantRes    *api.JWTValidationResponseV1
		wantErr    error
		mockConfig func()
	}{
		{
			name: "Success",
			dto: &api.JWTValidationRequestV1{
				Token: validToken,
			},
			wantRes: &api.JWTValidationResponseV1{
				IsValid: true,
				Payload: &api.JWTPayloadV1{UserID: "42", Username: "test", FirstName: "test", LastName: "test"},
			},
			wantErr: nil,
			mockConfig: func() {
				repo.EXPECT().
					FindByUsername(gomock.Any(), gomock.Any()).
					Return(&domain.RepositoryAccountEntity{ID: "42"}, nil)
			},
		},
		{
			name:       "Empty DTO",
			dto:        nil,
			wantRes:    &api.JWTValidationResponseV1{IsValid: false, Payload: nil},
			wantErr:    ErrArgumentsNotFilled,
			mockConfig: func() {},
		},
		{
			name: "Token Not Valid",
			dto: &api.JWTValidationRequestV1{
				Token: fmt.Sprintf("%sa", validToken),
			},
			wantRes:    &api.JWTValidationResponseV1{IsValid: false, Payload: nil},
			wantErr:    jwt.ErrSignatureInvalid,
			mockConfig: func() {},
		},
		{
			name: "User not found",
			dto: &api.JWTValidationRequestV1{
				Token: validToken,
			},
			wantRes: &api.JWTValidationResponseV1{IsValid: false, Payload: nil},
			wantErr: nil,
			mockConfig: func() {
				repo.EXPECT().FindByUsername(gomock.Any(), gomock.Any()).Return(nil, errors.New("not found"))
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockConfig()
			response, err := service.Validation(ctx, tt.dto)

			if err != nil || tt.wantErr != nil {
				assert.Equal(t, lib.ErrorsUnwrap(err), tt.wantErr)
			} else {
				assert.Equal(t, response.IsValid, tt.wantRes.IsValid)
				assert.Equal(t, response.Payload, tt.wantRes.Payload)
			}
		})
	}
}
