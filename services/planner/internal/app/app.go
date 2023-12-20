package app

import (
	"context"
	"fmt"
	"time"

	"github.com/dsbasko/yandex-go-diploma-1/core/logger"
	"github.com/dsbasko/yandex-go-diploma-1/services/planner/internal/adapters/amqp"
	"github.com/dsbasko/yandex-go-diploma-1/services/planner/internal/config"
)

const token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MDMwNzAwMzAsImlhdCI6MTcwMjk4MzYzMCwiSldUUGF5bG9hZCI6eyJ1c2VyX2lkIjoiZDY2YTNmMzMtNjJjZS00OTE1LWE4N2EtOGM0ZDhjZDFiNWUyIiwidXNlcm5hbWUiOiJhZG1pbiIsImZpcnN0X25hbWUiOiJEbWl0cml5IiwibGFzdF9uYW1lIjoiQmFzZW5rbyJ9fQ.copn-8boa08PWZrGCq7QA-0sbiy4bFSAK4biWxcvIRw" //nolint:lll,gosec

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

	rmqAdapter, err := amqp.RunAdapter(log)
	if err != nil {
		return fmt.Errorf("amqp.RunAdapter: %w", err)
	}

	/*
		TODO {
		Код с response можно удалять.
		Сделан он для проверки и демонстрации того как будет работать общение между сервисами.
		Также буду думать над тем как упростить это общение
		↓↓↓↓↓↓↓↓↓↓↓↓↓
	*/
	if _, err = rmqAdapter.JWTValidate(ctx, token); err != nil {
		log.Errorf("rmqAdapter.JWTValidate: %v", err)
	}
	/* ↑↑↑↑↑↑↑↑↑↑↑↑ */

	// TODO Don't forget to remove this нахрен
	time.Sleep(24 * time.Hour)

	return nil
}
