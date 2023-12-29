package handles

import (
	"encoding/json"
	"net/http"

	coreMiddleware "github.com/dsbasko/yandex-go-diploma-1/core/rest/middleware"
	"github.com/dsbasko/yandex-go-diploma-1/services/planner/pkg/api"
)

func (h *Handler) GetUndated(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	authPayload := coreMiddleware.GetAuthPayload(r.Context())
	response, err := h.taskService.FindUndated(r.Context(), authPayload.UserID)
	if err != nil {
		h.log.Errorf("taskService.FindUndated: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if len(response.Data) == 0 {
		response = &api.GetTasksResponseV1{
			Data:  []api.GetTaskResponseV1{},
			Total: 0,
		}
	}

	responseBytes, err := json.Marshal(response)
	if err != nil {
		h.log.Errorf("json.Marshal: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if _, err = w.Write(responseBytes); err != nil {
		h.log.Errorf("Write: %v", err)
		w.WriteHeader(http.StatusBadRequest)
	}
}
