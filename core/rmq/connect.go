package rmq

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/dsbasko/yandex-go-diploma-1/core/logger"
	"github.com/rabbitmq/amqp091-go"
)

type Connector struct {
	conn *amqp091.Connection
	log  *logger.Logger
	dsn  string
}

type ConnectorOptions struct {
	MaxRetries   int
	RetryTimeOut time.Duration
}

func Connect(ctx context.Context, log *logger.Logger, dsn string, options ConnectorOptions) (*Connector, error) {
	conn, err := dial(ctx, dsn, &options)
	if err != nil {
		return nil, fmt.Errorf("dial: %w", err)
	}

	connector := Connector{
		conn: conn,
		log:  log,
		dsn:  dsn,
	}
	go connector.healthCheck(ctx, &options)

	return &connector, nil
}

func dial(ctx context.Context, dsn string, options *ConnectorOptions) (*amqp091.Connection, error) {
	for i := 0; i < options.MaxRetries; i++ {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
		}

		conn, err := amqp091.Dial(dsn)
		if err == nil && !conn.IsClosed() {
			return conn, nil
		}

		if i == options.MaxRetries-1 {
			return nil, fmt.Errorf("amqp091.Dial: %w", err)
		}

		time.Sleep(options.RetryTimeOut)
	}

	return nil, errors.New("unknown error")
}

func (c *Connector) healthCheck(ctx context.Context, options *ConnectorOptions) {
	ticker := time.NewTicker(options.RetryTimeOut)
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			if c.conn.IsClosed() {
				conn, err := dial(ctx, c.dsn, options)
				if err != nil {
					c.log.Errorf("dial: %v", err)
					cancel()
				}

				c.conn = conn
			}
		}
	}
}
