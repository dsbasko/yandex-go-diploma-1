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

	RmqHost     string `env:"RMQ_HOST"`
	RmqPort     string `env:"RMQ_PORT"`
	RmqAuthUser string `env:"RMQ_AUTH_USER"`
	RmqAuthPass string `env:"RMQ_AUTH_PASS"`

	PsqlHost     string `env:"PSQL_HOST"`
	PsqlPort     string `env:"PSQL_PORT"`
	PsqlUser     string `env:"PSQL_USER"`
	PsqlPass     string `env:"PSQL_PASS"`
	PsqlDB       string `env:"PSQL_DB"`
	PsqlMaxPools int32  `env:"PSQL_MAX_POOLS"`
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

func GetRmqConnectingString() string {
	return fmt.Sprintf(
		"amqp://%s:%s@%s:%s/",
		cfg.RmqAuthUser,
		cfg.RmqAuthPass,
		cfg.RmqHost,
		cfg.RmqPort,
	)
}

func GetPsqlMaxPools() int32 {
	return cfg.PsqlMaxPools
}

func GetPsqlConnectingString() string {
	return fmt.Sprintf(
		"postgresql://%s:%s@%s:%s/%s",
		cfg.PsqlUser,
		cfg.PsqlPass,
		cfg.PsqlHost,
		cfg.PsqlPort,
		cfg.PsqlDB,
	)
}
