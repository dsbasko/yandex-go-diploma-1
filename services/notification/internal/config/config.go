package config

import (
	"fmt"
	"sync"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type config struct {
	Env         string `env:"ENV"`
	ServiceName string `env:"SERVICE_NAME"`

	RestReadTimeout  int `env:"REST_READ_TIMEOUT"`
	RestWriteTimeout int `env:"REST_WRITE_TIMEOUT"`
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

func GetRestReadTimeout() time.Duration {
	return time.Duration(cfg.RestReadTimeout) * time.Millisecond
}

func GetRestWriteTimeout() time.Duration {
	return time.Duration(cfg.RestWriteTimeout) * time.Millisecond
}
