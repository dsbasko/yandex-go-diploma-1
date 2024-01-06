package middleware

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/dsbasko/yandex-go-diploma-1/core/logger"
	"github.com/dsbasko/yandex-go-diploma-1/core/rmq"
	apiAuth "github.com/dsbasko/yandex-go-diploma-1/services/auth/pkg/api"
	chiMiddleware "github.com/go-chi/chi/v5/middleware"
	"github.com/rabbitmq/amqp091-go"
)

type AuthKey string

var CheckAuthKey AuthKey = "auth-payload"

func CheckAuth(log *logger.Logger, conn *rmq.Connector) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log = log.With(RequestIDKey, chiMiddleware.GetReqID(r.Context()))

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
			dtoBytes, err := json.Marshal(apiAuth.JWTValidationRequestV1{Token: token})
			if err != nil {
				w.WriteHeader(http.StatusUnauthorized)
				log.Warnf("Unauthorized. json.Marshal: %v", err)
				return
			}

			body, err := conn.SimplePublishAndWaitResponse(r.Context(), &rmq.SimplePublisherConfig{
				Exchange:  apiAuth.AMQPExchange,
				Key:       apiAuth.AMQPJWTValidationKey,
				Mandatory: true,
				Msg: amqp091.Publishing{
					Body: dtoBytes,
				},
			})
			if err != nil {
				w.WriteHeader(http.StatusUnauthorized)
				log.Warnf("Unauthorized. conn.SimplePublishAndWaitResponse: %v", err)
				return
			}

			var validate apiAuth.JWTValidationResponseV1
			if err = json.Unmarshal(body, &validate); err != nil {
				w.WriteHeader(http.StatusUnauthorized)
				log.Warnf("Unauthorized. json.Unmarshal: %v", err)
				return
			}

			if !validate.IsValid {
				w.WriteHeader(http.StatusUnauthorized)
				log.Warnf("Unauthorized. %+v", validate)
				return
			}

			ctx := context.WithValue(r.Context(), CheckAuthKey, validate.Payload)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func CheckAuthMock(mockToken string) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			if !strings.HasPrefix(authHeader, "Bearer ") {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			token := strings.TrimPrefix(authHeader, "Bearer ")
			if token != mockToken {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			validate := apiAuth.JWTValidationResponseV1{
				IsValid: true,
				Payload: &apiAuth.JWTPayloadV1{
					UserID:    "mock",
					Username:  "mock",
					FirstName: "mock",
					LastName:  "mock",
				},
			}

			ctx := context.WithValue(r.Context(), CheckAuthKey, validate.Payload)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func GetAuthPayload(ctx context.Context) *apiAuth.JWTPayloadV1 {
	if ctx == nil {
		return nil
	}

	if payload, ok := ctx.Value(CheckAuthKey).(*apiAuth.JWTPayloadV1); ok {
		return payload
	}

	return nil
}
