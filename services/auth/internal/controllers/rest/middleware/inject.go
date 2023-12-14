package middleware

import (
	"github.com/dsbasko/yandex-go-diploma-1/core/logger"
	"github.com/dsbasko/yandex-go-diploma-1/core/rest/middleware"
	"github.com/go-chi/chi/v5"
)

func Inject(log *logger.Logger, handler *chi.Mux) {
	mw := middleware.New(log)

	handler.Use(mw.CompressEncoding)
	handler.Use(mw.CompressDecoding)
	handler.Use(mw.Logger)
	handler.Use(mw.RequestID)
}
