package handles

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

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
	"github.com/stretchr/testify/require"
)

func TestHandler_UpdateDueDate(t *testing.T) {
	log := logger.NewMock()
	repo := repositories.NewMock(t)
	taskService := task.NewService(log, repo)
	now := time.Now()

	router := chi.NewRouter()
	h := New(log, repo, taskService)
	router.
		With(coreMiddleware.CheckAuthMock("42")).
		Patch("/{id}/due_date", h.UpdateDueDate)

	ts := httptest.NewServer(router)
	defer ts.Close()

	tests := []struct {
		name           string
		token          string
		contentType    string
		id             string
		body           *api.ChangeDueDateRequestV1
		wantStatusCode int
		wantBody       func() string
		repoCfg        func()
	}{
		{
			name:           "Empty Auth Token",
			wantStatusCode: http.StatusUnauthorized,
			id:             "42",
			repoCfg:        func() {},
			wantBody:       func() string { return "" },
		},
		{
			name:           "Invalid Auth Token",
			token:          "43",
			id:             "42",
			wantStatusCode: http.StatusUnauthorized,
			repoCfg:        func() {},
			wantBody:       func() string { return "" },
		},
		{
			name:           "Wrong Content-Type",
			token:          "42",
			id:             "42",
			wantStatusCode: http.StatusBadRequest,
			repoCfg:        func() {},
			wantBody:       func() string { return "" },
		},
		{
			name:           "Empty Body",
			token:          "42",
			id:             "42",
			contentType:    ContentTypeApplicationJSON,
			wantStatusCode: http.StatusBadRequest,
			body:           nil,
			repoCfg:        func() {},
			wantBody:       func() string { return "" },
		},
		{
			name:        "Success",
			token:       "42",
			id:          "42",
			contentType: ContentTypeApplicationJSON,
			body: &api.ChangeDueDateRequestV1{
				DueDate: now,
			},
			wantStatusCode: http.StatusOK,
			wantBody: func() string {
				response, _ := json.Marshal(api.UpdateTaskResponseV1{
					ID:          "42",
					Name:        "test name",
					Description: "test description",
					DueDate:     now,
				})
				return string(response)
			},
			repoCfg: func() {
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
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.repoCfg()
			var bodyBytes []byte
			var err error
			if tt.body != nil {
				bodyBytes, err = json.Marshal(tt.body)
			}
			require.Nil(t, err)

			resp, body := test.Request(t, ts, &test.RequestArgs{
				Method:      "PATCH",
				Path:        fmt.Sprintf("/%s/due_date", tt.id),
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