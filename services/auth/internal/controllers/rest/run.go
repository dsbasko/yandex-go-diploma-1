package rest

import (
	"context"
	"net/http"

	"github.com/dsbasko/yandex-go-diploma-1/core/logger"
	coreMiddleware "github.com/dsbasko/yandex-go-diploma-1/core/rest/middleware"
	"github.com/dsbasko/yandex-go-diploma-1/services/auth/internal/config"
	"github.com/dsbasko/yandex-go-diploma-1/services/auth/internal/controllers/rest/handles"
	"github.com/dsbasko/yandex-go-diploma-1/services/auth/internal/controllers/rest/middlewares"
	"github.com/dsbasko/yandex-go-diploma-1/services/auth/internal/interfaces"
	"github.com/dsbasko/yandex-go-diploma-1/services/auth/internal/services/account"
	"github.com/dsbasko/yandex-go-diploma-1/services/auth/internal/services/jwt"
	"github.com/go-chi/chi/v5"
)

func RunController(
	ctx context.Context,
	log *logger.Logger,
	repo interfaces.Repository,
	accountService *account.Service,
	jwtService *jwt.Service,
) error {
	handler := chi.NewRouter()
	coreMiddlewares := coreMiddleware.New(log)
	h := handles.New(log, repo, accountService, jwtService)

	handler.Use(coreMiddlewares.RequestID)
	handler.Use(coreMiddlewares.Logger)
	handler.Use(coreMiddlewares.CompressEncoding)
	handler.Use(coreMiddlewares.CompressDecoding)

	handler.Get("/ping", h.Ping)
	handler.Post("/register", h.Register)
	handler.Post("/login", h.Login)
	handler.With(middlewares.CheckAuth(log, jwtService)).Post("/change_password", h.ChangePassword)

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

	log.Infof("starting rest server at the address: http://localhost/auth")
	return server.ListenAndServe()
}
