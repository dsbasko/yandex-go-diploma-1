package config

import (
	"fmt"
	"sync"

	"github.com/ilyakaznacheev/cleanenv"
)

type config struct {
	Env         string `env:"ENV"`
	ServiceName string `env:"SERVICE_NAME"`

	RmqHost     string `env:"RMQ_HOST"`
	RmqPort     string `env:"RMQ_PORT"`
	RmqAuthUser string `env:"RMQ_AUTH_USER"`
	RmqAuthPass string `env:"RMQ_AUTH_PASS"`
}

var (
	cfg  config
	once sync.Once
	err  error
)

func Init() error {
	once.Do(func() {
		err = cleanenv.ReadEnv(&cfg)
	})

	if err != nil {
		return fmt.Errorf("cleanenv.ReadEnv: %w", err)
	}

	return nil
}

func GetEnv() string {
	return cfg.Env
}

func GetServiceName() string {
	return cfg.ServiceName
}

func GetRmqConnectingString() string {
	return fmt.Sprintf(
		"amqp://%s:%s@%s:%s/",
		cfg.RmqAuthUser,
		cfg.RmqAuthPass,
		cfg.RmqHost,
		cfg.RmqPort,
	)
}
