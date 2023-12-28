package handles

import (
	"testing"

	"github.com/dsbasko/yandex-go-diploma-1/core/logger"
	"github.com/dsbasko/yandex-go-diploma-1/services/planner/internal/repositories"
	"github.com/dsbasko/yandex-go-diploma-1/services/planner/internal/services/task"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	log := logger.NewMock()
	repo := repositories.NewMock(t)
	taskService := task.NewService(log, repo)
	handler := New(log, repo, taskService)

	t.Run("Implement Type", func(t *testing.T) {
		mockHandler := &Handler{log: log, repo: repo, taskService: taskService}
		assert.Equal(t, handler, mockHandler)
	})

	t.Run("Not Nil", func(t *testing.T) {
		assert.NotNil(t, handler)
	})
}
