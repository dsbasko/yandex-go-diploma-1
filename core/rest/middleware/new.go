package middleware

import "github.com/dsbasko/yandex-go-diploma-1/core/logger"

type Middleware struct {
	log *logger.Logger
}

func New(log *logger.Logger) Middleware {
	return Middleware{log: log}
}
