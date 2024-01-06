package handles

import (
	"testing"

	"github.com/dsbasko/yandex-go-diploma-1/core/logger"
	"github.com/dsbasko/yandex-go-diploma-1/services/notification/internal/repositories"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	log := logger.NewMock()
	repo := repositories.NewMock(t)
	handler := New(log, repo)

	t.Run("Implement Type", func(t *testing.T) {
		mockHandler := &Handler{log: log, repo: repo}
		assert.Equal(t, handler, mockHandler)
	})

	t.Run("Not Nil", func(t *testing.T) {
		assert.NotNil(t, handler)
	})
}
