package config

import (
	"fmt"
	"sync"

	"github.com/ilyakaznacheev/cleanenv"
)

type config struct {
	Env         string `env:"ENV"`
	ServiceName string `env:"SERVICE_NAME"`
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
