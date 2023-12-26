package rmq

import (
	"fmt"

	"github.com/rabbitmq/amqp091-go"
)

type ExchangeConfig struct {
	Name       string
	Kind       string
	Durable    bool
	AutoDelete bool
	Internal   bool
	NoWait     bool
	Args       amqp091.Table
}

func (c *Connector) ExchangeDeclare(exchangeCfg *ExchangeConfig) error {
	ch, err := c.conn.Channel()
	if err != nil {
		return fmt.Errorf("conn.Channel: %w", err)
	}
	defer ch.Close()

	return ch.ExchangeDeclare(
		exchangeCfg.Name,
		exchangeCfg.Kind,
		exchangeCfg.Durable,
		exchangeCfg.AutoDelete,
		exchangeCfg.Internal,
		exchangeCfg.NoWait,
		exchangeCfg.Args,
	)
}
