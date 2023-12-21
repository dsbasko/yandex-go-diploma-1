package rmq

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/dsbasko/yandex-go-diploma-1/core/logger"
	"github.com/rabbitmq/amqp091-go"
)

const MaxRetries int = 6
const RetryTimeOut time.Duration = 5 * time.Second

type Connector struct {
	conn *amqp091.Connection
	log  *logger.Logger
	dsn  string
}

func Connect(ctx context.Context, log *logger.Logger, dsn string) (*Connector, error) {
	conn, err := dial(ctx, dsn)
	if err != nil {
		return nil, fmt.Errorf("dial: %w", err)
	}

	connector := Connector{
		conn: conn,
		log:  log,
		dsn:  dsn,
	}
	go connector.healthCheck(ctx)

	return &connector, nil
}

func dial(ctx context.Context, dsn string) (*amqp091.Connection, error) {
	for i := 0; i < MaxRetries; i++ {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
		}

		conn, err := amqp091.Dial(dsn)
		if err == nil && !conn.IsClosed() {
			return conn, nil
		}

		if i == MaxRetries-1 {
			return nil, fmt.Errorf("amqp091.Dial: %w", err)
		}

		time.Sleep(RetryTimeOut)
	}

	return nil, errors.New("unknown error")
}

func (c *Connector) healthCheck(ctx context.Context) {
	ticker := time.NewTicker(RetryTimeOut)
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			if c.conn.IsClosed() {
				conn, err := dial(ctx, c.dsn)
				if err != nil {
					c.log.Errorf("dial: %v", err)
					cancel()
				}

				c.conn = conn
			}
		}
	}
}
