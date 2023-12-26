package rmq

import (
	"context"
	"fmt"

	"github.com/rabbitmq/amqp091-go"
)

type ConsumeConfig struct {
	Queue     string
	Consumer  string
	AutoAck   bool
	Exclusive bool
	NoLocal   bool
	NoWait    bool
	Args      amqp091.Table
}

func (c *Connector) Consume(
	ctx context.Context,
	cfgConsume *ConsumeConfig,
) (<-chan amqp091.Delivery, func(), error) {
	ch, err := c.conn.Channel()
	if err != nil {
		return nil, func() {}, fmt.Errorf("conn.Channel: %w", err)
	}

	msg, err := ch.ConsumeWithContext(
		ctx,
		cfgConsume.Queue,
		cfgConsume.Consumer,
		cfgConsume.AutoAck,
		cfgConsume.Exclusive,
		cfgConsume.NoLocal,
		cfgConsume.NoWait,
		cfgConsume.Args,
	)
	if err != nil {
		return nil, func() {}, fmt.Errorf("ConsumeWithContext: %w", err)
	}

	return msg, func() { ch.Close() }, nil
}

type SimpleConsumeConfig struct {
	Exchange  string
	Queue     string
	Key       string
	Consumer  string
	PullMsgFn func(msg amqp091.Delivery)
}

func (c *Connector) SimpleConsume(
	ctx context.Context,
	cfgConsume *SimpleConsumeConfig,
) error {
	if _, err := c.QueueDeclare(&QueueConfig{
		Name:    cfgConsume.Queue,
		Durable: true,
	}); err != nil {
		return fmt.Errorf("QueueDeclare: %w", err)
	}

	if err := c.Bind(&BindConfig{
		Exchange: cfgConsume.Exchange,
		Queue:    cfgConsume.Queue,
		Key:      cfgConsume.Key,
	}); err != nil {
		return fmt.Errorf("Bind: %w", err)
	}

	ch, err := c.conn.Channel()
	if err != nil {
		return fmt.Errorf("conn.Channel: %w", err)
	}

	msgs, err := ch.ConsumeWithContext(
		ctx,
		cfgConsume.Queue,
		cfgConsume.Consumer,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return fmt.Errorf("ConsumeWithContext: %w", err)
	}

	go func() {
		for msg := range msgs {
			select {
			case <-ctx.Done():
				ch.Close()
				return
			default:
			}

			c.log.Debugw(
				"message has been sent and wait response",
				"payload", string(msg.Body),
				"correlation_id", msg.CorrelationId,
			)

			cfgConsume.PullMsgFn(msg)
		}
	}()

	return nil
}
