package app

import (
	"context"
	"fmt"

	"github.com/dsbasko/yandex-go-diploma-1/core/logger"
	"github.com/dsbasko/yandex-go-diploma-1/services/planner/internal/adapters/rmq"
	"github.com/dsbasko/yandex-go-diploma-1/services/planner/internal/config"
	"github.com/dsbasko/yandex-go-diploma-1/services/planner/internal/controllers/rest"
	"github.com/dsbasko/yandex-go-diploma-1/services/planner/internal/repositories"
	"github.com/dsbasko/yandex-go-diploma-1/services/planner/internal/services/task"
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

	adapter, err := rmq.RunAdapter(ctx, log)
	if err != nil {
		return fmt.Errorf("rmq.RunAdapter: %w", err)
	}

	taskService := task.NewService(log, repo)

	response, err := taskService.FindToday(ctx,
		"657d98c7-0eb3-473d-a471-94c1214fde40",
	)
	fmt.Println("response", response)

	// HTTP REST триггер
	errRestCh := make(chan error)
	go func() {
		if err = rest.RunController(ctx, log, repo, adapter, taskService); err != nil {
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
