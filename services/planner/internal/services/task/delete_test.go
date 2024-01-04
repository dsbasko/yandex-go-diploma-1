package task

import (
	"context"
	"testing"

	"github.com/dsbasko/yandex-go-diploma-1/core/lib"
	"github.com/dsbasko/yandex-go-diploma-1/core/logger"
	"github.com/dsbasko/yandex-go-diploma-1/services/planner/internal/repositories"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestService_DeleteByID(t *testing.T) {
	ctx := context.Background()
	log := logger.NewMock()
	repo := repositories.NewMock(t)
	service := NewService(log, repo)

	tests := []struct {
		name     string
		ctx      context.Context
		id       string
		userID   string
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
			name:    "Success",
			ctx:     ctx,
			userID:  "42",
			id:      "42",
			wantErr: nil,
			repoConf: func() {
				repo.EXPECT().
					DeleteByID(gomock.Any(), gomock.Any(), gomock.Any()).
					Return(nil)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.repoConf()
			err := service.DeleteByID(tt.ctx, tt.userID, tt.id)
			assert.Equal(t, lib.ErrorsUnwrap(err), tt.wantErr)
		})
	}
}
