package app

import (
	"fmt"
	"time"

	"github.com/dsbasko/yandex-go-diploma-1/core/logger"
)

func Run() error {
	// TODO указать переменную окружения из конфига
	log, err := logger.NewLogger("dev", "auth")
	if err != nil {
		return fmt.Errorf("logger.NewLogger: %w", err)
	}
	// TODO Don`t forget to remove this нахрен
	_ = log

	// TODO Don`t forget to remove this нахрен
	time.Sleep(24 * time.Hour)

	return nil
}
