package rest

import (
	"context"
	"net/http"

	"github.com/dsbasko/yandex-go-diploma-1/core/logger"
	coreMiddleware "github.com/dsbasko/yandex-go-diploma-1/core/rest/middleware"
	"github.com/dsbasko/yandex-go-diploma-1/services/notification/internal/config"
	"github.com/dsbasko/yandex-go-diploma-1/services/notification/internal/controllers/rest/handles"

	"github.com/go-chi/chi/v5"
)

func RunController(ctx context.Context, log *logger.Logger) error {
	handler := chi.NewRouter()
	coreMiddlewares := coreMiddleware.New(log)
	h := handles.New(log)

	handler.Use(coreMiddlewares.RequestID)
	handler.Use(coreMiddlewares.Logger)
	handler.Use(coreMiddlewares.CompressEncoding)
	handler.Use(coreMiddlewares.CompressDecoding)

	handler.Get("/ping", h.Ping)

	routes := handler.Routes()
	for _, route := range routes {
		for handle := range route.Handlers {
			log.Debugf("Mapped [%v] %v route", handle, route.Pattern)
		}
	}

	server := http.Server{
		Addr:         ":3000",
		Handler:      handler,
		ReadTimeout:  config.GetRestReadTimeout(),
		WriteTimeout: config.GetRestWriteTimeout(),
	}

	go func() {
		<-ctx.Done()
		log.Info("shutdown rest server")
		err := server.Shutdown(context.Background())
		if err != nil {
			log.Errorf("server.Shutdown: %v", err)
		}
	}()

	log.Infof("starting rest server at the address: http://localhost/notification")
	return server.ListenAndServe()
}
