package rmq

import (
	"fmt"

	"github.com/rabbitmq/amqp091-go"
)

type Connector struct {
	conn *amqp091.Connection
}

func Connect(dsn string) (*Connector, error) {
	conn, err := amqp091.Dial(dsn)
	if err != nil {
		return nil, fmt.Errorf("amqp091.Dial: %w", err)
	}

	return &Connector{
		conn: conn,
	}, nil
}
