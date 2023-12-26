module github.com/dsbasko/yandex-go-diploma-1/core

go 1.21

require (
	github.com/dsbasko/yandex-go-diploma-1/services/auth v1.0.0
	github.com/go-chi/chi/v5 v5.0.10
	github.com/google/uuid v1.5.0
	github.com/jackc/pgx/v5 v5.5.1
	github.com/rabbitmq/amqp091-go v1.9.0
	go.uber.org/zap v1.26.0
)

require (
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20221227161230-091c0ba34f0a // indirect
	github.com/jackc/puddle/v2 v2.2.1 // indirect
	go.uber.org/multierr v1.10.0 // indirect
	golang.org/x/crypto v0.14.0 // indirect
	golang.org/x/sync v0.4.0 // indirect
	golang.org/x/text v0.13.0 // indirect
)

replace github.com/dsbasko/yandex-go-diploma-1/services/auth => ../services/auth
