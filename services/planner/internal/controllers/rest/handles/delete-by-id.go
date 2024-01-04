package handles

import (
	"net/http"

	coreMiddleware "github.com/dsbasko/yandex-go-diploma-1/core/rest/middleware"
	"github.com/go-chi/chi/v5"
)

func (h *Handler) DeleteByID(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	id := chi.URLParam(r, "id")
	authPayload := coreMiddleware.GetAuthPayload(r.Context())
	err := h.taskService.DeleteByID(r.Context(), authPayload.UserID, id)
	if err != nil {
		h.log.Errorf("taskService.DeleteByID: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}
