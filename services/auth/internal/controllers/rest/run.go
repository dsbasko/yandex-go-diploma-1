package rest

import (
	"context"
	"net/http"

	"github.com/dsbasko/yandex-go-diploma-1/core/logger"
	"github.com/dsbasko/yandex-go-diploma-1/services/auth/internal/config"
	"github.com/dsbasko/yandex-go-diploma-1/services/auth/internal/controllers/rest/handler"
	"github.com/dsbasko/yandex-go-diploma-1/services/auth/internal/controllers/rest/middleware"
	"github.com/dsbasko/yandex-go-diploma-1/services/auth/internal/services/account"
	"github.com/dsbasko/yandex-go-diploma-1/services/auth/internal/services/jwt"
	"github.com/go-chi/chi/v5"
)

func RunController(
	ctx context.Context,
	log *logger.Logger,
	accountService *account.Service,
	jwtService *jwt.Service,
) error {
	handlers := chi.NewRouter()

	middleware.Inject(log, handlers)
	handler.Inject(log, handlers, accountService, jwtService)

	server := http.Server{
		Addr:         ":3000",
		Handler:      handlers,
		ReadTimeout:  config.GetRestReadTimeout(),
		WriteTimeout: config.GetRestWriteTimeout(),
	}

	// Если корневой контекст закрывается, то тормозим сервер
	go func() {
		<-ctx.Done()
		log.Info("shutdown rest server")
		err := server.Shutdown(context.Background())
		if err != nil {
			log.Errorf("server.Shutdown: %v", err)
		}
	}()

	log.Infof("starting rest server at the address: http://localhost/auth")
	return server.ListenAndServe()
}
