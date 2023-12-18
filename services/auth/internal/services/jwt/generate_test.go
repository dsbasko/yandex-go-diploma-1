package jwt

import (
	"testing"
	"time"

	"github.com/dsbasko/yandex-go-diploma-1/core/logger"
	"github.com/dsbasko/yandex-go-diploma-1/services/auth/internal/domain"
	"github.com/dsbasko/yandex-go-diploma-1/services/auth/internal/repositories"
	"github.com/stretchr/testify/assert"
)

func TestService_Generate(t *testing.T) {
	log, _ := logger.NewMock()
	repo := repositories.NewMock(t)
	service := NewService(log, repo)

	tests := []struct {
		name    string
		dto     *domain.RepositoryAccountEntity
		wantJWT bool
		wantErr bool
	}{
		{
			name: "Success",
			dto: &domain.RepositoryAccountEntity{
				ID:        "42",
				Username:  "test",
				Password:  "test",
				FirstName: "test",
				LastName:  "test",
				LastLogin: time.Now(),
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			wantJWT: true,
			wantErr: false,
		},
		{
			name:    "Empty DTO",
			dto:     nil,
			wantJWT: false,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			jwt, err := service.Generate(tt.dto)

			if tt.wantJWT {
				assert.NotZero(t, jwt)
			}

			if tt.wantErr {
				assert.NotZero(t, err)
			}
		})
	}
}
