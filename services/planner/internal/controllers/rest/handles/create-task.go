package handles

import (
	"encoding/json"
	"net/http"

	coreMiddleware "github.com/dsbasko/yandex-go-diploma-1/core/rest/middleware"
	"github.com/dsbasko/yandex-go-diploma-1/services/planner/pkg/api"
)

func (h *Handler) CreateTask(w http.ResponseWriter, r *http.Request) {
	var dto api.CreateTaskRequestV1
	defer r.Body.Close()

	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		h.log.Errorf("json.Decode: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	authPayload := coreMiddleware.GetAuthPayload(r.Context())
	dto.UserID = authPayload.UserID

	response, err := h.taskService.Create(r.Context(), &dto)
	if err != nil {
		h.log.Errorf("taskService.Register: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	responseBytes, err := json.Marshal(response)
	if err != nil {
		h.log.Errorf("json.Marshal: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if _, err = w.Write(responseBytes); err != nil {
		h.log.Errorf("Write: %v", err)
		w.WriteHeader(http.StatusBadRequest)
	}
}
