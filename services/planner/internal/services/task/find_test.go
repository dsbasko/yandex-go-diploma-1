package task

import (
	"context"
	"testing"

	"github.com/dsbasko/yandex-go-diploma-1/core/lib"
	"github.com/dsbasko/yandex-go-diploma-1/core/logger"
	"github.com/dsbasko/yandex-go-diploma-1/services/planner/internal/entities"
	"github.com/dsbasko/yandex-go-diploma-1/services/planner/internal/repositories"
	"github.com/dsbasko/yandex-go-diploma-1/services/planner/pkg/api"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestService_FindById(t *testing.T) {
	ctx := context.Background()
	log := logger.NewMock()
	repo := repositories.NewMock(t)
	service := NewService(log, repo)

	tests := []struct {
		name     string
		ctx      context.Context
		id       string
		userID   string
		wantRes  *api.GetTaskResponseV1
		wantErr  error
		repoConf func()
	}{
		{
			name:     "Arguments Not Filled",
			wantErr:  ErrArgumentsNotFilled,
			repoConf: func() {},
		},
		{
			name:     "Empty User ID",
			ctx:      ctx,
			wantErr:  ErrEmptyUserID,
			repoConf: func() {},
		},
		{
			name:     "Empty ID",
			ctx:      ctx,
			userID:   "42",
			wantErr:  ErrEmptyID,
			repoConf: func() {},
		},
		{
			name:    "Not Found",
			ctx:     ctx,
			userID:  "42",
			id:      "42",
			wantErr: nil,
			repoConf: func() {
				repo.EXPECT().
					FindByID(gomock.Any(), gomock.Any(), gomock.Any()).
					Return(nil, nil)
			},
		},
		{
			name:    "Found",
			ctx:     ctx,
			userID:  "42",
			id:      "42",
			wantErr: nil,
			wantRes: &api.GetTaskResponseV1{
				ID:          "42",
				UserID:      "42",
				Name:        "test task",
				Description: "test description",
			},
			repoConf: func() {
				repo.EXPECT().
					FindByID(gomock.Any(), gomock.Any(), gomock.Any()).
					Return(&entities.RepositoryTaskEntity{
						ID:          "42",
						UserID:      "42",
						Name:        "test task",
						Description: "test description",
					}, nil)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.repoConf()
			response, err := service.FindByID(tt.ctx, tt.userID, tt.id)

			assert.Equal(t, response, tt.wantRes)
			assert.Equal(t, lib.ErrorsUnwrap(err), tt.wantErr)
		})
	}
}

func TestService_FindToday(t *testing.T) {
	ctx := context.Background()
	log := logger.NewMock()
	repo := repositories.NewMock(t)
	service := NewService(log, repo)

	tests := []struct {
		name     string
		ctx      context.Context
		userID   string
		wantRes  *api.GetTasksResponseV1
		wantErr  error
		repoConf func()
	}{
		{
			name:   "Found Once",
			ctx:    ctx,
			userID: "42",
			wantRes: &api.GetTasksResponseV1{
				Data: []api.GetTaskResponseV1{
					{
						ID:          "42",
						UserID:      "42",
						Name:        "test",
						Description: "test",
					},
				},
				Total: 1,
			},
			wantErr: nil,
			repoConf: func() {
				repo.EXPECT().
					FindByUserIDAndDate(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
					Return(&[]entities.RepositoryTaskEntity{
						{
							ID:          "42",
							UserID:      "42",
							Name:        "test",
							Description: "test",
						},
					}, nil)
			},
		},
		{
			name:   "Found Many",
			ctx:    ctx,
			userID: "42",
			wantRes: &api.GetTasksResponseV1{
				Data: []api.GetTaskResponseV1{
					{
						ID:          "42",
						UserID:      "42",
						Name:        "test",
						Description: "test",
					},
					{
						ID:          "43",
						UserID:      "42",
						Name:        "test",
						Description: "test",
					},
				},
				Total: 2,
			},
			wantErr: nil,
			repoConf: func() {
				repo.EXPECT().
					FindByUserIDAndDate(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
					Return(&[]entities.RepositoryTaskEntity{
						{
							ID:          "42",
							UserID:      "42",
							Name:        "test",
							Description: "test",
						}, {
							ID:          "43",
							UserID:      "42",
							Name:        "test",
							Description: "test",
						},
					}, nil)
			},
		},
		{
			name:    "Not Found",
			ctx:     ctx,
			userID:  "42",
			wantRes: &api.GetTasksResponseV1{},
			wantErr: nil,
			repoConf: func() {
				repo.EXPECT().
					FindByUserIDAndDate(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
					Return(&[]entities.RepositoryTaskEntity{}, nil)
			},
		},
		{
			name:     "Arguments Not Filled",
			ctx:      nil,
			userID:   "",
			wantRes:  nil,
			wantErr:  ErrArgumentsNotFilled,
			repoConf: func() {},
		},
		{
			name:     "Empty UserID",
			ctx:      ctx,
			userID:   "",
			wantRes:  nil,
			wantErr:  ErrEmptyUserID,
			repoConf: func() {},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.repoConf()
			response, err := service.FindToday(tt.ctx, tt.userID)

			if tt.wantRes != nil && response != nil && response.Total != 0 && tt.wantRes.Total != 0 {
				assert.Equal(t, response, tt.wantRes)
			}
			assert.Equal(t, err, tt.wantErr)
		})
	}
}
