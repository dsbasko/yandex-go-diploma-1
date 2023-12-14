package rest

import (
	"context"
	"net/http"

	"github.com/dsbasko/yandex-go-diploma-1/core/logger"
	"github.com/dsbasko/yandex-go-diploma-1/services/auth/internal/config"
	"github.com/dsbasko/yandex-go-diploma-1/services/auth/internal/controllers/rest/handler"
	"github.com/go-chi/chi/v5"
)

func RunServer(ctx context.Context, log *logger.Logger) error {
	handlers := chi.NewRouter()
	// middleware.Inject(log, handler)
	handler.Inject(log, handlers)

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