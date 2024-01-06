package handles

import (
	"net/http"
)

func (h *Handler) Ping(w http.ResponseWriter, r *http.Request) {
	if err := h.repo.Ping(r.Context()); err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	w.WriteHeader(http.StatusOK)
	if _, err := w.Write([]byte("pong")); err != nil {
		h.log.Errorf("http.ResponseWriter: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}
