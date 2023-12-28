package task

import (
	"context"
	"strings"
	"testing"

	"github.com/dsbasko/yandex-go-diploma-1/core/logger"
	"github.com/dsbasko/yandex-go-diploma-1/services/planner/internal/domain"
	"github.com/dsbasko/yandex-go-diploma-1/services/planner/internal/repositories"
	"github.com/dsbasko/yandex-go-diploma-1/services/planner/pkg/api"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestService_Create(t *testing.T) {
	ctx := context.Background()
	log := logger.NewMock()
	repo := repositories.NewMock(t)
	service := NewService(log, repo)

	tests := []struct {
		name     string
		dto      *api.CreateTaskRequestV1
		repoConf func()
		wantRes  *api.CreateTaskResponseV1
		wantErr  error
	}{
		{
			name: "Success",
			dto: &api.CreateTaskRequestV1{
				UserID:      "42",
				Name:        "test task",
				Description: "test description",
			},
			wantRes: &api.CreateTaskResponseV1{
				UUID: "42",
				Name: "test task",
			},
			wantErr: nil,
			repoConf: func() {
				repo.EXPECT().
					CreateTask(gomock.Any(), gomock.Any()).
					Return(&domain.RepositoryTaskEntity{
						ID:          "42",
						Name:        "test task",
						Description: "test description",
					}, nil)
			},
		},
		{
			name: "Empty User ID",
			dto: &api.CreateTaskRequestV1{
				Name:        "test task",
				Description: "test description",
			},
			wantRes:  nil,
			wantErr:  ErrEmptyUserID,
			repoConf: func() {},
		},
		{
			name: "Short Name",
			dto: &api.CreateTaskRequestV1{
				UserID:      "42",
				Name:        "test",
				Description: "test description",
			},
			wantRes:  nil,
			wantErr:  ErrNameMinLength,
			repoConf: func() {},
		},
		{
			name: "Long Name",
			dto: &api.CreateTaskRequestV1{
				UserID:      "42",
				Name:        strings.Repeat("test task", 42),
				Description: "test description",
			},
			wantRes:  nil,
			wantErr:  ErrNameMaxLength,
			repoConf: func() {},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.repoConf()
			response, err := service.Create(ctx, tt.dto)

			assert.Equal(t, err, tt.wantErr)
			assert.Equal(t, response, tt.wantRes)
		})
	}
}
