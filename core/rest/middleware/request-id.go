package middleware

import (
	"net/http"

	chiMiddleware "github.com/go-chi/chi/v5/middleware"
)

func (m *Middleware) RequestID(next http.Handler) http.Handler {
	m.log.Debug("adding request id to header is enabled")
	fn := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("request-id", chiMiddleware.GetReqID(r.Context()))
		next.ServeHTTP(w, r)
	}

	return chiMiddleware.RequestID(http.HandlerFunc(fn))
}
