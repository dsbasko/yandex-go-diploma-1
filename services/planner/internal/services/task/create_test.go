package task

import (
	"context"
	"strings"
	"testing"

	"github.com/dsbasko/yandex-go-diploma-1/core/errors"
	"github.com/dsbasko/yandex-go-diploma-1/core/logger"
	"github.com/dsbasko/yandex-go-diploma-1/services/planner/internal/entities"
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
		ctx      context.Context
		userID   string
		dto      *api.CreateTaskRequestV1
		wantRes  *api.CreateTaskResponseV1
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
			dto:      &api.CreateTaskRequestV1{},
			wantErr:  ErrEmptyUserID,
			repoConf: func() {},
		},
		{
			name:   "Short Name",
			userID: "42",
			ctx:    ctx,
			dto: &api.CreateTaskRequestV1{
				Name: "",
			},
			wantErr:  ErrNameMinLength,
			repoConf: func() {},
		},
		{
			name:   "Long Name",
			userID: "42",
			ctx:    ctx,
			dto: &api.CreateTaskRequestV1{
				Name: strings.Repeat("test name", 42),
			},
			wantErr:  ErrNameMaxLength,
			repoConf: func() {},
		},
		{
			name:   "Success",
			userID: "42",
			ctx:    ctx,
			dto: &api.CreateTaskRequestV1{
				Name:        "test name",
				Description: "test description",
			},
			wantErr: nil,
			wantRes: &api.CreateTaskResponseV1{
				Name: "test name",
				ID:   "42",
			},
			repoConf: func() {
				repo.EXPECT().
					Create(gomock.Any(), gomock.Any()).
					Return(&entities.RepositoryTaskEntity{
						ID:   "42",
						Name: "test name",
					}, nil)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.repoConf()
			if tt.dto != nil {
				tt.dto.UserID = tt.userID
			}
			response, err := service.Create(tt.ctx, tt.dto)

			assert.Equal(t, response, tt.wantRes)
			assert.Equal(t, errors.Unwrap(err), tt.wantErr)
		})
	}
}
