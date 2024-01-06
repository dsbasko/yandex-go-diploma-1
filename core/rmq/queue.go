package rmq

import (
	"fmt"

	"github.com/rabbitmq/amqp091-go"
)

type QueueConfig struct {
	Name       string
	Durable    bool
	AutoDelete bool
	Exclusive  bool
	NoWait     bool
	Args       amqp091.Table
}

func (c *Connector) QueueDeclare(queueCfg *QueueConfig) (amqp091.Queue, error) {
	ch, err := c.conn.Channel()
	if err != nil {
		return amqp091.Queue{}, fmt.Errorf("conn.Channel: %w", err)
	}
	defer ch.Close()

	return ch.QueueDeclare(
		queueCfg.Name,
		queueCfg.Durable,
		queueCfg.AutoDelete,
		queueCfg.Exclusive,
		queueCfg.NoWait,
		queueCfg.Args,
	)
}
