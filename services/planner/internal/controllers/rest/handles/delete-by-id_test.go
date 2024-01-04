package handles

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/dsbasko/yandex-go-diploma-1/core/logger"
	coreMiddleware "github.com/dsbasko/yandex-go-diploma-1/core/rest/middleware"
	"github.com/dsbasko/yandex-go-diploma-1/core/test"
	"github.com/dsbasko/yandex-go-diploma-1/services/planner/internal/repositories"
	"github.com/dsbasko/yandex-go-diploma-1/services/planner/internal/services/task"
	"github.com/go-chi/chi/v5"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestHandler_DeleteByID(t *testing.T) {
	log := logger.NewMock()
	repo := repositories.NewMock(t)
	taskService := task.NewService(log, repo)

	router := chi.NewRouter()
	h := New(log, repo, taskService)
	router.With(coreMiddleware.CheckAuthMock("42")).Delete("/{id}", h.DeleteByID)

	ts := httptest.NewServer(router)
	defer ts.Close()

	tests := []struct {
		name           string
		token          string
		id             string
		wantStatusCode int
		repoCfg        func()
	}{
		{
			name:           "Empty Auth Token",
			id:             "42",
			wantStatusCode: http.StatusUnauthorized,
			repoCfg:        func() {},
		},
		{
			name:           "Invalid Auth Token",
			id:             "42",
			token:          "43",
			wantStatusCode: http.StatusUnauthorized,
			repoCfg:        func() {},
		},
		{
			name:           "Success",
			id:             "42",
			token:          "42",
			wantStatusCode: http.StatusOK,
			repoCfg: func() {
				repo.EXPECT().
					DeleteByID(gomock.Any(), gomock.Any(), gomock.Any()).
					Return(nil)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.repoCfg()
			resp, _ := test.Request(t, ts, &test.RequestArgs{
				Method:   "DELETE",
				Path:     "/" + tt.id,
				JWTToken: tt.token,
			})
			defer resp.Body.Close()

			assert.Equal(t, resp.StatusCode, tt.wantStatusCode)
		})
	}
}
