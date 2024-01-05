package handles

import (
	"encoding/json"
	"net/http"

	coreMiddleware "github.com/dsbasko/yandex-go-diploma-1/core/rest/middleware"
	"github.com/dsbasko/yandex-go-diploma-1/services/planner/pkg/api"
	"github.com/go-chi/chi/v5"
)

func (h *Handler) UpdateIsArchive(w http.ResponseWriter, r *http.Request) {
	var dto api.ArchiveTaskRequestV1
	defer r.Body.Close()

	id := chi.URLParam(r, "id")
	if id == "" {
		h.log.Error(ErrEmptyID)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if r.Header.Get("Content-Type") != ContentTypeApplicationJSON {
		h.log.Error(ErrWrongContentType)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if r.ContentLength == 0 {
		h.log.Error(ErrEmptyBody)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		h.log.Errorf("json.Decode: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	authPayload := coreMiddleware.GetAuthPayload(r.Context())

	response, err := h.taskService.UpdateIsArchive(r.Context(), authPayload.UserID, id, &dto)
	if err != nil {
		h.log.Errorf("taskService.UpdateIsArchive: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	responseBytes, err := json.Marshal(response)
	if err != nil {
		h.log.Errorf("json.Marshal: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", ContentTypeApplicationJSON)
	w.WriteHeader(http.StatusOK)
	if _, err = w.Write(responseBytes); err != nil {
		h.log.Errorf("Write: %v", err)
		w.WriteHeader(http.StatusBadRequest)
	}
}
