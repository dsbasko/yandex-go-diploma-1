package rmq

import (
	"fmt"

	"github.com/rabbitmq/amqp091-go"
)

type BindConfig struct {
	Exchange string
	Queue    string
	Key      string
	NoWait   bool
	Args     amqp091.Table
}

func (c *Connector) Bind(bindCfg *BindConfig) error {
	ch, err := c.conn.Channel()
	if err != nil {
		return fmt.Errorf("conn.Channel: %w", err)
	}
	defer ch.Close()

	return ch.QueueBind(
		bindCfg.Queue,
		bindCfg.Key,
		bindCfg.Exchange,
		bindCfg.NoWait,
		bindCfg.Args,
	)
}
