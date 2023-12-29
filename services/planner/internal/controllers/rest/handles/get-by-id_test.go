package handles

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/dsbasko/yandex-go-diploma-1/core/logger"
	coreMiddleware "github.com/dsbasko/yandex-go-diploma-1/core/rest/middleware"
	"github.com/dsbasko/yandex-go-diploma-1/core/test"
	"github.com/dsbasko/yandex-go-diploma-1/services/planner/internal/entities"
	"github.com/dsbasko/yandex-go-diploma-1/services/planner/internal/repositories"
	"github.com/dsbasko/yandex-go-diploma-1/services/planner/internal/services/task"
	"github.com/dsbasko/yandex-go-diploma-1/services/planner/pkg/api"
	"github.com/go-chi/chi/v5"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestHandler_GetByID(t *testing.T) {
	log := logger.NewMock()
	repo := repositories.NewMock(t)
	taskService := task.NewService(log, repo)

	router := chi.NewRouter()
	h := New(log, repo, taskService)
	router.With(coreMiddleware.CheckAuthMock("42")).Get("/{id}", h.GetByID)

	ts := httptest.NewServer(router)
	defer ts.Close()

	tests := []struct {
		name           string
		token          string
		contentType    string
		id             string
		wantStatusCode int
		wantBody       func() string
		repoCfg        func()
	}{
		{
			name:           "Empty Auth Token",
			id:             "42",
			wantStatusCode: http.StatusUnauthorized,
			repoCfg:        func() {},
			wantBody:       func() string { return "" },
		},
		{
			name:           "Invalid Auth Token",
			id:             "42",
			token:          "43",
			wantStatusCode: http.StatusUnauthorized,
			repoCfg:        func() {},
			wantBody:       func() string { return "" },
		},
		{
			name:           "Not Found",
			id:             "42",
			token:          "42",
			wantStatusCode: http.StatusNoContent,
			repoCfg: func() {
				repo.EXPECT().
					FindByID(gomock.Any(), gomock.Any(), gomock.Any()).
					Return(nil, nil)
			},
			wantBody: func() string { return "" },
		},
		{
			name:           "Found",
			id:             "42",
			token:          "42",
			wantStatusCode: http.StatusOK,
			repoCfg: func() {
				repo.EXPECT().
					FindByID(gomock.Any(), gomock.Any(), gomock.Any()).
					Return(&entities.RepositoryTaskEntity{
						ID:          "42",
						UserID:      "42",
						Name:        "test task",
						Description: "test description",
					}, nil)
			},
			wantBody: func() string {
				response, _ := json.Marshal(api.GetTaskResponseV1{
					ID:          "42",
					UserID:      "42",
					Name:        "test task",
					Description: "test description",
				})
				return string(response)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.repoCfg()
			resp, body := test.Request(t, ts, &test.RequestArgs{
				Method:   "GET",
				Path:     "/" + tt.id,
				JWTToken: tt.token,
			})
			defer resp.Body.Close()

			assert.Equal(t, resp.StatusCode, tt.wantStatusCode)
			assert.Equal(t, body, tt.wantBody())
		})
	}
}
