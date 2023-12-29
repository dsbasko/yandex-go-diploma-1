package handles

import (
	"encoding/json"
	"net/http"

	coreMiddleware "github.com/dsbasko/yandex-go-diploma-1/core/rest/middleware"
	"github.com/go-chi/chi/v5"
)

func (h *Handler) GetByID(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	id := chi.URLParam(r, "id")
	authPayload := coreMiddleware.GetAuthPayload(r.Context())
	response, err := h.taskService.FindByID(r.Context(), authPayload.UserID, id)
	if err != nil {
		h.log.Errorf("taskService.FindToday: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if response == nil {
		w.WriteHeader(http.StatusNoContent)
		return
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
