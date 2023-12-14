package app

import (
	"fmt"
	"time"

	"github.com/dsbasko/yandex-go-diploma-1/core/logger"
	"github.com/dsbasko/yandex-go-diploma-1/services/auth/internal/config"
)

func Run() error {
	err := config.Init()
	if err != nil {
		return fmt.Errorf("config.Init: %w", err)
	}

	log, err := logger.NewLogger(config.GetEnv(), config.GetServiceName())
	if err != nil {
		return fmt.Errorf("logger.NewLogger: %w", err)
	}
	// TODO Don`t forget to remove this нахрен
	_ = log

	// TODO Don`t forget to remove this нахрен
	time.Sleep(24 * time.Hour)

	return nil
}
