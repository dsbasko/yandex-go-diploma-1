package handles

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/dsbasko/yandex-go-diploma-1/core/logger"
	coreMiddleware "github.com/dsbasko/yandex-go-diploma-1/core/rest/middleware"
	"github.com/dsbasko/yandex-go-diploma-1/core/test"
	"github.com/dsbasko/yandex-go-diploma-1/services/planner/internal/domain"
	"github.com/dsbasko/yandex-go-diploma-1/services/planner/internal/repositories"
	"github.com/dsbasko/yandex-go-diploma-1/services/planner/internal/services/task"
	"github.com/dsbasko/yandex-go-diploma-1/services/planner/pkg/api"
	"github.com/go-chi/chi/v5"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHandler_CreateTask(t *testing.T) {
	log := logger.NewMock()
	repo := repositories.NewMock(t)
	taskService := task.NewService(log, repo)

	router := chi.NewRouter()
	h := New(log, repo, taskService)
	router.
		With(coreMiddleware.CheckAuthMock("42")).
		Post("/", h.CreateTask)

	ts := httptest.NewServer(router)
	defer ts.Close()

	tests := []struct {
		name           string
		token          string
		contentType    string
		body           api.CreateTaskRequestV1
		wantStatusCode int
		wantBody       func() string
		repoCfg        func()
	}{
		{
			name:           "Empty Auth Token",
			wantStatusCode: http.StatusUnauthorized,
			repoCfg:        func() {},
			wantBody:       func() string { return "" },
		},
		{
			name:           "Invalid Auth Token",
			token:          "43",
			wantStatusCode: http.StatusUnauthorized,
			repoCfg:        func() {},
			wantBody:       func() string { return "" },
		},
		{
			name:           "Wrong Content-Type",
			token:          "42",
			contentType:    "application/json",
			wantStatusCode: http.StatusBadRequest,
			repoCfg:        func() {},
			wantBody:       func() string { return "" },
		},
		{
			name:           "Empty Body",
			token:          "42",
			contentType:    "application/json",
			wantStatusCode: http.StatusBadRequest,
			repoCfg:        func() {},
			wantBody:       func() string { return "" },
		},
		{
			name:        "Success",
			token:       "42",
			contentType: "application/json",
			body: api.CreateTaskRequestV1{
				Name:        "test task",
				Description: "test description",
			},
			wantStatusCode: http.StatusCreated,
			wantBody: func() string {
				response, _ := json.Marshal(api.CreateTaskResponseV1{
					UUID: "42",
					Name: "test task",
				})
				return string(response)
			},
			repoCfg: func() {
				repo.EXPECT().
					CreateTask(gomock.Any(), gomock.Any()).
					Return(&domain.RepositoryTaskEntity{
						ID:   "42",
						Name: "test task",
					}, nil)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.repoCfg()
			bodyBytes, err := json.Marshal(tt.body)
			require.Nil(t, err)

			resp, body := test.Request(t, ts, &test.RequestArgs{
				Method:      "POST",
				Path:        "/",
				JWTToken:    tt.token,
				ContentType: tt.contentType,
				Body:        bodyBytes,
			})
			defer resp.Body.Close()

			assert.Equal(t, resp.StatusCode, tt.wantStatusCode)
			assert.Equal(t, body, tt.wantBody())
		})
	}
}
