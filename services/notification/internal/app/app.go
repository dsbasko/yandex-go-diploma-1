package app

import (
	"context"
	"fmt"

	"github.com/dsbasko/yandex-go-diploma-1/core/logger"
	"github.com/dsbasko/yandex-go-diploma-1/services/notification/internal/config"
	"github.com/dsbasko/yandex-go-diploma-1/services/notification/internal/controllers/rest"
)

func Run() error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := config.Init()
	if err != nil {
		return fmt.Errorf("config.Init: %w", err)
	}

	log, err := logger.NewLogger(config.GetEnv(), config.GetServiceName())
	if err != nil {
		return fmt.Errorf("logger.NewLogger: %w", err)
	}

	// HTTP REST триггер
	errRestCh := make(chan error)
	go func() {
		if err = rest.RunController(ctx, log); err != nil {
			errRestCh <- fmt.Errorf("rest.Run: %v", err)
		}
	}()

	select {
	case <-ctx.Done():
		return fmt.Errorf("ctx.Done: %w", ctx.Err())
	case err = <-errRestCh:
		return err
	}
}
