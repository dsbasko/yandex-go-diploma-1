package jwt

import (
	"testing"

	"github.com/dsbasko/yandex-go-diploma-1/core/logger"
	"github.com/dsbasko/yandex-go-diploma-1/services/auth/internal/repositories"
	"github.com/stretchr/testify/assert"
)

func TestNewService(t *testing.T) {
	log := logger.NewMock()
	repo := repositories.NewMock(t)
	service := NewService(log, repo)

	t.Run("Implement Type", func(t *testing.T) {
		mockService := &Service{log: log, repo: repo}
		assert.Equal(t, service, mockService)
	})

	t.Run("Not Nil", func(t *testing.T) {
		assert.NotNil(t, service)
	})
}
