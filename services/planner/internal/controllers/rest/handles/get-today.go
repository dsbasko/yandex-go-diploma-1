package handles

import (
	"encoding/json"
	"net/http"

	coreMiddleware "github.com/dsbasko/yandex-go-diploma-1/core/rest/middleware"
)

func (h *Handler) GetToday(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	authPayload := coreMiddleware.GetAuthPayload(r.Context())
	response, err := h.taskService.FindToday(r.Context(), authPayload.UserID)
	if err != nil {
		h.log.Errorf("taskService.FindToday: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var responseBytes []byte
	if len(response.Data) == 0 {
		responseBytes = []byte("[]")
	} else {
		responseBytes, err = json.Marshal(response)
		if err != nil {
			h.log.Errorf("json.Marshal: %v", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if _, err = w.Write(responseBytes); err != nil {
		h.log.Errorf("Write: %v", err)
		w.WriteHeader(http.StatusBadRequest)
	}
}
