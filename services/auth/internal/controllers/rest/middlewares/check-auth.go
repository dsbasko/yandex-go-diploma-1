package middlewares

import (
	"context"
	"net/http"
	"strings"

	"github.com/dsbasko/yandex-go-diploma-1/core/logger"
	coreMiddleware "github.com/dsbasko/yandex-go-diploma-1/core/rest/middleware"
	"github.com/dsbasko/yandex-go-diploma-1/services/auth/internal/services/jwt"
	"github.com/dsbasko/yandex-go-diploma-1/services/auth/pkg/api"
	chiMiddleware "github.com/go-chi/chi/v5/middleware"
)

func CheckAuth(log *logger.Logger, jwtService *jwt.Service) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log = log.With(coreMiddleware.RequestIDKey, chiMiddleware.GetReqID(r.Context()))

			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				w.WriteHeader(http.StatusUnauthorized)
				log.Warn("Unauthorized. Empty auth header")
				return
			}

			if !strings.HasPrefix(authHeader, "Bearer ") {
				w.WriteHeader(http.StatusUnauthorized)
				log.Warn("Unauthorized. Bearer not found")
				return
			}

			token := strings.TrimPrefix(authHeader, "Bearer ")
			validate, err := jwtService.Validation(r.Context(), &api.JWTValidationRequestV1{Token: token})
			if err != nil || !validate.IsValid {
				w.WriteHeader(http.StatusUnauthorized)
				log.Warn("Unauthorized. Token not valid")
				return
			}

			if !validate.IsValid {
				w.WriteHeader(http.StatusUnauthorized)
				log.Warnf("Unauthorized. %+v", validate)
				return
			}

			ctx := context.WithValue(r.Context(), coreMiddleware.CheckAuthKey, validate.Payload)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
