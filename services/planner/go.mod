module github.com/dsbasko/yandex-go-diploma-1/services/planner

go 1.21

require (
	github.com/dsbasko/yandex-go-diploma-1/core v1.0.0
	github.com/dsbasko/yandex-go-diploma-1/services/auth v0.0.0-00010101000000-000000000000
	github.com/ilyakaznacheev/cleanenv v1.5.0
)

require (
	github.com/BurntSushi/toml v1.2.1 // indirect
	github.com/joho/godotenv v1.5.1 // indirect
	github.com/rabbitmq/amqp091-go v1.9.0 // indirect
	go.uber.org/multierr v1.10.0 // indirect
	go.uber.org/zap v1.26.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	olympos.io/encoding/edn v0.0.0-20201019073823-d3554ca0b0a3 // indirect
)

replace github.com/dsbasko/yandex-go-diploma-1/core => ../../core

replace github.com/dsbasko/yandex-go-diploma-1/services/auth => ../../services/auth
