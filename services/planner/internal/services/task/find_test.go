package task

import (
	"context"
	"testing"

	"github.com/dsbasko/yandex-go-diploma-1/core/logger"
	"github.com/dsbasko/yandex-go-diploma-1/services/planner/internal/domain"
	"github.com/dsbasko/yandex-go-diploma-1/services/planner/internal/repositories"
	"github.com/dsbasko/yandex-go-diploma-1/services/planner/pkg/api"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestService_FindToday(t *testing.T) {
	ctx := context.Background()
	log := logger.NewMock()
	repo := repositories.NewMock(t)
	service := NewService(log, repo)

	tests := []struct {
		name     string
		ctx      context.Context
		userID   string
		wantRes  *api.GetTodayResponseV1
		wantErr  error
		repoConf func()
	}{
		{
			name:   "Found Once",
			ctx:    ctx,
			userID: "42",
			wantRes: &api.GetTodayResponseV1{
				Data: []api.GetTodayResponseV1Data{
					{
						UUID:        "42",
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
					Return(&[]domain.RepositoryTaskEntity{
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
			wantRes: &api.GetTodayResponseV1{
				Data: []api.GetTodayResponseV1Data{
					{
						UUID:        "42",
						UserID:      "42",
						Name:        "test",
						Description: "test",
					},
					{
						UUID:        "43",
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
					Return(&[]domain.RepositoryTaskEntity{
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
			wantRes: &api.GetTodayResponseV1{},
			wantErr: nil,
			repoConf: func() {
				repo.EXPECT().
					FindByUserIDAndDate(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
					Return(&[]domain.RepositoryTaskEntity{}, nil)
			},
		},
		{
			name:     "Empty ArgumentsNotFilled",
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
