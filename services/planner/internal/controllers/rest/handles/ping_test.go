package handles

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/dsbasko/yandex-go-diploma-1/core/logger"
	"github.com/dsbasko/yandex-go-diploma-1/core/test"
	"github.com/dsbasko/yandex-go-diploma-1/services/planner/internal/repositories"
	"github.com/dsbasko/yandex-go-diploma-1/services/planner/internal/services/task"
	"github.com/go-chi/chi/v5"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestHandler_Ping(t *testing.T) {
	log := logger.NewMock()
	repo := repositories.NewMock(t)
	taskService := task.NewService(log, repo)

	router := chi.NewRouter()
	h := New(log, repo, taskService)
	router.Get("/ping", h.Ping)

	ts := httptest.NewServer(router)
	defer ts.Close()

	tests := []struct {
		name           string
		wantStatusCode int
		wantBody       string
		repoCfg        func()
	}{
		{
			name:           "Success",
			wantStatusCode: http.StatusOK,
			repoCfg: func() {
				repo.EXPECT().Ping(gomock.Any()).Return(nil)
			},
			wantBody: "pong",
		},
		{
			name:           "Fail",
			wantStatusCode: http.StatusBadRequest,
			repoCfg: func() {
				repo.EXPECT().Ping(gomock.Any()).Return(errors.New(""))
			},
			wantBody: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.repoCfg()
			resp, body := test.Request(t, ts, &test.RequestArgs{
				Method: "GET",
				Path:   "/ping",
			})
			defer resp.Body.Close()

			assert.Equal(t, resp.StatusCode, tt.wantStatusCode)
			assert.Equal(t, body, tt.wantBody)
		})
	}
}
