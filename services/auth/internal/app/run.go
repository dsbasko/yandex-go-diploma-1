package app

import (
	"context"
	"fmt"

	"github.com/dsbasko/yandex-go-diploma-1/core/logger"
	"github.com/dsbasko/yandex-go-diploma-1/services/auth/internal/config"
	"github.com/dsbasko/yandex-go-diploma-1/services/auth/internal/controllers/rest"
	"github.com/dsbasko/yandex-go-diploma-1/services/auth/internal/controllers/rmq"
	"github.com/dsbasko/yandex-go-diploma-1/services/auth/internal/repositories"
	"github.com/dsbasko/yandex-go-diploma-1/services/auth/internal/services/account"
	"github.com/dsbasko/yandex-go-diploma-1/services/auth/internal/services/jwt"
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

	repo, err := repositories.NewRepository(ctx)
	if err != nil {
		return fmt.Errorf("repositories.NewRepository: %w", err)
	}

	accountService := account.NewService(log, repo)
	jwtService := jwt.NewService(log, repo)

	rmqCancel, err := rmq.RunController(ctx, log, jwtService)
	defer rmqCancel()
	if err != nil {
		return fmt.Errorf("rmq.RunController: %w", err)
	}

	// HTTP REST триггер
	errRestCh := make(chan error)
	go func() {
		if err = rest.RunController(ctx, log, repo, accountService, jwtService); err != nil {
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
