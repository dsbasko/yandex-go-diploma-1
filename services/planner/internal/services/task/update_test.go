package task

import (
	"context"
	"strings"
	"testing"
	"time"

	"github.com/dsbasko/yandex-go-diploma-1/core/errors"
	"github.com/dsbasko/yandex-go-diploma-1/core/logger"
	"github.com/dsbasko/yandex-go-diploma-1/services/planner/internal/entities"
	"github.com/dsbasko/yandex-go-diploma-1/services/planner/internal/repositories"
	"github.com/dsbasko/yandex-go-diploma-1/services/planner/pkg/api"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestService_UpdateOnce(t *testing.T) {
	ctx := context.Background()
	log := logger.NewMock()
	repo := repositories.NewMock(t)
	service := NewService(log, repo)

	tests := []struct {
		name     string
		ctx      context.Context
		userID   string
		id       string
		dto      *api.UpdateTaskRequestV1
		wantRes  *api.UpdateTaskResponseV1
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
			dto:      &api.UpdateTaskRequestV1{},
			wantErr:  ErrEmptyUserID,
			repoConf: func() {},
		},
		{
			name:     "Empty User ID",
			ctx:      ctx,
			userID:   "42",
			dto:      &api.UpdateTaskRequestV1{},
			wantErr:  ErrEmptyID,
			repoConf: func() {},
		},
		{
			name:   "Short Name",
			userID: "42",
			id:     "42",
			ctx:    ctx,
			dto: &api.UpdateTaskRequestV1{
				Name: "42",
			},
			wantErr:  ErrNameMinLength,
			repoConf: func() {},
		},
		{
			name:   "Long Name",
			userID: "42",
			id:     "42",
			ctx:    ctx,
			dto: &api.UpdateTaskRequestV1{
				Name: strings.Repeat("test name", 42),
			},
			wantErr:  ErrNameMaxLength,
			repoConf: func() {},
		},
		{
			name:   "Success",
			userID: "42",
			id:     "42",
			ctx:    ctx,
			dto: &api.UpdateTaskRequestV1{
				Name:        "test name",
				Description: "test description",
			},
			wantErr: nil,
			wantRes: &api.UpdateTaskResponseV1{
				ID:          "42",
				Name:        "test name",
				Description: "test description",
			},
			repoConf: func() {
				repo.EXPECT().
					UpdateOnce(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
					Return(&entities.RepositoryTaskEntity{
						ID:          "42",
						Name:        "test name",
						Description: "test description",
					}, nil)
			},
		},
		{
			name:   "Success Only Description",
			userID: "42",
			id:     "42",
			ctx:    ctx,
			dto: &api.UpdateTaskRequestV1{
				Description: "test description",
			},
			wantErr: nil,
			wantRes: &api.UpdateTaskResponseV1{
				ID:          "42",
				Name:        "test name",
				Description: "test description",
			},
			repoConf: func() {
				repo.EXPECT().
					UpdateOnce(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
					Return(&entities.RepositoryTaskEntity{
						ID:          "42",
						Name:        "test name",
						Description: "test description",
					}, nil)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.repoConf()
			response, err := service.UpdateOnce(tt.ctx, tt.userID, tt.id, tt.dto)

			assert.Equal(t, response, tt.wantRes)
			assert.Equal(t, errors.Unwrap(err), tt.wantErr)
		})
	}
}

func TestService_UpdateIsArchive(t *testing.T) {
	ctx := context.Background()
	log := logger.NewMock()
	repo := repositories.NewMock(t)
	service := NewService(log, repo)

	tests := []struct {
		name     string
		ctx      context.Context
		userID   string
		id       string
		dto      *api.ChangeIsArchiveRequestV1
		wantRes  *api.ChangeIsArchiveResponseV1
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
			dto:      &api.ChangeIsArchiveRequestV1{},
			wantErr:  ErrEmptyUserID,
			repoConf: func() {},
		},
		{
			name:     "Empty User ID",
			ctx:      ctx,
			userID:   "42",
			dto:      &api.ChangeIsArchiveRequestV1{},
			wantErr:  ErrEmptyID,
			repoConf: func() {},
		},
		{
			name:   "Success To UpdateIsArchive",
			userID: "42",
			id:     "42",
			ctx:    ctx,
			dto: &api.ChangeIsArchiveRequestV1{
				IsArchive: true,
			},
			wantErr: nil,
			wantRes: &api.ChangeIsArchiveResponseV1{
				ID:          "42",
				Name:        "test name",
				Description: "test description",
				IsArchive:   true,
			},
			repoConf: func() {
				repo.EXPECT().
					UpdateIsArchive(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
					Return(&entities.RepositoryTaskEntity{
						ID:          "42",
						Name:        "test name",
						Description: "test description",
						IsArchive:   true,
					}, nil)
			},
		},
		{
			name:   "Success From UpdateIsArchive",
			userID: "42",
			id:     "42",
			ctx:    ctx,
			dto: &api.ChangeIsArchiveRequestV1{
				IsArchive: false,
			},
			wantErr: nil,
			wantRes: &api.ChangeIsArchiveResponseV1{
				ID:          "42",
				Name:        "test name",
				Description: "test description",
				IsArchive:   false,
			},
			repoConf: func() {
				repo.EXPECT().
					UpdateIsArchive(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
					Return(&entities.RepositoryTaskEntity{
						ID:          "42",
						Name:        "test name",
						Description: "test description",
						IsArchive:   false,
					}, nil)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.repoConf()
			response, err := service.UpdateIsArchive(tt.ctx, tt.userID, tt.id, tt.dto)

			assert.Equal(t, response, tt.wantRes)
			assert.Equal(t, errors.Unwrap(err), tt.wantErr)
		})
	}
}

func TestService_UpdateDueDate(t *testing.T) {
	ctx := context.Background()
	log := logger.NewMock()
	repo := repositories.NewMock(t)
	service := NewService(log, repo)
	now := time.Now()
	var zeroTime time.Time

	tests := []struct {
		name     string
		ctx      context.Context
		userID   string
		id       string
		dto      *api.ChangeDueDateRequestV1
		wantRes  *api.ChangeDueDateResponseV1
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
			dto:      &api.ChangeDueDateRequestV1{},
			wantErr:  ErrEmptyUserID,
			repoConf: func() {},
		},
		{
			name:     "Empty User ID",
			ctx:      ctx,
			userID:   "42",
			dto:      &api.ChangeDueDateRequestV1{},
			wantErr:  ErrEmptyID,
			repoConf: func() {},
		},
		{
			name:   "Success Add Due Date",
			userID: "42",
			id:     "42",
			ctx:    ctx,
			dto: &api.ChangeDueDateRequestV1{
				DueDate: now,
			},
			wantErr: nil,
			wantRes: &api.ChangeDueDateResponseV1{
				ID:          "42",
				Name:        "test name",
				Description: "test description",
				DueDate:     now,
			},
			repoConf: func() {
				repo.EXPECT().
					UpdateDueDate(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
					Return(&entities.RepositoryTaskEntity{
						ID:          "42",
						Name:        "test name",
						Description: "test description",
						DueDate:     now,
					}, nil)
			},
		},
		{
			name:    "Success Remove Due Date",
			userID:  "42",
			id:      "42",
			ctx:     ctx,
			dto:     &api.ChangeDueDateRequestV1{},
			wantErr: nil,
			wantRes: &api.ChangeDueDateResponseV1{
				ID:          "42",
				Name:        "test name",
				Description: "test description",
				DueDate:     zeroTime,
			},
			repoConf: func() {
				repo.EXPECT().
					UpdateDueDate(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
					Return(&entities.RepositoryTaskEntity{
						ID:          "42",
						Name:        "test name",
						Description: "test description",
						DueDate:     zeroTime,
					}, nil)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.repoConf()
			response, err := service.UpdateDueDate(tt.ctx, tt.userID, tt.id, tt.dto)

			assert.Equal(t, response, tt.wantRes)
			assert.Equal(t, errors.Unwrap(err), tt.wantErr)
		})
	}
}
