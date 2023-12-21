package handles

import (
	"net/http"
)

func (h *Handler) Ping(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write([]byte("pong")); err != nil {
		h.log.Errorf("http.ResponseWriter: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}
