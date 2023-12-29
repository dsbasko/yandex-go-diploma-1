package rest

import (
	"context"
	"net/http"

	"github.com/dsbasko/yandex-go-diploma-1/core/logger"
	coreMiddleware "github.com/dsbasko/yandex-go-diploma-1/core/rest/middleware"
	"github.com/dsbasko/yandex-go-diploma-1/core/rmq"
	"github.com/dsbasko/yandex-go-diploma-1/services/planner/internal/config"
	"github.com/dsbasko/yandex-go-diploma-1/services/planner/internal/controllers/rest/handles"
	"github.com/dsbasko/yandex-go-diploma-1/services/planner/internal/interfaces"
	"github.com/dsbasko/yandex-go-diploma-1/services/planner/internal/services/task"

	"github.com/go-chi/chi/v5"
)

func RunController(
	ctx context.Context,
	log *logger.Logger,
	repo interfaces.Repository,
	adapter *rmq.Connector,
	taskService *task.Service,
) error {
	handler := chi.NewRouter()
	coreMiddlewares := coreMiddleware.New(log)
	h := handles.New(log, repo, taskService)

	handler.Use(coreMiddlewares.RequestID)
	handler.Use(coreMiddlewares.Logger)
	handler.Use(coreMiddlewares.CompressEncoding)
	handler.Use(coreMiddlewares.CompressDecoding)

	handler.Get("/ping", h.Ping)
	handler.With(coreMiddleware.CheckAuth(log, adapter)).Post("/", h.CreateTask)
	handler.With(coreMiddleware.CheckAuth(log, adapter)).Get("/{id}", h.GetByID)
	handler.With(coreMiddleware.CheckAuth(log, adapter)).Get("/today", h.GetToday)
	handler.With(coreMiddleware.CheckAuth(log, adapter)).Get("/archive", h.GetArchive)

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

	log.Infof("starting rest server at the address: http://localhost/planner")
	return server.ListenAndServe()
}
