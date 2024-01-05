package handles

import (
	"encoding/json"
	"net/http"

	"github.com/dsbasko/yandex-go-diploma-1/core/lib"
	"github.com/dsbasko/yandex-go-diploma-1/core/rest/middleware"
	"github.com/dsbasko/yandex-go-diploma-1/services/auth/pkg/api"
)

func (h *Handler) ChangePassword(w http.ResponseWriter, r *http.Request) {
	var dto api.ChangePasswordRequestV1
	defer r.Body.Close()

	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		h.log.Errorf("json.Decode: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	authPayload := middleware.GetAuthPayload(r.Context())
	if authPayload == nil {
		h.log.Error("Unauthorized")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	response, err := h.accountService.ChangePassword(r.Context(), authPayload.UserID, &dto)
	if err != nil {
		h.log.Errorf("accountService.Register: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		if _, err = w.Write([]byte(lib.ErrorsUnwrap(err).Error())); err != nil {
			h.log.Errorf("Write: %v", err)
			w.WriteHeader(http.StatusBadRequest)
		}
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
