package rmq

import (
	"fmt"

	"github.com/dsbasko/yandex-go-diploma-1/core/logger"
	"github.com/rabbitmq/amqp091-go"
)

type Connector struct {
	conn *amqp091.Connection
	log  *logger.Logger
}

func Connect(log *logger.Logger, dsn string) (*Connector, error) {
	conn, err := amqp091.Dial(dsn)
	if err != nil {
		return nil, fmt.Errorf("amqp091.Dial: %w", err)
	}

	return &Connector{
		conn: conn,
		log:  log,
	}, nil
}
